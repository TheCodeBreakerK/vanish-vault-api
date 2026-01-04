package router

import (
	authHandler "github.com/TheCodeBreakerK/vanish-vault-api/internal/handler/auth"
	infraHandler "github.com/TheCodeBreakerK/vanish-vault-api/internal/handler/infra"
	roomHandler "github.com/TheCodeBreakerK/vanish-vault-api/internal/handler/room"
	secretHandler "github.com/TheCodeBreakerK/vanish-vault-api/internal/handler/secret"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Router) setupRoutes(engine *gin.Engine) {
	r.log.Info("Setting up all routes")

	repo := repository.New(r.db)

	engine.GET("/healthz", infraHandler.NewHealthCheckHandler(r.log, r.db))
	engine.HEAD("/healthz", infraHandler.NewHealthCheckHandler(r.log, r.db))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := engine.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.GET("/login/:provider", authHandler.NewLoginHandler(r.cfg, r.log))
		auth.GET("/callback/:provider", authHandler.NewCallbackHandler(repo, r.log))
		auth.POST("/refresh", authHandler.NewRefreshHandler(repo, r.log))
	}

	rooms := v1.Group("/rooms")
	{
		rooms.POST("", roomHandler.NewCreateRoomHandler(repo, r.log))
		rooms.GET("", roomHandler.NewListRoomsHandler(repo, r.log))

		roomID := rooms.Group("/:id")
		{
			roomID.GET("", roomHandler.NewGetRoomHandler(repo, r.log))
			roomID.DELETE("", roomHandler.NewDeleteRoomHandler(repo, r.log))
			roomID.POST("/join", roomHandler.NewJoinRoomHandler(repo, r.log))
			roomID.POST("/leave", roomHandler.NewLeaveRoomHandler(repo, r.log))

			secrets := roomID.Group("/secrets")
			{
				secrets.POST("", secretHandler.NewCreateSecretHandler(repo, r.log))
				secrets.GET("", secretHandler.NewListSecretsHandler(repo, r.log))
				secrets.GET("/:secretId", secretHandler.NewGetSecretHandler(repo, r.log))
			}
		}
	}
}
