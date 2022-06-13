package utils

import (
	"errors"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	dir, _            = Dirname()
	TestDirectoryRoot = dir + "/../../"
)

func ValueFromOrFail[T any](t *testing.T, value T, err error) T {
	t.Helper()

	AssertNilOrFail(t, err)

	return value
}

func ValueFromSupplierOrFail[T any](t *testing.T, supplier func() (T, error)) T {
	t.Helper()

	value, err := supplier()

	AssertNilOrFail(t, err)

	return value
}

func SuccessOrFailNow[T any](t *testing.T, tester func() (T, error)) {
	t.Helper()

	_, err := tester()

	AssertNilOrFail(t, err)
}

func Filename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename") // nolint:goerr113
	}

	return filename, nil
}

func Dirname() (string, error) {
	filename, err := Filename()
	if err != nil {
		return "", err
	}

	return filepath.Dir(filename), nil
}
