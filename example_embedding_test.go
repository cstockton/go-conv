package conv_test

import (
	"fmt"

	"github.com/cstockton/go-conv"
)

// V will return MyValue which embeds conv.Value.
func V(value interface{}) MyValue {
	return MyValue{conv.Value{value}}
}

type MyString string

type MyValue struct{ conv.Value }

// Int behaves like conv.Int() except when the underlying type is a MyString the
// length is returned instead.
func (v MyValue) Int() int {
	if s, ok := v.Indirect().(MyString); ok {
		return len(s)
	}
	return v.Value.Int()
}

func Example_embedding() {
	fmt.Println("Int():", V(`4321`).Int())
	fmt.Println("Int():", V(MyString(`4321`)).Int())

	// Output:
	// Int(): 4321
	// Int(): 4
}
