package errors

import (
	"errors"
	"testing"

	"github.com/rjansen/migi/internal/testutils"
	"github.com/stretchr/testify/assert"
)

type (
	testList struct {
		name  string
		errs  []error
		match testListMatch
	}

	testListMatch struct {
		message string
	}
)

func TestList(t *testing.T) {
	scenarios := []testList{
		{
			name: "when has no errors",
			errs: []error{},
			match: testListMatch{
				message: "",
			},
		},
		{
			name: "when has only one error",
			errs: []error{
				errors.New("mock_error_1"),
			},
			match: testListMatch{
				message: "errors.List{mock_error_1}",
			},
		},
		{
			name: "when has many errors",
			errs: []error{
				errors.New("mock_error_1"),
				errors.New("mock_error_2"),
				errors.New("mock_error_3"),
				errors.New("mock_error_4"),
				errors.New("mock_error_5"),
			},
			match: testListMatch{
				message: "errors.List{mock_error_1, mock_error_2, mock_error_3, mock_error_4, mock_error_5}",
			},
		},
	}

	for index, scenario := range scenarios {
		t.Run(
			testutils.TestName(t, scenario.name, index),
			func(t *testing.T) {
				list := NewList(scenario.errs...)
				assert.EqualError(t, list, scenario.match.message)
			},
		)
	}
}

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
