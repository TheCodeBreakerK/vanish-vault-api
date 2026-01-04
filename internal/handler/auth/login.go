// Package auth contains handlers related to authentication.
package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/TheCodeBreakerK/vanish-vault-api/configs"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/dto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

// NewLoginHandler initiates the OAuth2 login process.
// @Summary      Get OAuth2 Login URL
// @Description  Returns the URL to redirect the user to the specified authentication provider or redirects automatically based on Accept header.
// @Tags         Auth
// @Produce      json
// @Param        provider   path      string  true  "OAuth2 Provider (e.g., google, github)" Enums(google, github)
// @Success      200        {object}  dto.LoginResponse "Returns JSON with auth URL (if Accept: application/json)"
// @Success      307        {string}  string  "Redirects to provider (if accessed via browser)"
// @Failure      400        {object}  dto.ErrorResponse "Invalid provider"
// @Failure      500        {object}  dto.ErrorResponse "Internal Server Error"
// @Router       /api/v1/auth/login/{provider} [get]
func NewLoginHandler(cfg *configs.Conf, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		oauthConf := map[string]*oauth2.Config{
			"google": {
				ClientID:     cfg.GoogleClientID,
				ClientSecret: cfg.GoogleSecret,
				RedirectURL:  "http://localhost/api/v1/auth/callback/google",
				Scopes: []string{
					"https://www.googleapis.com/auth/userinfo.email",
					"https://www.googleapis.com/auth/userinfo.profile",
				},
				Endpoint: google.Endpoint,
			},
			"github": {
				ClientID:     cfg.GithubClientID,
				ClientSecret: cfg.GithubSecret,
				RedirectURL:  "http://localhost/api/v1/auth/callback/github",
				Scopes:       []string{"user:email"},
				Endpoint:     github.Endpoint,
			},
		}

		provider := c.Param("provider")

		config, exists := oauthConf[provider]
		if !exists {
			log.Warn("Attempt to log in with an invalid provider", zap.String("provider", provider))
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: "Login provider not supported.",
			})
			return
		}

		state, err := generateRandomState()
		if err != nil {
			log.Error("Failed to generate random state", zap.Error(err))
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: "Internal error when starting authentication",
			})
			return
		}

		c.SetCookie("oauth_state", state, 900, "/", "", true, true)

		url := config.AuthCodeURL(state, oauth2.AccessTypeOffline)

		if c.GetHeader("Accept") != "application/json" {
			log.Info("Redirecting user to provider", zap.String("provider", provider))
			c.Redirect(http.StatusTemporaryRedirect, url)
			return
		}

		c.JSON(http.StatusOK, dto.LoginResponse{URL: url})
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
