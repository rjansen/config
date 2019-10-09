package migi

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/rjansen/abend"
	"github.com/rjansen/migi/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	testOption struct {
		name  string
		value interface{}
	}

	testOptions struct {
		name     string
		options  []testOption
		sources  []Source
		expected testOptionsExpected
	}

	testOptionsExpected struct {
		loadError error
		options   map[string]interface{}
	}
)

func TestOptions(t *testing.T) {
	tests := []testOptions{
		{
			name: "when all requested var exists",
			options: []testOption{
				{name: "string_key", value: testutils.StringPointer("")},
				{name: "int_key", value: testutils.IntPointer(0)},
				{name: "float_key", value: testutils.FloatPointer(0.0)},
				{name: "bool_key", value: testutils.BoolPointer(false)},
				{name: "time_key", value: testutils.TimePointer(time.Time{})},
				{name: "duration_key", value: testutils.DurationPointer(time.Duration(0))},
				{name: "default_string_key", value: testutils.StringPointer("default_string_value")},
				{name: "default_int_key", value: testutils.IntPointer(976)},
				{name: "default_float_key", value: testutils.FloatPointer(455.55)},
				{name: "default_bool_key", value: testutils.BoolPointer(true)},
				{name: "default_time_key", value: testutils.TimePointer(testutils.NewTime(t, "2006-01-02", "2012-05-30"))},
				{name: "default_duration_key", value: testutils.DurationPointer(time.Second * 60)},
			},
			sources: []Source{
				&mockSource{
					options: map[string]interface{}{
						"string_key":   "string_value",
						"int_key":      333,
						"float_key":    float32(333.33),
						"bool_key":     true,
						"time_key":     testutils.NewTime(t, "2006-01-02", "1999-10-05"),
						"duration_key": time.Minute * 5,
					},
				},
			},
			expected: testOptionsExpected{
				options: map[string]interface{}{
					"string_key":           "string_value",
					"int_key":              333,
					"float_key":            float32(333.33),
					"bool_key":             true,
					"time_key":             testutils.NewTime(t, "2006-01-02", "1999-10-05"),
					"duration_key":         time.Minute * 5,
					"default_string_key":   "default_string_value",
					"default_int_key":      976,
					"default_float_key":    455.55,
					"default_bool_key":     true,
					"default_time_key":     testutils.NewTime(t, "2006-01-02", "2012-05-30"),
					"default_duration_key": time.Second * 60,
				},
			},
		},
		{
			name: "when load raises error",
			sources: []Source{
				&mockSource{
					loadError: errors.New("mock_error_1"),
				},
				&mockSource{
					loadError: errors.New("mock_error_2"),
				},
				&mockSource{
					loadError: errors.New("mock_error_3"),
				},
			},
			expected: testOptionsExpected{
				loadError: errors.New("errors.List{mock_error_1, mock_error_2, mock_error_3}"),
			},
		},
		{
			name: "when get option raises error",
			options: []testOption{
				{name: "string_key", value: testutils.StringPointer("")},
				{name: "int_key", value: testutils.IntPointer(0)},
				{name: "invalid_type_key", value: testutils.IntPointer(0)},
				{name: "float_key", value: testutils.FloatPointer(0.0)},
				{name: "bool_key", value: testutils.BoolPointer(false)},
				{name: "time_key", value: testutils.TimePointer(time.Time{})},
				{name: "duration_key", value: testutils.DurationPointer(time.Duration(0))},
			},
			sources: []Source{
				&mockSource{
					options: map[string]interface{}{
						"string_key":       errors.New("mock_err_string_key"),
						"int_key":          errors.New("mock_err_int_key"),
						"invalid_type_key": "invalid_type_key",
						"float_key":        errors.New("mock_err_float_key"),
						"bool_key":         errors.New("mock_err_bool_key"),
						"time_key":         errors.New("mock_err_time_key"),
						"duration_key":     errors.New("mock_err_duration_key"),
					},
				},
			},
			expected: testOptionsExpected{
				loadError: abend.NewList(
					errors.New("mock_err_string_key"),
					errors.New("mock_err_int_key"),
					errors.New("errors.OptionInvalidType{Name='invalid_type_key', Source='string', Target='int'}"),
					errors.New("mock_err_float_key"),
					errors.New("mock_err_bool_key"),
					errors.New("mock_err_time_key"),
					errors.New("mock_err_duration_key"),
				),
			},
		},
	}
	for index, test := range tests {
		t.Run(
			testutils.TestName(t, test.name, index),
			func(t *testing.T) {
				var (
					stringOptions   = make(map[string]*string)
					intOptions      = make(map[string]*int)
					floatOptions    = make(map[string]*float32)
					boolOptions     = make(map[string]*bool)
					timeOptions     = make(map[string]*time.Time)
					durationOptions = make(map[string]*time.Duration)
					options         = NewOptions(test.sources...)
				)

				require.NotNil(t, options)

				for _, option := range test.options {
					switch value := option.value.(type) {
					case *string:
						stringOptions[option.name] = options.String(
							option.name, *value, fmt.Sprintf("the option: name='%s'", option.name),
						)
					case *int:
						intOptions[option.name] = options.Int(
							option.name, *value, fmt.Sprintf("the option: name='%s'", option.name),
						)
					case *float32:
						floatOptions[option.name] = options.Float(
							option.name, *value, fmt.Sprintf("the option: name='%s'", option.name),
						)
					case *bool:
						boolOptions[option.name] = options.Bool(
							option.name, *value, fmt.Sprintf("the option: name='%s'", option.name),
						)
					case *time.Time:
						timeOptions[option.name] = options.Time(
							option.name, *value, fmt.Sprintf("the option: name='%s'", option.name),
						)
					case *time.Duration:
						durationOptions[option.name] = options.Duration(
							option.name, *value, fmt.Sprintf("the option: name='%s'", option.name),
						)
					}
				}

				loadError := options.Load()
				if test.expected.loadError != nil {
					require.EqualError(t, loadError, test.expected.loadError.Error())
				} else {
					require.Nil(t, loadError)
				}

				for key, value := range test.expected.options {
					switch value.(type) {
					case string:
						assert.Equal(t, value, *stringOptions[key])
					case int:
						assert.Equal(t, value, *intOptions[key])
					case float32:
						assert.Equal(t, value, *floatOptions[key])
					case bool:
						assert.Equal(t, value, *boolOptions[key])
					case time.Time:
						assert.Equal(t, value, *timeOptions[key])
					case time.Duration:
						assert.Equal(t, value, *durationOptions[key])
					}
				}
			},
		)
	}
}
