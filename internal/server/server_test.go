package server

import (
	"context"
	"testing"
	"time"
	"webmalc/no-more-excuses/internal/server/mocks"

	"github.com/stretchr/testify/assert"
)

// Must run the server.
func TestServer_Run(t *testing.T) {
	l := &mocks.Logger{}
	s := NewServer(l)

	ctx, cancel := context.WithCancel(context.Background())
	l.On("Debugf", "server: run the blocker").Return(nil)
	l.On("Error", "server: cleanup and shutdown").Return(nil)
	go s.Run(ctx)
	time.Sleep(time.Millisecond)
	cancel()
	time.Sleep(time.Millisecond)
	l.AssertExpectations(t)
}

// Must return the server.
func TestNewServer(t *testing.T) {
	l := &mocks.Logger{}
	s := NewServer(l)
	assert.NotNil(t, s)
	assert.Equal(t, l, s.logger)
	assert.Equal(t, float64(2), s.config.IntervalDuration.Seconds())
}
