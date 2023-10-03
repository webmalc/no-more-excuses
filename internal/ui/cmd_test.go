package ui

import (
	"testing"
	"webmalc/no-more-excuses/common/logger"
	"webmalc/no-more-excuses/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestNewCmd(t *testing.T) {
	log := logger.NewLogger()
	appRepo := repositories.NewAppRepository(log)
	apps := appRepo.GetApps()
	cmd := NewCmd(apps)
	assert.Equal(t, apps, cmd.apps)
	assert.NotNil(t, cmd.model)
}
