package migi

import (
	"time"

	"github.com/rjansen/abend"
)

type (
	// Options is an interface wich provides access to software configuration
	Options interface {
		String(name string, defaultValue string, description string) *string
		StringVar(pointer *string, name string, defaultValue string, description string)
		Int(name string, defaultValue int, description string) *int
		IntVar(pointer *int, name string, defaultValue int, description string)
		Bool(name string, defaultValue bool, description string) *bool
		BoolVar(pointer *bool, name string, defaultValue bool, description string)
		Float(name string, defaultValue float32, description string) *float32
		FloatVar(pointer *float32, name string, defaultValue float32, description string)
		Time(name string, defaultValue time.Time, description string) *time.Time
		TimeVar(pointer *time.Time, name string, defaultValue time.Time, description string)
		Duration(name string, defaultValue time.Duration, description string) *time.Duration
		DurationVar(pointer *time.Duration, name string, defaultValue time.Duration, description string)
		Load() error
	}

	// Source is an interface to define how options are loaded
	Source interface {
		Load() error
		Bool(name string) (bool, error)
		Int(name string) (int, error)
		Float(name string) (float32, error)
		Time(name string) (time.Time, error)
		Duration(name string) (time.Duration, error)
		String(name string) (string, error)
	}

	// option is a configured value
	option struct {
		name         string
		description  string
		defaultValue interface{}
		pointer      interface{}
		setted       bool
	}

	// options is a default Options implementation
	options struct {
		register []*option
		sources  []Source
	}
)

func (o *option) scan(sources ...Source) []error {
	var errs []error
	for _, source := range sources {
		err := o.read(source)
		if err != nil {
			if _, is := err.(OptionNotFound); !is {
				errs = append(errs, err)
			}
			continue
		}
		o.setted = true
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (o *option) read(source Source) error {
	switch pointer := o.pointer.(type) {
	case *string:
		value, err := source.String(o.name)
		if err != nil {
			return err
		}
		*pointer = value
	case *int:
		value, err := source.Int(o.name)
		if err != nil {
			return err
		}
		*pointer = value
	case *float32:
		value, err := source.Float(o.name)
		if err != nil {
			return err
		}
		*pointer = value
	case *bool:
		value, err := source.Bool(o.name)
		if err != nil {
			return err
		}
		*pointer = value
	case *time.Time:
		value, err := source.Time(o.name)
		if err != nil {
			return err
		}
		*pointer = value
	case *time.Duration:
		value, err := source.Duration(o.name)
		if err != nil {
			return err
		}
		*pointer = value
	default:
		return NewOptionInvalidType(o.name, o.pointer, "[*string, *int, *float, *bool, *time.Time, *time.Duration]")
	}

	return nil
}

func (o *option) set(value interface{}) {
	switch pointer := o.pointer.(type) {
	case *string:
		v := value.(string)
		*pointer = v
	case *int:
		v := value.(int)
		*pointer = v
	case *float32:
		v := value.(float32)
		*pointer = v
	case *bool:
		v := value.(bool)
		*pointer = v
	case *time.Time:
		v := value.(time.Time)
		*pointer = v
	case *time.Duration:
		v := value.(time.Duration)
		*pointer = v
	}
}

func (o *option) setDefault() {
	o.set(o.defaultValue)
}

func (o *options) String(name string, defaultValue string, description string) *string {
	pointer := new(string)
	o.StringVar(pointer, name, defaultValue, description)

	return pointer
}

func (o *options) StringVar(pointer *string, name string, defaultValue string, description string) {
	o.register = append(o.register,
		&option{
			name:         name,
			description:  description,
			defaultValue: defaultValue,
			pointer:      pointer,
		},
	)
}

func (o *options) Int(name string, defaultValue int, description string) *int {
	pointer := new(int)
	o.IntVar(pointer, name, defaultValue, description)

	return pointer
}

func (o *options) IntVar(pointer *int, name string, defaultValue int, description string) {
	o.register = append(o.register,
		&option{
			name:         name,
			description:  description,
			defaultValue: defaultValue,
			pointer:      pointer,
		},
	)
}

func (o *options) Float(name string, defaultValue float32, description string) *float32 {
	pointer := new(float32)
	o.FloatVar(pointer, name, defaultValue, description)

	return pointer
}

func (o *options) FloatVar(pointer *float32, name string, defaultValue float32, description string) {
	o.register = append(o.register,
		&option{
			name:         name,
			description:  description,
			defaultValue: defaultValue,
			pointer:      pointer,
		},
	)
}

func (o *options) Bool(name string, defaultValue bool, description string) *bool {
	pointer := new(bool)
	o.BoolVar(pointer, name, defaultValue, description)

	return pointer
}

func (o *options) BoolVar(pointer *bool, name string, defaultValue bool, description string) {
	o.register = append(o.register,
		&option{
			name:         name,
			description:  description,
			defaultValue: defaultValue,
			pointer:      pointer,
		},
	)
}

func (o *options) Time(name string, defaultValue time.Time, description string) *time.Time {
	pointer := new(time.Time)
	o.TimeVar(pointer, name, defaultValue, description)

	return pointer
}

func (o *options) TimeVar(pointer *time.Time, name string, defaultValue time.Time, description string) {
	o.register = append(o.register,
		&option{
			name:         name,
			description:  description,
			defaultValue: defaultValue,
			pointer:      pointer,
		},
	)
}

func (o *options) Duration(name string, defaultValue time.Duration, description string) *time.Duration {
	pointer := new(time.Duration)
	o.DurationVar(pointer, name, defaultValue, description)

	return pointer
}

func (o *options) DurationVar(pointer *time.Duration, name string, defaultValue time.Duration, description string) {
	o.register = append(o.register,
		&option{
			name:         name,
			description:  description,
			defaultValue: defaultValue,
			pointer:      pointer,
		},
	)
}

func (o *options) loadSources() error {
	var errs []error
	for _, source := range o.sources {
		err := source.Load()
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return abend.NewList(errs...)
	}
	return nil
}

func (o *options) Load() error {
	err := o.loadSources()
	if err != nil {
		return err
	}

	var errs []error
	for _, option := range o.register {
		scanErrs := option.scan(o.sources...)
		if len(scanErrs) > 0 {
			errs = append(errs, scanErrs...)
			continue
		}

		if !option.setted {
			option.setDefault()
		}
	}

	if len(errs) > 0 {
		return abend.NewList(errs...)
	}

	return nil
}

// NewOptions creates an options instance with the provided sources
func NewOptions(sources ...Source) Options {
	return &options{
		sources: sources,
	}
}
