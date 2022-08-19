package errors

type DomainError struct {
	Message string
	err     error
}

func NewDomainError(message string, err error) *DomainError { // nolint: wsl
	// TODO, use option withers instead of parameters?

	return &DomainError{Message: message, err: err}
}

func (d DomainError) Error() string {
	return d.Message
}

func (d DomainError) Unwrap() error {
	return d.err
}
