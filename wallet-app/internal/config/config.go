package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL     string `mapstructure:"DATABASE_URL"`
	Port            string `mapstructure:"PORT"`
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	Environment     string `mapstructure:"ENVIRONMENT"`
	RedisURL        string `mapstructure:"REDIS_URL"`
	MaxTransactions int    `mapstructure:"MAX_TRANSACTIONS"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
