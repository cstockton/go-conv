package conv

// Builds the README.md and examples_test.go from Markdown templates.
//go:generate go test -v -args -generate

// New returns a new Value.
func New(v interface{}) Value {
	return Value{v}
}

// Converter groups all the conversion interfaces of this package.
type Converter interface {
	BoolConverter
	DurationConverter
	NumericConverter
	StringConverter
	TimeConverter
}
