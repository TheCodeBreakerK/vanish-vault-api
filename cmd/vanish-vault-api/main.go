// Package main is the entry point for the VanishVault API server.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/TheCodeBreakerK/vanish-vault-api/api/docs"
)

// @title           VanishVault API
// @version         1.0.0
// @description     A secure, ephemeral secret-sharing service with OAuth2 and private rooms.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     https://github.com/TheCodeBreakerK/vanish-vault-api/support
// @contact.email   kelvin.oliveira.dev@pm.me

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()

	r.Any("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
		})
	})

	runDBMigrations()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func runDBMigrations() {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		user, pass, host, dbName)

	m, err := migrate.New(
		"file://db/migrations",
		dbURL,
	)

	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("could not apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
