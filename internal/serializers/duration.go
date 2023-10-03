package serializers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Duration is a duration struct.
type Duration struct {
	duration time.Duration
	timeStr  string
}

// Serialize serializes the duration.
func (s *Duration) Serialize() string {
	t := time.Now()
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	t = t.Add(s.duration)

	return t.Format("15:04")
}

// Deserialize deserializes the duration.
func (s *Duration) Deserialize() (time.Duration, error) {
	partsCount := 2
	maxHour := 23
	maxMinutes := 59
	parts := strings.Split(s.timeStr, ":")
	if len(parts) != partsCount {
		return 0, errors.New("invalid time. Format: HH:MM")
	}

	hour, err1 := strconv.Atoi(parts[0])
	minutes, err2 := strconv.Atoi(parts[1])

	if err := errors.Join(err1, err2); err != nil {
		return 0, err
	}
	if hour > maxHour {
		return 0, errors.New("invalid time. hour > 23")
	}
	if minutes > maxMinutes {
		return 0, errors.New("invalid time. minutes > 59")
	}

	return time.ParseDuration(
		fmt.Sprintf("%dh%dm", hour, minutes),
	)
}

func NewDuration(duration time.Duration, timeStr string) *Duration {
	return &Duration{
		duration: duration,
		timeStr:  timeStr,
	}
}
