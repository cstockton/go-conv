package conv_test

import (
	"fmt"

	"github.com/cstockton/go-conv"
)

func Example() {

	// Conversions may be done with package functions or the Value struct.
	fmt.Printf("Bool: %#v\n", conv.Bool("true"))
	fmt.Printf("Bool: %#v == %#v\n",
		conv.Value{"true"}.Bool(), conv.Bool("true")) // Identical
	fmt.Printf("Duration: %v\n", conv.Duration("-5ns"))  // -5ns
	fmt.Printf("Duration: %v\n", conv.Duration("+10ns")) // 10ns
	fmt.Printf("Float64(): %v\n", conv.Float64("-1.23"))
	fmt.Printf("Int64(): %v\n", conv.Int64("-1.23"))
	fmt.Printf("Uint64(): %v\n", conv.Uint64("-1.23"))

	// Output:
	// Bool: true
	// Bool: true == true
	// Duration: -5ns
	// Duration: 10ns
	// Float64(): -1.23
	// Int64(): -1
	// Uint64(): 0
}

func ExampleNew() {
	// conv.New returns a conv.Value. conv.New("Foo") returns a value equivalent
	// to conv.Value{"Foo"}.
	v := conv.New("1h")
	fmt.Printf("Duration: %v\n", v.Duration())

	// Output:
	// Duration: 1h0m0s
}

func ExampleValue() {
	v := conv.Value{"12345.6789"}
	fmt.Printf("Complex64: %v\n", v.Complex64())
	fmt.Printf("Duration: %v\n", v.Duration())

	// Output:
	// Complex64: (12345.679+0i)
	// Duration: 12.345µs
}

func ExampleValue_Flatten() {
	v := conv.Value{conv.Value{conv.Value{conv.Value{"root"}}}}
	fmt.Printf("%#v\n", v)

	v = v.Flatten()
	fmt.Printf("%#v\n", v)

	// Output:
	// conv.Value{V:conv.Value{V:conv.Value{V:conv.Value{V:"root"}}}}
	// conv.Value{V:"root"}
}
