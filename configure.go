package migi

import (
	"io"

	"github.com/paked/configure"
)

// NewOptions creates an options instance with the provided sources
func NewOptions(sources ...OptionsSource) Options {
	stack := make([]configure.Checker, len(sources))
	for index, source := range sources {
		stack[index] = source
	}
	return configure.New(stack...)
}

// NewEnvironmentSource creates an options source from session environment variables
func NewEnviromentSource() OptionsSource {
	return configure.NewEnvironment()
}

// NewJSONSource creates an options source from the provided json reader
func NewJSONSource(reader io.Reader) OptionsSource {
	return configure.NewJSON(
		func() (io.Reader, error) {
			return reader, nil
		},
	)
}
