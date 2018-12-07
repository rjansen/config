package migi

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type testNewOptions struct {
	Scenario
	setupFunc    func(*testing.T, *testNewOptions)
	tearDownFunc func(*testing.T, testNewOptions)
	options      map[string]string //TODO: Change to a interface map and test bool and int options
	sources      []OptionsSource
}

func (scenario testNewOptions) Setup(t *testing.T) {
	scenario.setupFunc(t, &scenario)
}

func (scenario testNewOptions) TearDown(t *testing.T) {
	scenario.tearDownFunc(t, scenario)
}

func newTestNewOptions(
	name string,
	setup func(*testing.T, *testNewOptions), tearDown func(*testing.T, testNewOptions),
	options map[string]string, sources ...OptionsSource,
) testNewOptions {
	return testNewOptions{
		Scenario: Scenario{
			Name: name,
		},
		setupFunc:    setup,
		tearDownFunc: tearDown,
		options:      options,
		sources:      sources,
	}
}

func TestNewOptions(test *testing.T) {
	scenarios := []testNewOptions{
		newTestNewOptions(
			"with environment source options",
			func(t *testing.T, scenario *testNewOptions) {
				for key, value := range scenario.options {
					os.Setenv(strings.ToUpper(key), value)
				}
			},
			func(t *testing.T, scenario testNewOptions) {
				for key, _ := range scenario.options {
					os.Unsetenv(strings.ToUpper(key))
				}
			},
			map[string]string{
				"key1":    "value1",
				"key2":    "value2",
				"keynnnn": "valuennn",
				"keyint":  "10099",
				"keybool": "true",
			},
			NewEnviromentSource(),
		),
		newTestNewOptions(
			"with json source options",
			func(t *testing.T, scenario *testNewOptions) {},
			func(t *testing.T, scenario testNewOptions) {},
			map[string]string{
				"jsonkey1":               "value1",
				"json.key2":              "value2",
				"json-keynnnn":           "valuennn",
				"json_keyint":            "10099",
				"json.json.json.keybool": "true",
			},
			NewJSONSource(
				strings.NewReader(`
					{
						"jsonkey1": "value1",
						"json.key2":    "value2",
						"json-keynnnn": "valuennn",
						"json_keyint":  "10099",
						"json.json.json.keybool": "true"
					}
				`),
			),
		),
	}
	for index, scenario := range scenarios {
		test.Run(
			fmt.Sprintf("[%d]-%s", index, scenario.Name),
			func(t *testing.T) {
				scenario.Setup(t)
				defer scenario.TearDown(t)
				options := NewOptions(scenario.sources...)
				optionValueMap := make(map[string]*string)
				for key, value := range scenario.options {
					optionValueMap[key] = options.String(
						key, "", fmt.Sprintf("invalid_option: %s=%s", key, value),
					)
				}
				options.Parse()
				for key, value := range scenario.options {
					optionValue := optionValueMap[key]
					if optionValue == nil || *optionValue == "" {
						t.Errorf("error_assert: test='%s' message=error_expected_option: expected='%s=%s' got='%v'",
							scenario.Name, key, value, optionValue,
						)
					}
				}
			},
		)
	}

}
