package configs

import (
	"context"
	"time"

	migrations "github.com/TheCodeBreakerK/vanish-vault-api/db/migrations"
	migrate "github.com/golang-migrate/migrate/v4"

	// Driver necessary for golang-migrate to talk to postgres
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	iofs "github.com/golang-migrate/migrate/v4/source/iofs"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	zap "go.uber.org/zap"
)

// NewDatabase initializes a new pgx connection pool and runs migrations.
func NewDatabase(ctx context.Context, cfg *Conf) *pgxpool.Pool {
	log := GetLogger()
	dsn := cfg.GetDBURL()

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("Failed to parse database config", zap.Error(err))
	}

	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatal("Failed to create database pool", zap.Error(err))
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Failed to ping database", zap.Error(err))
	}

	log.Info("Connected to the database successfully")

	runMigrations(dsn, log)

	return pool
}

func runMigrations(dbURL string, log *zap.Logger) {
	d, err := iofs.New(migrations.FS, ".")
	if err != nil {
		log.Fatal("Failed to create migration source", zap.Error(err))
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dbURL)
	if err != nil {
		log.Fatal("Failed to create migration instance", zap.Error(err))
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to apply migrations", zap.Error(err))
	}

	log.Info("Database migrations applied successfully")
}
