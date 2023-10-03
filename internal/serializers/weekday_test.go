package serializers

import (
	"testing"
	"time"
	"webmalc/no-more-excuses/common/test"

	"github.com/stretchr/testify/assert"
)

func TestWeekday_Deserialize(t *testing.T) {
	weekday, err := NewWeekday("mon").Deserialize()
	assert.Nil(t, err)
	assert.Equal(t, time.Monday, weekday)

	weekday, _ = NewWeekday("sun").Deserialize()
	assert.Equal(t, time.Sunday, weekday)

	_, err = NewWeekday("invalid").Deserialize()
	assert.Error(t, err)
}

func TestNewWeekday(t *testing.T) {
	serializer := NewWeekday("mon")
	assert.Equal(t, "mon", serializer.weekday)
}

func TestMain(m *testing.M) {
	test.Run(m)
}
