package repositories

import (
	"time"
	"webmalc/no-more-excuses/internal/dto"
	"webmalc/no-more-excuses/internal/utils"
)

// The app repository.
type AppRepository struct {
	config *Config
	logger ErrorLogger
}

// GetApps returns the list of apps.
func (s *AppRepository) GetApps() map[string]dto.App {
	apps := make(map[string]dto.App)
	for name, data := range s.config.Apps {
		startBase, endBase, err := utils.GetDurationsFromStrRange(
			data.BaseSchedule,
		)
		if err != nil {
			s.logger.Errorf(
				"invalid time range: %s. Error: %s", data.BaseSchedule, err,
			)

			continue
		}
		app := s.newApp(name, data.Path)

		for weekdayStr, schedule := range data.Weekdays {
			start, end := s.getAppTime(schedule, startBase, endBase)
			s.addIntervalToApp(&app, weekdayStr, start, end)
		}
		apps[name] = app
	}

	return apps
}

func (s *AppRepository) newApp(name, path string) dto.App {
	return dto.App{
		Name:     name,
		Path:     path,
		Weekdays: make(map[time.Weekday]dto.AppInterval),
	}
}

func (s *AppRepository) getAppTime(
	schedule string, startBase, endBase time.Duration,
) (time.Duration, time.Duration) {
	start, end, err := utils.GetDurationsFromStrRange(schedule)
	if err != nil {
		start, end = startBase, endBase
	}

	return start, end
}

func (s *AppRepository) addIntervalToApp(
	app *dto.App, weekdayStr string, start, end time.Duration,
) {
	weekday, err := utils.GetWeekdayFromStr(weekdayStr)
	if err != nil {
		s.logger.Errorf(
			"invalid weekdays: %s. Error: %s", weekday, err,
		)
	}
	app.Weekdays[weekday] = dto.AppInterval{
		StartTime: start,
		EndTime:   end,
	}
}

// NewAppRepository creates a new app repository.
func NewAppRepository(logger ErrorLogger) *AppRepository {
	return &AppRepository{
		config: NewConfig(),
		logger: logger,
	}
}
