package room

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewListRoomsHandler handles listing rooms available to the user.
// @Summary      List Rooms
// @Description  Lists rooms available to the user (those created by them or public, depending on rules).
// @Tags         Rooms
// @Produce      json
// @Security     BearerAuth
// @Success      200        {array}   map[string]any "List of rooms"
// @Failure      401        {object}  map[string]any "Unauthorized"
// @Router       /api/v1/rooms [get]
func NewListRoomsHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "List rooms endpoint not yet implemented",
		})
	}
}
