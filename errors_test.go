package migi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionNotFound(t *testing.T) {
	name := "my_option"
	err := NewOptionNotFound(name)

	assert.EqualError(t, err, "errors.OptionNotFound{Name='my_option'}")
}

func TestOptionInvalidOption(t *testing.T) {
	name := "my_option"
	source := 877
	target := "string"
	err := NewOptionInvalidType(name, source, target)

	assert.EqualError(t, err, "errors.OptionInvalidType{Name='my_option', Source='int', Target='string'}")
}
