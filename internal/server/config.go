package server

import (
	"time"

	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type Config struct {
	IntervalDuration time.Duration
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{
		IntervalDuration: time.Second * time.Duration(
			viper.GetInt("run_per_seconds"),
		),
	}

	return config
}
