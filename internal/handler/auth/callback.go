// Package auth contains handlers related to authentication.
package auth

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewCallbackHandler handles the OAuth2 callback, exchanges the authorization code for tokens, and generates access JWT.
// @Summary      OAuth2 Callback
// @Description  Receives the authorization code from the provider, exchanges it for tokens, and generates the VanishVault access JWT.
// @Tags         Auth
// @Param        provider   path      string  true  "OAuth2 Provider" Enums(google, github)
// @Param        code       query     string  true  "Authorization code returned by the provider"
// @Param        state      query     string  true  "CSRF security state"
// @Success      200        {object}  map[string]string "Access tokens (Access Token and Refresh Token)"
// @Failure      400        {object}  map[string]any "Token exchange failure or invalid state"
// @Failure      500        {object}  map[string]any "Internal error processing login"
// @Router       /api/v1/auth/callback/{provider} [get]
func NewCallbackHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Callback endpoint not yet implemented",
		})
	}
}
