package repositories

import (
	"testing"
	"webmalc/no-more-excuses/common/test"

	"github.com/stretchr/testify/assert"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, "/usr/bin/steam", c.Apps["steam"].Path)
	assert.Equal(t, "07:00-10:00", c.Apps["game"].Weekdays["mon"])
	assert.Equal(t, "", c.Apps["steam"].Weekdays["tue"])
	assert.Equal(t, "08:00-11:00", c.Apps["steam"].Weekdays["mon"])
	assert.Equal(t, "07:00-18:00", c.Apps["steam"].BaseSchedule)
	assert.Equal(t, "09:00-15:00", c.Apps["game"].BaseSchedule)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
