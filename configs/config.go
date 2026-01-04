// Package configs handles application configuration and environment variables.
package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Conf holds all the configuration for the application.
type Conf struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBUser         string `mapstructure:"POSTGRES_USER"`
	DBPassword     string `mapstructure:"POSTGRES_PASSWORD"`

	GoogleClientID string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleSecret   string `mapstructure:"GOOGLE_CLIENT_SECRET"`

	GithubClientID string `mapstructure:"GITHUB_CLIENT_ID"`
	GithubSecret   string `mapstructure:"GITHUB_CLIENT_SECRET"`

	RedisAddr      string `mapstructure:"REDIS_ADDR"`
	RedisPassword  string `mapstructure:"REDIS_PASSWORD"`
	RedisDB        int    `mapstructure:"REDIS_DB"`
}

// LoadConfig reads the .env file and unmarshals it into the Conf struct.
func LoadConfig(log *zap.Logger) *Conf {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("POSTGRES_PORT", "5432")

	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_ADDR", "redis:${REDIS_PORT}")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Error("Failed to read config file", zap.Error(err))
			return nil
		}
		log.Info("Config file not found, using environment variables")
	}

	var cfg Conf
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error("Failed to unmarshal config", zap.Error(err))
		return nil
	}

	return &cfg
}

// GetDBURL returns the formatted postgres connection string.
func (c *Conf) GetDBURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

// GetLogger returns the global zap logger instance.
func GetLogger() *zap.Logger {
	ensureInitialized()
	return log
}
