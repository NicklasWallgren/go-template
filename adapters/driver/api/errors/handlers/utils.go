package handlers

import "errors"

func AsError[T error](err error) *T {
	var expectedErrorType *T
	if errors.As(err, &expectedErrorType) { // nolint: govet
		return expectedErrorType
	}

	return nil
}
