// Package configs handles application configuration and environment variables.
package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Conf holds all the configuration for the application.
type Conf struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Debug       bool   `mapstructure:"DEBUG"`

	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
}

// LoadConfig reads the .env file and unmarshals it into the Conf struct.
func LoadConfig() *Conf {
	log := GetLogger()

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("ENVIRONMENT", "development")

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
