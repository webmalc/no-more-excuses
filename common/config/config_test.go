package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// Should return the config filename based on the environment variable.
func Test_getFilename(t *testing.T) {
	assert.Equal(t, "config", getFilename())

	os.Setenv("NO_MORE_EXCUSES_ENV", "test")
	assert.Equal(t, "config.test", getFilename())

	os.Setenv("NO_MORE_EXCUSES_ENV", "prod")
	assert.Equal(t, "config.prod", getFilename())
}

// Should setup the main configuration.
func TestSetup(t *testing.T) {
	timezone := viper.GetString("timezone")
	assert.Equal(t, "", timezone)

	os.Setenv("NO_MORE_EXCUSES_ENV", "test")
	Setup()
	timezone = viper.GetString("timezone")
	assert.Equal(t, "Europe/Rome", timezone)
}

// Should panic with the invalid environment variable.
func TestSetupPanic(t *testing.T) {
	os.Setenv("NO_MORE_EXCUSES_ENV", "invalid")
	assert.Panics(t, Setup)
}

// Should set default values.
func Test_setDefaults(t *testing.T) {
	viper.Reset()
	assert.Empty(t, viper.GetString("log_path"))
	assert.False(t, viper.IsSet("is_prod"))
	setDefaults("test")
	assert.True(t, viper.IsSet("is_prod"))
	assert.Equal(t, "UTC", viper.GetString("timezone"))
}
