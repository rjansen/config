package migi

import (
	"fmt"
	"testing"

	"github.com/rjansen/yggdrasil"
)

type testRegister struct {
	Scenario
	options Options
	err     error
}

func newTestRegister(name string, options Options, err error) testRegister {
	return testRegister{
		Scenario: Scenario{
			Name: name,
		},
		options: options,
		err:     err,
	}
}

func TestRegister(test *testing.T) {
	scenarios := []testRegister{
		newTestRegister(
			"Register the Options reference",
			NewOptions(),
			nil,
		),
		newTestRegister(
			"Register a nil options reference",
			nil,
			nil,
		),
	}

	for index, scenario := range scenarios {
		test.Run(
			fmt.Sprintf("[%d]-%s", index, scenario.Name),
			func(t *testing.T) {
				roots := yggdrasil.NewRoots()
				err := Register(&roots, scenario.options)
				if err != scenario.err {
					t.Errorf("err_assert: assert=errOptionsRegister expected=%+v got=%+v", scenario.err, err)
				}
				tree := roots.NewTreeDefault()
				options, err := tree.Reference(optionsPath)
				if err != nil {
					t.Errorf("err_assert: assert=errTreeReference expected=nil got=%+v", err)
				}
				if options != scenario.options {
					t.Errorf("err_assert: assert=errTreeReferenceValue expected=%+v got=%+v", scenario.options, options)
				}
			},
		)
	}
}

type testReference struct {
	Scenario
	references map[yggdrasil.Path]yggdrasil.Reference
	tree       yggdrasil.Tree
	err        error
}

func newTestReference(name string, references map[yggdrasil.Path]yggdrasil.Reference, err error) testReference {
	return testReference{
		Scenario: Scenario{
			Name: name,
		},
		references: references,
		err:        err,
	}
}

func (scenario *testReference) Setup(t *testing.T) {
	roots := yggdrasil.NewRoots()
	for path, reference := range scenario.references {
		err := roots.Register(path, reference)
		if err != nil {
			t.Errorf("err_assert: assert=setup.TestReference expected=nil got=%+v", err)
		}
	}
	scenario.tree = roots.NewTreeDefault()
}

func TestReference(test *testing.T) {
	scenarios := []testReference{
		newTestReference(
			"Access the options Reference",
			map[yggdrasil.Path]yggdrasil.Reference{
				optionsPath: yggdrasil.NewReference(NewOptions()),
			},
			nil,
		),
		newTestReference(
			"Access the a nil options Reference",
			map[yggdrasil.Path]yggdrasil.Reference{
				optionsPath: yggdrasil.NewReference(nil),
			},
			nil,
		),
		newTestReference(
			"When options was not register returns path not found",
			nil,
			yggdrasil.ErrPathNotFound,
		),
		newTestReference(
			"When a invalid options was register returns ...error",
			map[yggdrasil.Path]yggdrasil.Reference{
				optionsPath: yggdrasil.NewReference(new(struct{})),
			},
			ErrInvalidReference,
		),
	}

	for index, scenario := range scenarios {
		test.Run(
			fmt.Sprintf("[%d]-%s", index, scenario.Name),
			func(t *testing.T) {
				scenario.Setup(t)

				_, err := Reference(scenario.tree)
				if err != scenario.err {
					t.Errorf("err_assert: assert=errOptionsReference expected=%+v got=%+v", scenario.err, err)
				}

				mustTestFunc := func() {
					defer func() {
						err := recover()
						if err != scenario.err {
							t.Errorf("err_assert: assert=errOptionsMustReference expected=%+v got=%+v", scenario.err, err)
						}
					}()
					_ = MustReference(scenario.tree)
				}
				mustTestFunc()
			},
		)
	}
}
