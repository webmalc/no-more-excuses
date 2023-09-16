package dto

import "time"

// The Application DTO object.
type App struct {
	Name     string
	Path     string
	Weekdays map[time.Weekday]AppInterval
}

// The Application Interval DTO object.
type AppInterval struct {
	// Duration for the start of the day 00:00. E.g. 09:30 == 9h30m
	StartTime time.Duration
	EndTime   time.Duration
}
