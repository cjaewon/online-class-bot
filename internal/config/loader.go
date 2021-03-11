package config

import (
	"onlineclassbot/internal/logger"
	"runtime"

	"github.com/spf13/viper"
)

// Config is the onlineclassbot's base config structure
type Config struct {
	// Schedules is online class calendar, key is a cron expression, value is the response
	Schedules map[string]string
}

// ReadInConfig loads configuration and returns unmashaled instances
func ReadInConfig() *Config {
	v := viper.New()

	v.SetConfigName(".onlineclassbot")
	v.SetConfigType("toml")

	if runtime.GOOS != "windows" {
		v.AddConfigPath("$HOME")
	}

	v.AddConfigPath(".")
	v.SetDefault("schedules", map[string]string{})

	if err := v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			logger.Log().Infow("Using a default configuration")
		default:
			logger.Log().Fatalw("Failed to load a config file", "err", err)
		}
	} else {
		logger.Log().Infow("Using a provided configuration", "config_file", v.ConfigFileUsed())
	}

	var config Config

	if err := v.Unmarshal(&config); err != nil {
		logger.Log().Fatalw("Failed to unmarshal config file", "err", err)
	}

	return &config
}
