package testutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTime(t *testing.T) {
	format := time.RFC3339
	now, err := time.Parse(format, "2019-05-23T00:00:00Z")
	assert.NoError(t, err)

	time := NewTime(t, format, now.Format(format))
	assert.Equal(t, now, time)
}
