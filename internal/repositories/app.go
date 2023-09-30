package repositories

import (
	"time"
	"webmalc/no-more-excuses/internal/dto"
	"webmalc/no-more-excuses/internal/serializers"
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
		serializer := serializers.NewDurationRange(0, 0, data.BaseSchedule)
		startBase, endBase, err := serializer.Deserialize()
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
	serializer := serializers.NewDurationRange(0, 0, schedule)
	start, end, err := serializer.Deserialize()
	if err != nil {
		start, end = startBase, endBase
	}

	return start, end
}

func (s *AppRepository) addIntervalToApp(
	app *dto.App, weekdayStr string, start, end time.Duration,
) {
	serializer := serializers.NewWeekday(weekdayStr)
	weekday, err := serializer.Deserialize()
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
