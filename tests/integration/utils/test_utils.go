package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertNilOrFail(t *testing.T, subject any) {
	t.Helper()

	if !assert.Nil(t, subject) {
		t.FailNow()
	}
}
