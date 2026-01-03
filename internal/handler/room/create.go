// Package room contains handlers related to secure room management.
package room

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewCreateRoomHandler handles the creation of a new secure room.
// @Summary      Create Secure Room
// @Description  Creates a new encrypted private room for ephemeral data sharing.
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request    body      map[string]any  true  "Room data (name, optional password, expiration time)"
// @Success      201        {object}  map[string]any "Created room details"
// @Failure      400        {object}  map[string]any "Invalid input data"
// @Failure      401        {object}  map[string]any "Unauthorized"
// @Router       /api/v1/rooms [post]
func NewCreateRoomHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Create room endpoint not yet implemented",
		})
	}
}
