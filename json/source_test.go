package environment

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/rjansen/migi"
	"github.com/rjansen/migi/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	testSource struct {
		name    string
		jsonRaw []byte
		match   testSourceMatch
	}

	testSourceMatch struct {
		loadError error
		options   map[string]interface{}
	}
)

func testName(index int, name string) string {
	return fmt.Sprintf("%d-%s", index, name)
}

func newTime(t *testing.T, str string) time.Time {
	newTime, err := time.Parse("2006-01-02", str)
	assert.Nil(t, err)
	return newTime
}

func TestEnvironmentSource(t *testing.T) {
	scenarios := []testSource{
		{
			name: "load vars from json",
			jsonRaw: []byte(`{
				"string_key": "string_value",
			 	"int_key": 333,
				"float_key": 455.55,
				"bool_key": true,
				"time_key": "2019-05-23T00:00:00Z",
				"duration_key": "5m",
				"int_string_key": "550",
				"float_string_key": "555.78",
				"bool_string_key": "true"
			}`),
			match: testSourceMatch{
				options: map[string]interface{}{
					"string_key":       "string_value",
					"int_key":          333,
					"float_key":        float32(455.55),
					"bool_key":         true,
					"time_key":         testutils.NewTime(t, time.RFC3339, "2019-05-23T00:00:00Z"),
					"duration_key":     time.Minute * 5,
					"int_string_key":   550,
					"float_string_key": float32(555.78),
					"bool_string_key":  true,
				},
			},
		},
	}

	for index, scenario := range scenarios {
		t.Run(
			testName(index, scenario.name),
			func(t *testing.T) {
				source := NewSource(bytes.NewReader(scenario.jsonRaw))
				require.NotNil(t, source)
				require.Implements(t, (*migi.Source)(nil), source)

				loadError := source.Load()
				if scenario.match.loadError != nil {
					require.EqualError(t, loadError, scenario.match.loadError.Error())
				} else {
					require.Nil(t, loadError)
				}

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
