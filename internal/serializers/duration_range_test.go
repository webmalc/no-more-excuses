package serializers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurationRange_Serialize(t *testing.T) {
	assert.Equal(t, "00:00-00:00", NewDurationRange(0, 0, "").Serialize())
	assert.Equal(t, "02:15-04:30", NewDurationRange(
		2*time.Hour+15*time.Minute, 4*time.Hour+30*time.Minute, "",
	).Serialize())
}

func TestDurationRange_Deserialize(t *testing.T) {
	_, _, err := NewDurationRange(0, 0, "invalid").Deserialize()
	assert.Contains(t, err.Error(), "HH:MM-HH:MM")

	_, _, err = NewDurationRange(0, 0, "foo:bar-foo:bar").Deserialize()
	assert.Contains(t, err.Error(), "Atoi")

	_, _, err = NewDurationRange(0, 0, "19:00-10:00").Deserialize()
	assert.Contains(t, err.Error(), "> end time")

	start, end, err := NewDurationRange(0, 0, "10:30-19:00").Deserialize()
	assert.Nil(t, err)
	assert.Equal(t, 10*time.Hour+30*time.Minute, start)
	assert.Equal(t, 19*time.Hour, end)
}

func TestNewDurationRange(t *testing.T) {
	serializer := NewDurationRange(2*time.Hour, 4*time.Hour, "10:15")
	assert.Equal(t, "10:15", serializer.timeRange)
	assert.Equal(t, 2*time.Hour, serializer.start.duration)
	assert.Equal(t, 4*time.Hour, serializer.end.duration)
}
