package cmd

import (
	"os"
	"strings"
)

type Reader interface {
	// Find an option within the input that corresponds with the given name.
	// If an option is not found false is returned. Options are values that
	// don't have an explicit value, they are a state that is either true or
	// false.
	Option(name string) bool

	// Find an argument within the input that corresponds with the given name.
	// If an argument is not found false will be returned. Arguments are named
	// values.
	Argument(name string) (string, bool)

	// Fetch all arguments.
	Arguments() map[string]string

	// Find a value at the given position. If a value does not exist at the
	// given position then false is returned.
	Value(index int) (string, bool)

	// Fetch all values.
	Values() []string
}

type ArgumentReader struct {}

func (r ArgumentReader) Option(name string) bool {
	for _, arg := range os.Args[1:] {
		formatted := strings.Replace(arg, "-", "", 1)
		if formatted == name {
			return true
		}
	}
	return false
}

func (r ArgumentReader) Arguments() map[string]string {
	arguments := make(map[string]string)
	for _, arg := range os.Args[1:] {
		formatted := strings.Replace(arg, "--", "", 1)
		segments := strings.Split(formatted, "=")

		if len(segments) == 2 {
			arguments[segments[0]] = segments[1]
		}
	}
	return arguments
}

func (r ArgumentReader) Argument(name string) (string, bool) {
	arguments := r.Arguments()
	if val, ok := arguments[name]; ok {
		return val, true
	}
	return "", false
}

func (r ArgumentReader) Value(index int) (string, bool) {
	values := r.Values()
	if len(values) > index {
		return values[index], true
	}
	return "", false
}

func (r ArgumentReader) Values() []string {
	values := make([]string, 0)
	for _, value := range os.Args[2:] {
		if !strings.Contains(value, "-") {
			values = append(values, value)
		}
	}
	return values
}
