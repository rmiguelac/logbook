package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
}

func Load() Config {
	viper.AutomaticEnv()
	return Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
	}
}
