package auth

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewRefreshHandler handles the token refresh process.
// @Summary      Refresh JWT Token
// @Description  Generates a new Access Token using a valid Refresh Token, keeping the session active without re-login.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request    body      map[string]string  true  "Object containing the refresh_token"
// @Success      200        {object}  map[string]string "New Access Token"
// @Failure      401        {object}  map[string]any "Invalid or expired Refresh Token"
// @Router       /api/v1/auth/refresh [post]
func NewRefreshHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Refresh endpoint not yet implemented",
		})
	}
}
