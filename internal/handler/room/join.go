// Package room contains handlers related to secure room management.
package room

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewJoinRoomHandler handles the process of joining a secure room.
// @Summary      Join Room
// @Description  Allows an authenticated user to join a room (may require password in body).
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id         path      string          true  "Room ID"
// @Param        request    body      map[string]any  false "Room password (if applicable)"
// @Success      200        {object}  map[string]any "Join confirmation"
// @Failure      403        {object}  map[string]any "Incorrect password or access denied"
// @Router       /api/v1/rooms/{id}/join [post]
func NewJoinRoomHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Join room endpoint not yet implemented",
		})
	}
}
