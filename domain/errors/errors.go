package errors

type DomainError struct {
	Message string
	err     error
}

// TODO, pass options?

func NewDomainError(message string, err error) *DomainError {
	return &DomainError{Message: message, err: err}
}

func (d DomainError) Error() string {
	return d.Message // TODO
}

func (d DomainError) Unwrap() error {
	return d.err
}
