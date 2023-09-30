package cmd

import (
	"os"
	"testing"

	"webmalc/no-more-excuses/cmd/mocks"
	"webmalc/no-more-excuses/common/test"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.ConfigViewer{}
	c := &mocks.ContextRunner{}
	cr := NewCommandRouter(m, c, r)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	l := &mocks.ErrorLogger{}
	r := &mocks.ConfigViewer{}
	c := &mocks.ContextRunner{}
	cr := NewCommandRouter(l, c, r)
	assert.Equal(t, l, cr.logger)
	assert.Equal(t, c, cr.serverRunner)
	assert.Equal(t, r, cr.configViewer)
	assert.NotNil(t, cr.rootCmd)
}

func TestCommandRouter_server(t *testing.T) {
	configViewer := &mocks.ConfigViewer{}
	server := &mocks.ContextRunner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, server, configViewer)
	server.On("Run", mock.Anything).Return(nil).Once()
	cr.server(&cobra.Command{}, []string{})
	server.AssertExpectations(t)
}

func TestCommandRouter_configViewer(t *testing.T) {
	configViewer := &mocks.ConfigViewer{}
	server := &mocks.ContextRunner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, server, configViewer)
	configViewer.On("ShowConfig", mock.Anything).Return(nil).Once()
	cr.configShow(&cobra.Command{}, []string{})
	configViewer.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
