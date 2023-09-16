package repositories

// ErrorLogger logs errors.
type ErrorLogger interface {
	Errorf(format string, args ...interface{})
}
