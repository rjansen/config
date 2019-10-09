package migi

import (
	"time"
)

type mockSource struct {
	options   map[string]interface{}
	loadError error
}

func (m mockSource) Load() error {
	if m.loadError != nil {
		return m.loadError
	}
	return nil
}

func (m mockSource) getValue(name string) (interface{}, error) {
	value, ok := m.options[name]
	if !ok {
		return false, NewOptionNotFound(name)
	}
	err, is := value.(error)
	if is {
		return nil, err
	}
	return value, nil
}

func (m mockSource) String(name string) (string, error) {
	value, err := m.getValue(name)
	if err != nil {
		return "", err
	}
	stringValue, is := value.(string)
	if !is {
		return "", NewOptionInvalidType(name, value, "string")
	}
	return stringValue, nil
}

func (m mockSource) Int(name string) (int, error) {
	value, err := m.getValue(name)
	if err != nil {
		return 0, err
	}
	intValue, is := value.(int)
	if !is {
		return 0, NewOptionInvalidType(name, value, "int")
	}
	return intValue, nil
}

func (m mockSource) Float(name string) (float32, error) {
	value, err := m.getValue(name)
	if err != nil {
		return 0.0, err
	}
	floatValue, is := value.(float32)
	if !is {
		return 0.0, NewOptionInvalidType(name, value, "float32")
	}
	return floatValue, nil
}

func (m mockSource) Bool(name string) (bool, error) {
	value, err := m.getValue(name)
	if err != nil {
		return false, err
	}
	boolValue, is := value.(bool)
	if !is {
		return false, NewOptionInvalidType(name, value, "float32")
	}
	return boolValue, nil
}

func (m mockSource) Time(name string) (time.Time, error) {
	value, err := m.getValue(name)
	if err != nil {
		return time.Time{}, err
	}
	timeValue, is := value.(time.Time)
	if !is {
		return time.Time{}, NewOptionInvalidType(name, value, "time.Time")
	}
	return timeValue, nil
}

func (m mockSource) Duration(name string) (time.Duration, error) {
	value, err := m.getValue(name)
	if err != nil {
		return time.Duration(0), err
	}
	durationValue, is := value.(time.Duration)
	if !is {
		return time.Duration(0), NewOptionInvalidType(name, value, "time.Duration")
	}
	return durationValue, nil
}
