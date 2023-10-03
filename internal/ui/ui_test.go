package ui

import (
	"testing"
	"webmalc/no-more-excuses/common/logger"
	"webmalc/no-more-excuses/common/test"
	"webmalc/no-more-excuses/internal/repositories"
	"webmalc/no-more-excuses/internal/ui/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUI_ShowConfig(t *testing.T) {
	ui := &UI{}
	m := &mocks.UI{}
	ui.ui = m
	m.On("ShowConfig").Return(nil)
	ui.ShowConfig()
	m.AssertExpectations(t)
}

func TestNewUI(t *testing.T) {
	log := logger.NewLogger()
	appRepo := repositories.NewAppRepository(log)
	ui := NewUI(appRepo)
	assert.NotNil(t, ui.ui)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
