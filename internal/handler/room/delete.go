// Package room contains handlers related to secure room management.
package room

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewDeleteRoomHandler handles the deletion of a secure room.
// @Summary      Delete Room
// @Description  Permanently removes a room and all associated secrets (Immediate purge).
// @Tags         Rooms
// @Security     BearerAuth
// @Param        id         path      string  true  "Room ID (UUID)"
// @Success      204        "No Content - Room successfully deleted"
// @Failure      403        {object}  map[string]any "No permission to delete this room"
// @Router       /api/v1/rooms/{id} [delete]
func NewDeleteRoomHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Delete room endpoint not yet implemented",
		})
	}
}
