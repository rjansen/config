package environment

import (
	"os"
	"strconv"
	"time"

	"github.com/rjansen/migi"
)

type source struct{}

func (e *source) Load() error {
	return nil
}

func (e *source) lookup(name string) (string, error) {
	value, ok := os.LookupEnv(name)
	if !ok {
		return "", migi.NewOptionNotFound(name)
	}

	return value, nil
}

func (e *source) String(name string) (string, error) {
	value, err := e.lookup(name)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (e *source) Int(name string) (int, error) {
	envValue, err := e.lookup(name)
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseInt(envValue, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(value), nil
}

func (e *source) Float(name string) (float32, error) {
	envValue, err := e.lookup(name)
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(envValue, 32)
	if err != nil {
		return 0, err
	}

	return float32(value), nil
}

func (e *source) Bool(name string) (bool, error) {
	envValue, err := e.lookup(name)
	if err != nil {
		return false, err
	}

	value, err := strconv.ParseBool(envValue)
	if err != nil {
		return false, err
	}

	return value, nil
}

func (e *source) Time(name string) (time.Time, error) {
	envValue, err := e.lookup(name)
	if err != nil {
		return time.Time{}, err
	}

	value, err := time.Parse(time.RFC3339, envValue)
	if err != nil {
		return time.Time{}, err
	}

	return value, nil
}

func (e *source) Duration(name string) (time.Duration, error) {
	envValue, err := e.lookup(name)
	if err != nil {
		return time.Duration(0), err
	}

	value, err := time.ParseDuration(envValue)
	if err != nil {
		return time.Duration(0), err
	}

	return value, nil
}

func NewSource() *source {
	return new(source)
}
