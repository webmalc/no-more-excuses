package serializers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration_Serialize(t *testing.T) {
	serializer := NewDuration(10*time.Hour+15*time.Minute, "")
	assert.Equal(t, "10:15", serializer.Serialize())
}

func TestDuration_Deserialize(t *testing.T) {
	result, err := NewDuration(0, "14:15").Deserialize()
	assert.Nil(t, err)
	assert.Equal(t, time.Hour*14+time.Minute*15, result)

	_, err = NewDuration(0, "invalid").Deserialize()
	assert.Contains(t, err.Error(), "HH:MM")

	_, err = NewDuration(0, "as:df").Deserialize()
	assert.Contains(t, err.Error(), "Atoi")

	_, err = NewDuration(0, "99:00").Deserialize()
	assert.Contains(t, err.Error(), "hour")

	_, err = NewDuration(0, "12:99").Deserialize()
	assert.Contains(t, err.Error(), "minutes")
}

func TestNewDuration(t *testing.T) {
	serializer := NewDuration(0, "15:04")
	assert.Equal(t, "15:04", serializer.timeStr)
	assert.Equal(t, time.Duration(0), serializer.duration)
}
