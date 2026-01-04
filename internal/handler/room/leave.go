package room

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewLeaveRoomHandler handles the process of leaving a secure room.
// @Summary      Leave Room
// @Description  Removes the current user from the room's participant list.
// @Tags         Rooms
// @Security     BearerAuth
// @Param        id         path      string  true  "Room ID"
// @Success      200        {object}  map[string]any "Leave confirmation"
// @Router       /api/v1/rooms/{id}/leave [post]
func NewLeaveRoomHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Leave room endpoint not yet implemented",
		})
	}
}
