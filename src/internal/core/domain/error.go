package domain

type Severity string

const (
	SEVERITY_SUCCESS = "success"
	SEVERITY_WARNING = "warning"
	SEVERITY_ERROR   = "error"
)

type Error struct {
	Severity Severity
	Message  string
	Error    error
}
