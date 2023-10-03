package serializers

import (
	"errors"
	"strings"
	"time"
)

// Weekday is a weekday struct.
type Weekday struct {
	weekday  string
	weekdays map[string]int
}

// Deserialize deserializes the weekday.
func (s *Weekday) Deserialize() (time.Weekday, error) {
	weekday := strings.ToLower(s.weekday)
	if _, ok := s.weekdays[weekday]; !ok {
		return 0, errors.New("invalid weekday")
	}

	return time.Weekday(s.weekdays[weekday]), nil
}

// NewWeekday returns a new weekday.
func NewWeekday(weekday string) *Weekday {
	return &Weekday{
		weekday: weekday,
		weekdays: map[string]int{
			"sun": 0,
			"mon": 1,
			"tue": 2,
			"wen": 3,
			"thu": 4,
			"fri": 5,
			"sat": 6,
		},
	}
}
