package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const prefix = "no_more_excuses"

// getFilename returns filename based on the environment variable.
func getFilename() string {
	fileName := "config"
	env := os.Getenv(strings.ToUpper(prefix + "_env"))
	if env != "" {
		fileName += "." + env
	}

	return fileName
}

// setDefaults sets default values.
func setDefaults(baseDir string) {
	viper.Set("base_dir", filepath.Dir(filepath.Dir(baseDir))+"/")
	viper.SetDefault("is_prod", false)
	viper.SetDefault("timezone", "UTC")
	secondsInHour := 60
	viper.SetDefault("run_per_seconds", secondsInHour)
}

// setPaths set paths.
func setPaths(baseDir string) {
	viper.AddConfigPath(".")
	viper.AddConfigPath(baseDir + "/yaml")
	viper.AddConfigPath("/etc/" + prefix)
	viper.AddConfigPath("$HOME/." + prefix)
	viper.AddConfigPath("$HOME/.config/" + prefix)
}

// setEnv sets the environment.
func setEnv() {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
}

// getBaseDir get the base directory.
func getBaseDir() string {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		panic("config: unable to determine the caller.")
	}

	return filepath.Dir(b)
}

// SetTimezone sets timezone.
func SetTimezone() {
	loc, _ := time.LoadLocation(viper.GetString("timezone"))
	time.Local = loc
}

// read reads configuration.
func read() {
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// Setup initializes the main configuration.
func Setup() {
	baseDir := getBaseDir()
	viper.SetConfigName(getFilename())
	setPaths(baseDir)
	setEnv()
	setDefaults(baseDir)
	read()
}
