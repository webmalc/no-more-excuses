package utils

import (
	"testing"
	"time"
	"webmalc/no-more-excuses/common/test"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	test.Run(m)
}

func TestGetWeek(t *testing.T) {
	assert.Equal(t, []time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
		time.Sunday,
	}, GetWeek())
}
