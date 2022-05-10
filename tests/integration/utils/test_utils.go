package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertNilOrFail(t *testing.T, subject any) {
	if !assert.Nil(t, subject) {
		t.FailNow()
	}
}
