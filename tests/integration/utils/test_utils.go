package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func AssertNilOrFail(t *testing.T, subject any) {
	if !assert.Nil(t, subject) {
		t.FailNow()
	}
}
