package room

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewGetRoomHandler handles fetching details of a specific room.
// @Summary      Get Room Details
// @Description  Returns metadata for a specific room by ID.
// @Tags         Rooms
// @Produce      json
// @Security     BearerAuth
// @Param        id         path      string  true  "Room ID (UUID)"
// @Success      200        {object}  map[string]any "Room data"
// @Failure      404        {object}  map[string]any "Room not found"
// @Router       /api/v1/rooms/{id} [get]
func NewGetRoomHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Get room endpoint not yet implemented",
		})
	}
}
