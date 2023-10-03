package serializers

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// DurationRange is a duration rage struct.
type DurationRange struct {
	start     *Duration
	end       *Duration
	timeRange string
}

// Serialize serializes the duration.
func (s *DurationRange) Serialize() string {
	return fmt.Sprintf("%s-%s", s.start.Serialize(), s.end.Serialize())
}

// Deserialize deserializes the duration.
func (s *DurationRange) Deserialize() (
	time.Duration, time.Duration, error,
) {
	partsCount := 2
	timeParts := strings.Split(s.timeRange, "-")
	if len(timeParts) != partsCount {
		return 0, 0, errors.New("invalid time range. Format: HH:MM-HH:MM")
	}
	begin, err1 := NewDuration(0, timeParts[0]).Deserialize()
	end, err2 := NewDuration(0, timeParts[1]).Deserialize()
	if err := errors.Join(err1, err2); err != nil {
		return 0, 0, err
	}
	if begin >= end {
		return 0, 0, errors.New("invalid time range. begin time > end time")
	}

	return begin, end, nil
}

// NewDurationRange creates a new duration range.
func NewDurationRange(
	start, end time.Duration, timeRange string,
) *DurationRange {
	return &DurationRange{
		start:     NewDuration(start, ""),
		end:       NewDuration(end, ""),
		timeRange: timeRange,
	}
}
