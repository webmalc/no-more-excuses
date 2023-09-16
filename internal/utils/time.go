package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GetWeekdayNumber returns the weekday from string. Format: Mon, Tue, Wed..
func GetWeekdayFromStr(weekday string) (time.Weekday, error) {
	weekday = strings.ToLower(weekday)
	weekdays := map[string]int{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wen": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}
	if _, ok := weekdays[weekday]; !ok {
		return 0, errors.New("invalid weekday")
	}

	return time.Weekday(weekdays[weekday]), nil
}

// Get durations from from string range [HH:MM-HH:MM].
func GetDurationsFromStrRange(timeRange string) (
	time.Duration, time.Duration, error,
) {
	partsCount := 2
	timeParts := strings.Split(timeRange, "-")
	if len(timeParts) != partsCount {
		return 0, 0, errors.New("invalid time range. Format: HH:MM-HH:MM")
	}
	begin, err1 := GetDurationsFromStr(timeParts[0])
	end, err2 := GetDurationsFromStr(timeParts[1])
	if err := errors.Join(err1, err2); err != nil {
		return 0, 0, err
	}
	if begin >= end {
		return 0, 0, errors.New("invalid time range. begin time > end time")
	}

	return begin, end, nil
}

// GetDurationsFromStr returns the durations from string HH:MM.
func GetDurationsFromStr(timeStr string) (time.Duration, error) {
	partsCount := 2
	maxHour := 23
	maxMinutes := 59
	parts := strings.Split(timeStr, ":")
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
