package repositories

import (
	"github.com/spf13/viper"
)

// Config is the repositories configuration struct.
type Config struct {
	// Apps map[string]map[string]interface{} `mapstructure:"apps"`
	Apps map[string]ConfigItem `mapstructure:"apps"`
}

type ConfigItem struct {
	Path         string            `mapstructure:"path"`
	BaseSchedule string            `mapstructure:"base_schedule"`
	Weekdays     map[string]string `mapstructure:"weekdays"`
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}
