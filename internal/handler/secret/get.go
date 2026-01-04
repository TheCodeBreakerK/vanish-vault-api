package secret

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewGetSecretHandler handles retrieving and decrypting a specific secret.
// @Summary      Read Secret (Decrypt)
// @Description  Retrieves and decrypts the content of a specific secret. May trigger "Burn on Read" rule (delete after reading).
// @Tags         Secrets
// @Produce      json
// @Security     BearerAuth
// @Param        id         path      string  true  "Room ID (UUID)"
// @Param        secretId   path      string  true  "Secret ID (UUID)"
// @Success      200        {object}  map[string]any "Decrypted secret content"
// @Failure      404        {object}  map[string]any "Secret not found or already expired"
// @Router       /api/v1/rooms/{id}/secrets/{secretId} [get]
func NewGetSecretHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Get secret endpoint not yet implemented",
		})
	}
}
