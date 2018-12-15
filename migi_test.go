package migi

import "testing"

type Scenario struct {
	Name string
}

type NoSetup struct{}

func (NoSetup) Setup(t *testing.T) error {
	return nil
}

type NoTearDown struct{}

func (NoTearDown) TearDown(t *testing.T) error {
	return nil
}

type NoDependency struct {
	Scenario
	NoSetup
	NoTearDown
}
