// Package infra contains handlers related to infrastructure.
package infra

import (
	"net/http"
	"time"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// NewHealthCheckHandler handles the health check of the application.
// @Summary      Health Check
// @Description  Checks if the service and its dependencies (database) are operational.
// @Tags         Infra
// @Produce      json
// @Success      200  {object}  dto.HealthCheckResponse "Service is up and running"
// @Failure      503  {object}  dto.ErrorResponse "Service or dependencies are down"
// @Router       /healthz [get]
func NewHealthCheckHandler(log *zap.Logger, db *pgxpool.Pool, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("Health check requested")

		if err := db.Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, dto.ErrorResponse{
				Code:    http.StatusServiceUnavailable,
				Status:  http.StatusText(http.StatusServiceUnavailable),
				Message: "Database connection error",
			})
			return
		}

		if err := rdb.Ping(c.Request.Context()).Err(); err != nil {
			c.JSON(http.StatusServiceUnavailable, dto.ErrorResponse{
				Code:    http.StatusServiceUnavailable,
				Status:  http.StatusText(http.StatusServiceUnavailable),
				Message: "Redis connection error",
			})
			return
		}

		c.JSON(http.StatusOK, dto.HealthCheckResponse{
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
			TS:      time.Now(),
			Service: "vanish-vault",
		})
	}
}
