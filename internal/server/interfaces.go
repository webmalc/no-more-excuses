package server

// Logger logs errors.
type Logger interface {
	Error(args ...interface{})
	Debugf(format string, args ...interface{})
}
