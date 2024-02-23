package errors

func New(text string) error {
	return &ValidationError{text}
}

type ValidationError struct {
	s string
}

func (e *ValidationError) Error() string {
	return e.s
}
