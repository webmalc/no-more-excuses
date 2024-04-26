package server

import (
	"testing"
	"time"

	"webmalc/no-more-excuses/common/test"

	"github.com/stretchr/testify/assert"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, time.Second*2, c.IntervalDuration)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
