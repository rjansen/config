package environment

import (
	"os"
	"testing"
	"time"

	"github.com/rjansen/migi"
	"github.com/rjansen/migi/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	testSource struct {
		name         string
		setupTest    func(*testing.T, *testSource)
		tearDownTest func(*testing.T, *testSource)
		match        testSourceMatch
	}

	testSourceMatch struct {
		options map[string]interface{}
	}
)

func (e testSource) setup(t *testing.T) {
	if e.setupTest != nil {
		e.setupTest(t, &e)
	}
}

func (e testSource) tearDown(t *testing.T) {
	if e.tearDownTest != nil {
		e.tearDownTest(t, &e)
	}
}

func TestSource(t *testing.T) {
	scenarios := []testSource{
		{
			name: "load vars from environment",
			setupTest: func(t *testing.T, _ *testSource) {
				os.Setenv("string_key", "string_value")
				os.Setenv("int_key", "333")
				os.Setenv("float_key", "333.33")
				os.Setenv("bool_key", "true")
				os.Setenv("time_key", "2019-05-23T00:00:00Z")
				os.Setenv("duration_key", "5m")
			},
			tearDownTest: func(t *testing.T, _ *testSource) {
				os.Unsetenv("string_key")
				os.Unsetenv("int_key")
				os.Unsetenv("float_key")
				os.Unsetenv("bool_key")
				os.Unsetenv("time_key")
				os.Unsetenv("duration_key")
			},
			match: testSourceMatch{
				options: map[string]interface{}{
					"string_key":   "string_value",
					"int_key":      333,
					"float_key":    float32(333.33),
					"bool_key":     true,
					"time_key":     testutils.NewTime(t, time.RFC3339, "2019-05-23T00:00:00Z"),
					"duration_key": time.Minute * 5,
				},
			},
		},
	}

	for index, scenario := range scenarios {
		t.Run(
			testutils.TestName(t, scenario.name, index),
			func(t *testing.T) {
				scenario.setup(t)
				defer scenario.tearDown(t)

				source := NewSource()
				require.NotNil(t, source)
				require.Implements(t, (*migi.Source)(nil), source)
				require.NoError(t, source.Load())

				for key, value := range scenario.match.options {
					switch value.(type) {
					case string:
						v, err := source.String(key)
						assert.Nil(t, err)
						assert.Equal(t, value, v)
					case int:
						v, err := source.Int(key)
						assert.Nil(t, err)
						assert.Equal(t, value, v)
					case float32:
						v, err := source.Float(key)
						assert.Nil(t, err)
						assert.Equal(t, value, v)
					case bool:
						v, err := source.Bool(key)
						assert.Nil(t, err)
						assert.Equal(t, value, v)
					case time.Time:
						v, err := source.Time(key)
						assert.Nil(t, err)
						assert.Equal(t, value, v)
					case time.Duration:
						v, err := source.Duration(key)
						assert.Nil(t, err)
						assert.Equal(t, value, v)
					}
				}

			},
		)
	}
}
