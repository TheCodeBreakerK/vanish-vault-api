// Package auth contains handlers related to authentication.
package auth

import (
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/configs"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/dto"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// NewLoginHandler initiates the OAuth2 login process.
// @Summary      Initiate OAuth2 Login
// @Description  Redirects to the auth provider or returns the URL based on the Accept header.
// @Tags         Auth
// @Produce      json
// @Param        provider   path      string  true  "Auth Provider" Enums(google, github)
// @Success      200        {object}  dto.LoginResponseDto "Returns JSON with auth URL"
// @Success      307        {string}  string  "Temporary Redirect to Provider"
// @Failure      400        {object}  dto.ErrorResponseDto "Invalid provider"
// @Failure      500        {object}  dto.ErrorResponseDto "Internal server error"
// @Router       /api/v1/auth/login/{provider} [get]
func NewLoginHandler(cfg *configs.Conf, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		provider := c.Param("provider")

		oauthConfig := service.GetOauthConfig(provider, cfg)
		if oauthConfig == nil {
			log.Warn("Attempt to log in with an invalid provider", zap.String("provider", provider))
			c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{
				Code:    http.StatusBadRequest,
				Message: "Login provider not supported.",
				Status:  http.StatusText(http.StatusBadRequest),
			})
			return
		}

		state, err := service.GenerateRandomState()
		if err != nil {
			log.Error("Failed to generate random state", zap.Error(err))
			c.JSON(http.StatusInternalServerError, dto.ErrorResponseDto{
				Code:    http.StatusInternalServerError,
				Message: "Internal error when starting authentication",
				Status:  http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		c.SetCookie("oauth_state", state, 900, "/", "", true, true)

		url := oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)

		if c.GetHeader("Accept") != "application/json" {
			log.Info("Redirecting user to provider", zap.String("provider", provider))
			c.Redirect(http.StatusTemporaryRedirect, url)
			return
		}

		c.JSON(http.StatusOK, dto.LoginResponseDto{URL: url})
	}
}
