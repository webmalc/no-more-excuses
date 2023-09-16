package repositories

import (
	"testing"
	"time"
	"webmalc/no-more-excuses/internal/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

func TestAppRepository_GetApps(t *testing.T) {
	logger := &mocks.ErrorLogger{}
	repo := NewAppRepository(logger)
	apps := repo.GetApps()

	assert.Equal(t, "steam", apps["steam"].Name)
	assert.Equal(t, "/usr/bin/steam", apps["steam"].Path)
	assert.Equal(t, time.Hour*8, apps["steam"].Weekdays[time.Monday].StartTime)
	assert.Equal(t, time.Hour*11, apps["steam"].Weekdays[time.Monday].EndTime)
	assert.Equal(t, time.Hour*7, apps["steam"].Weekdays[time.Tuesday].StartTime)
	assert.Equal(t, time.Hour*18, apps["steam"].Weekdays[time.Tuesday].EndTime)
	assert.Equal(t, time.Hour*0, apps["steam"].Weekdays[time.Wednesday].StartTime)
	assert.Equal(
		t,
		time.Hour*23+time.Minute*59,
		apps["steam"].Weekdays[time.Wednesday].EndTime,
	)

	assert.Equal(t, "game", apps["game"].Name)
	assert.Equal(t, "/home/bob/game", apps["game"].Path)
	assert.Equal(t, time.Hour*7, apps["game"].Weekdays[time.Monday].StartTime)
	assert.Equal(t, time.Hour*10, apps["game"].Weekdays[time.Monday].EndTime)
	assert.Equal(t, time.Hour*10, apps["game"].Weekdays[time.Tuesday].StartTime)
	assert.Equal(t, time.Hour*14, apps["game"].Weekdays[time.Tuesday].EndTime)
	assert.Equal(t, time.Hour*9, apps["game"].Weekdays[time.Saturday].StartTime)
	assert.Equal(t, time.Hour*15, apps["game"].Weekdays[time.Saturday].EndTime)
	assert.NotContains(t, apps["game"].Weekdays, time.Sunday)
	assert.NotContains(t, apps["game"].Weekdays, time.Wednesday)
}

func TestNewAppRepository(t *testing.T) {
	logger := &mocks.ErrorLogger{}
	repo := NewAppRepository(logger)
	assert.Equal(t, logger, repo.logger)
	assert.NotNil(t, repo.config)
}
