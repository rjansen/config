package testutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringPointer(t *testing.T) {
	value := "my_string"
	pointer := StringPointer(value)
	assert.Equal(t, value, *pointer)
}

func TestIntPointer(t *testing.T) {
	value := 999
	pointer := IntPointer(value)
	assert.Equal(t, value, *pointer)
}

func TestFloatPointer(t *testing.T) {
	value := float32(999.87)
	pointer := FloatPointer(value)
	assert.Equal(t, value, *pointer)
}

func TestBoolPointer(t *testing.T) {
	value := true
	pointer := BoolPointer(value)
	assert.Equal(t, value, *pointer)
}

func TestTimePointer(t *testing.T) {
	value := time.Now().UTC()
	pointer := TimePointer(value)
	assert.Equal(t, value, *pointer)
}

func TestDurationPointer(t *testing.T) {
	value := time.Second * 33
	pointer := DurationPointer(value)
	assert.Equal(t, value, *pointer)
}
