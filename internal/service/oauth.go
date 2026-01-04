package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/TheCodeBreakerK/vanish-vault-api/configs"
	"github.com/TheCodeBreakerK/vanish-vault-api/internal/dto"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type googleUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type githubUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Login string `json:"login"`
}

// GetOauthConfig returns the OAuth2 configuration for the specified provider.
// It supports "google" and "github" providers based on the application configuration.
func GetOauthConfig(provider string, cfg *configs.Conf) *oauth2.Config {
	switch provider {
	case "google":
		return &oauth2.Config{
			ClientID:     cfg.GoogleClientID,
			ClientSecret: cfg.GoogleSecret,
			RedirectURL:  "http://localhost:8080/api/v1/auth/callback/google",
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		}
	case "github":
		return &oauth2.Config{
			ClientID:     cfg.GithubClientID,
			ClientSecret: cfg.GithubSecret,
			RedirectURL:  "http://localhost:8080/api/v1/auth/callback/github",
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		}
	default:
		return nil
	}
}

// FetchUserInfo retrieves profile information from OAuth2 providers.
// It uses the provided zap.Logger for consistent structured logging.
func FetchUserInfo(
	provider string,
	token string,
	log *zap.Logger,
) (*dto.UserInfoResponseDto, error) {
	urls := map[string]string{
		"google": "https://www.googleapis.com/oauth2/v2/userinfo",
		"github": "https://api.github.com/user",
	}

	url, ok := urls[provider]
	if !ok {
		log.Error("Unsupported OAuth provider requested", zap.String("provider", provider))
		return nil, errors.New("unsupported provider")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("Failed to create HTTP request", zap.Error(err))
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Failed to execute request to provider",
			zap.String("provider", provider),
			zap.Error(err),
		)
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		log.Error("Provider returned error status",
			zap.Int("status", resp.StatusCode),
			zap.String("provider", provider),
		)
		return nil, err
	}

	if provider == "google" {
		var u googleUser
		if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
			log.Error("Failed to decode Google user response", zap.Error(err))
			return nil, err
		}
		return &dto.UserInfoResponseDto{
			ID:    u.ID,
			Email: u.Email,
		}, nil
	}

	var u githubUser
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		log.Error("Failed to decode GitHub user response", zap.Error(err))
		return nil, err
	}

	return &dto.UserInfoResponseDto{
		ID:    strconv.Itoa(u.ID),
		Email: u.Email,
	}, nil
}
