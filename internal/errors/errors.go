package errors

import (
	"errors"
	"fmt"
	"strings"
)

var New = errors.New

type List []error

func NewList(errs ...error) List {
	return List(errs)
}

func (list List) Error() string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return "errors.List{" + list[0].Error() + "}"
	}
	var builder strings.Builder
	builder.WriteString("errors.List{" + list[0].Error())
	for _, err := range list[1:] {
		builder.WriteString(", ")
		builder.WriteString(err.Error())
	}
	builder.WriteString("}")
	return builder.String()
}

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
