package migi

import (
	"fmt"
)

type OptionNotFound struct {
	Name string
}

func (e OptionNotFound) Error() string {
	return fmt.Sprintf("errors.OptionNotFound{Name='%s'}", e.Name)
}

func NewOptionNotFound(name string) error {
	return OptionNotFound{Name: name}
}

type OptionInvalidType struct {
	Name   string
	Source interface{}
	Target string
}

func (e OptionInvalidType) Error() string {
	return fmt.Sprintf("errors.OptionInvalidType{Name='%s', Source='%T', Target='%s'}", e.Name, e.Source, e.Target)
}

func NewOptionInvalidType(name string, source interface{}, target string) error {
	return OptionInvalidType{Name: name, Source: source, Target: target}
}
