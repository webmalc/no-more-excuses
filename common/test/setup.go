package test

import (
	"os"
	"testing"

	"webmalc/no-more-excuses/common/config"
)

// Setups the tests.
func setUp() {
	os.Setenv("NO_MORE_EXCUSES_ENV", "test")
	config.Setup()
}

// Run setups, runs and teardown the tests.
func Run(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
