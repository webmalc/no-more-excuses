package mocks

import (
	"github.com/stretchr/testify/mock"
)

// ConfigViewer mocks the object.
type ConfigViewer struct {
	mock.Mock
}

// Run is method mock.
func (s *ConfigViewer) ShowConfig() {
	s.Called()
}
