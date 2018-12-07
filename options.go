package migi

// Options is an interface to provide access of the sofwatre configuration
type Options interface {
	Bool(string, bool, string) *bool
	BoolVar(*bool, string, bool, string)
	Int(string, int, string) *int
	IntVar(*int, string, int, string)
	String(string, string, string) *string
	StringVar(*string, string, string, string)
	Parse()
}

// OptionsSource is an interface to define how options are load and read
type OptionsSource interface {
	Setup() error
	Bool(string) (bool, error)
	Int(string) (int, error)
	String(string) (string, error)
}
