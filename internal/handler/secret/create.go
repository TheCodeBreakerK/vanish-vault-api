// Package secret provides handlers for distribution of encrypted secrets.
package secret

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewCreateSecretHandler handles the creation of a new secret within a room.
// @Summary      Add Secret
// @Description  Stores a new encrypted secret (AES-256) within a specific room.
// @Tags         Secrets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id         path      string          true  "Room ID (UUID)"
// @Param        request    body      map[string]any  true  "Secret content and visibility settings"
// @Success      201        {object}  map[string]any "Created secret metadata (ID, date)"
// @Failure      403        {object}  map[string]any "User is not a member of the room"
// @Failure      404        {object}  map[string]any "Room not found"
// @Router       /api/v1/rooms/{id}/secrets [post]
func NewCreateSecretHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Create secret endpoint not yet implemented",
		})
	}
}
