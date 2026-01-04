// Package main is the entry point for the VanishVault API server.
package main

import (
	"context"

	// Importing the docs package to register Swagger documentation
	_ "github.com/TheCodeBreakerK/vanish-vault-api/api/docs"
	"github.com/TheCodeBreakerK/vanish-vault-api/configs"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/router"
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

// @host            localhost

// @tag.name         Infra
// @tag.description  Service health, monitoring and system-level endpoints

// @tag.name         Auth
// @tag.description  Authentication and session management

// @tag.name         Rooms
// @tag.description  Private encrypted room management

// @tag.name         Secrets
// @tag.description  Ephemeral secret management and P2P messaging

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	log := configs.GetLogger()
	defer configs.Sync()

	cfg := configs.LoadConfig(log)

	ctx := context.Background()

	dbPool := configs.NewDatabase(ctx, cfg, log)
	defer dbPool.Close()

	rdb := configs.NewRedisClient(ctx, cfg, log)
	defer rdb.Close()

	appRouter := router.NewRouter(cfg, log, dbPool, rdb)
	appRouter.Setup()
}
