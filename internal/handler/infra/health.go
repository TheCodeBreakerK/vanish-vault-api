// Package infra contains handlers related to infrastructure.
package infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// NewHealthCheckHandler handles the health check of the application.
// @Summary      Health Check
// @Description  Checks if the service and its dependencies (database) are operational.
// @Tags         Infra
// @Produce      json
// @Success      200  {object}  map[string]string "Service is up and running"
// @Failure      503  {object}  map[string]string "Service or dependencies are down"
// @Router       /healthz [get]
func NewHealthCheckHandler(log *zap.Logger, db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("Health check requested")

		if err := db.Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "error",
				"db": "disconnected",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"service": "vanish-vault",
		})
	}
}
