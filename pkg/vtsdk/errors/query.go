package virusTotalErrors

import "fmt"

type QueryError struct {
	target  string
	wrapped error
}

func NewQueryError(target string) *QueryError {
	return &QueryError{
		target: target,
	}
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("error querying for %s on VirusTotal: %s", e.target, e.wrapped)
}

func (e *QueryError) Wrap(wrapped error) *QueryError {
	e.wrapped = wrapped
	return e
}
