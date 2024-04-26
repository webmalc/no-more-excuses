package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Logger logs errors.
type Logger struct {
	mock.Mock
}

// Error is method mock.
func (m *Logger) Error(args ...interface{}) {
	m.Called(args...)
}

// Debugf is method mock.
func (m *Logger) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	m.Called(_ca...)
}
