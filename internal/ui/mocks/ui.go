package mocks

import (
	"github.com/stretchr/testify/mock"
)

// UI mocks the object.
type UI struct {
	mock.Mock
}

// Run is method mock.
func (s *UI) ShowConfig() {
	s.Called()
}
