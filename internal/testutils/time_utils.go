package testutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func NewTime(t *testing.T, layout, str string) time.Time {
	newTime, err := time.Parse(layout, str)
	assert.Nil(t, err)
	return newTime
}
