package domain

import "fmt"

type Severity string

const (
	SEVERITY_SUCCESS = "success"
	SEVERITY_WARNING = "warning"
	SEVERITY_ERROR   = "error"
)

type Error struct {
	Severity Severity
	Message  string
	Details  error
}

func (e *Error) Error() string {
	if e.Details != nil {
		return fmt.Sprintf("%v", e.Details)
	}
	return ""
}

func NewError(err error, message string, severity Severity) *Error {
	return &Error{
		Severity: severity,
		Message:  message,
		Details:  err,
	}
}
