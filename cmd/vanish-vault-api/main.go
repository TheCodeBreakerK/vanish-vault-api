// Package main is the entry point for the VanishVault API server.
package main

import (
	"context"

	_ "github.com/TheCodeBreakerK/vanish-vault-api/api/docs"
	"github.com/TheCodeBreakerK/vanish-vault-api/configs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
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
	log := configs.GetLogger()
	defer configs.Sync()

	cfg := configs.LoadConfig()

	dbPool := configs.NewDatabase(context.Background(), cfg)
	defer dbPool.Close()

	r := gin.Default()

	r.Any("/healthz", func(c *gin.Context) {
		if err := dbPool.Ping(c); err != nil {
			c.JSON(500, gin.H{"status": "error", "db": "disconnected"})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Info("Starting server...", zap.String("port", "8080"), zap.String("env", cfg.Environment))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Error("Failed to start server", zap.Error(err))
	}
}
