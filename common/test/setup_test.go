package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	assert.Equal(t, "test", os.Getenv("NO_MORE_EXCUSES_ENV"))
}

func TestMain(m *testing.M) {
	Run(m)
}
