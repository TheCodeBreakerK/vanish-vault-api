// Package auth contains handlers related to authentication.
package auth

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewLoginHandler initiates the OAuth2 login process by redirecting to the provider's auth page.
// @Summary      Initiate OAuth2 Login
// @Description  Redirects the user to the specified authentication provider (Google, GitHub) to start the secure session.
// @Tags         Auth
// @Param        provider   path      string  true  "OAuth2 Provider (e.g., google, github)" Enums(google, github)
// @Success      302        {string}  string  "Temporary redirection to the provider"
// @Failure      400        {object}  map[string]any "Invalid provider"
// @Router       /api/v1/auth/login/{provider} [get]
func NewLoginHandler(repo repository.Querier, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Login endpoint not yet implemented",
		})
	}
}
