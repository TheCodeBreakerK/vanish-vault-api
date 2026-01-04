package auth

import (
	"net/http"
	"time"

	"github.com/TheCodeBreakerK/vanish-vault-api/configs"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/dto"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/repository"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

// NewCallbackHandler handles the OAuth2 callback.
// @Summary      OAuth2 Callback
// @Description  Exchanges authorization code for a VanishVault JWT access token.
// @Tags         Auth
// @Produce      json
// @Param        provider   path      string  true  "google or github"
// @Param        code       query     string  true  "Authorization code"
// @Param        state      query     string  true  "CSRF state"
// @Success      200        {object}  dto.CallbackResponseDto
// @Failure      401        {object}  dto.ErrorResponseDto "Unauthorized or invalid state"
// @Failure      500        {object}  dto.ErrorResponseDto "Failed to process authentication"
// @Router       /api/v1/auth/callback/{provider} [get]
func NewCallbackHandler(
	repo repository.Querier,
	cfg *configs.Conf,
	log *zap.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		provider := c.Param("provider")
		state := c.Query("state")
		code := c.Query("code")

		cookieState, err := c.Cookie("oauth_state")
		if err != nil || state != cookieState {
			log.Warn("Invalid state parameter", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponseDto{
				Code:    http.StatusUnauthorized,
				Message: "Invalid state parameter",
				Status:  http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		c.SetCookie("oauth_state", "", -1, "/", "", true, true)

		oauthConfig := service.GetOauthConfig(provider, cfg)
		if oauthConfig == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponseDto{
				Code:    http.StatusBadRequest,
				Message: "Invalid provider",
				Status:  http.StatusText(http.StatusBadRequest),
			})
			return
		}

		token, err := oauthConfig.Exchange(c.Request.Context(), code)
		if err != nil {
			log.Error("Failed to exchange token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponseDto{
				Code:    http.StatusInternalServerError,
				Message: "Failed to exchange token",
				Status:  http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		userInfo, err := service.FetchUserInfo(provider, token.AccessToken, log)
		if err != nil {
			log.Error("Failed to fetch user info", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponseDto{
				Code:    http.StatusInternalServerError,
				Message: "Failed to fetch user info",
				Status:  http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		dbProvider := repository.AuthProviderType(provider)

		user, err := repo.GetUserByProvider(c, repository.GetUserByProviderParams{
			Provider:   dbProvider,
			ProviderID: userInfo.ID,
		})

		if err != nil {
			log.Info("User not found, creating new user", zap.String("email", userInfo.Email))

			user, err = repo.CreateUser(c, repository.CreateUserParams{
				Email: pgtype.Text{
					String: userInfo.Email,
					Valid:  userInfo.Email != "",
				},
				Provider:   dbProvider,
				ProviderID: userInfo.ID,
			})

			if err != nil {
				log.Error("Failed to create user in db", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponseDto{
					Code:    http.StatusInternalServerError,
					Message: "Failed to create user account",
					Status:  http.StatusText(http.StatusInternalServerError),
				})
				return
			}
		}

		jwtToken, err := service.GenerateToken(user.ID, cfg)
		if err != nil {
			log.Error("Failed to generate JWT", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponseDto{
				Code:    http.StatusInternalServerError,
				Message: "Failed to generate session token",
				Status:  http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		c.JSON(http.StatusOK, dto.CallbackResponseDto{
			Token:     jwtToken,
			TokenType: "Bearer",
			ExpiryAt:  time.Now().Add(time.Hour * time.Duration(cfg.JWTExpirationHours)).Unix(),
		})
	}
}
