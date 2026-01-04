// Package router contains the route definitions for the API server.
package router

import (
	"github.com/TheCodeBreakerK/vanish-vault-api/configs"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Router struct holds the configuration and handlers for setting up routes.
type Router struct {
	cfg *configs.Conf
	log *zap.Logger
	db  *pgxpool.Pool
	rdb *redis.Client
}

// NewRouter creates a new Router instance with the given configuration and handlers.
func NewRouter(
	cfg *configs.Conf,
	log *zap.Logger,
	db *pgxpool.Pool,
	rdb *redis.Client,
) *Router {
	return &Router{
		cfg: cfg,
		log: log,
		db:  db,
		rdb: rdb,
	}
}

// Setup initializes the Gin engine, sets up routes, and returns the configured engine.
func (r *Router) Setup() {
	router := gin.New()
	gin.SetMode(r.cfg.GinMode)

	router.Use(gin.Recovery())

	r.setupRoutes(router)

	r.log.Info("Starting server on port 8080")

	if err := router.Run(":8080"); err != nil {
		r.log.Fatal("Failed to start server", zap.Error(err))
	}
}
