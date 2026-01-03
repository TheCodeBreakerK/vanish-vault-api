// Package secret provides handlers for distribution of encrypted secrets.
package secret

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewListSecretsHandler handles listing all secrets in a room.
// @Summary      List Room Secrets
// @Description  Lists secrets available in the room. Sensitive content may be hidden in this listing.
// @Tags         Secrets
// @Produce      json
// @Security     BearerAuth
// @Param        id         path      string  true  "Room ID (UUID)"
// @Success      200        {array}   map[string]any "List of secrets (metadata)"
// @Failure      403        {object}  map[string]any "Access denied to the room"
// @Router       /api/v1/rooms/{id}/secrets [get]
func NewListSecretsHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "List secrets endpoint not yet implemented",
		})
	}
}
