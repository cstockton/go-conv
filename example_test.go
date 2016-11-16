package conv_test

import (
	"fmt"
	"log"
	"sort"
	"time"

	conv "github.com/cstockton/go-conv"
)

func Example() {

	// All top level conversion functions discard errors, returning the types zero
	// value instead.
	fmt.Println(conv.Time("bad time string")) // time.Time{}
	fmt.Println(conv.Time("Monday, 02 Jan 2006 15:04:05 -0700"))

	// Conversions are allowed as long as the underlying type is convertable, for
	// example:
	type MyString string
	fmt.Println(conv.Int(MyString(`123`))) // 123

	// Pointers will be dereferenced when appropriate.
	s := `123`
	fmt.Println(conv.Int(&s)) // 123

	// If you would like to know if errors occur you may use a conv.Conv. It is
	// safe for use by multiple Go routines and needs no initialization.
	var c conv.Conv
	i, err := c.Int(`Foo`)
	fmt.Printf("Got %v because of err: %v", i, err)
	// Got 0 because of err: cannot convert "Foo" (type string) to int

	// Output:
	// 0001-01-01 00:00:00 +0000 UTC
	// 2006-01-02 15:04:05 -0700 MST
	// 123
	// 123
	// Got 0 because of err: cannot convert "Foo" (type string) to int
}

// String conversion from any values outside the cases below will simply be the
// result of calling fmt.Sprintf("%v", value), meaning it can not fail. An error
// is still provided and you should check it to be future proof.
func Example_strings() {

	// String conversion from other string values will be returned without
	// modification.
	fmt.Println(conv.String(`Foo`))

	// As a special case []byte will also be returned after a Go string conversion
	// is applied.
	fmt.Println(conv.String([]byte(`Foo`)))

	// String conversion from types that do not have a valid conversion path will
	// still have sane string conversion for troubleshooting.
	fmt.Println(conv.String(struct{ msg string }{"Foo"}))

	// Output:
	// Foo
	// Foo
	// {Foo}
}

func Example_bools() {

	// Bool conversion from other bool values will be returned without
	// modification.
	fmt.Println(conv.Bool(true), conv.Bool(false))

	// Bool conversion from strings consider the following values true:
	//   "t", "T", "true", "True", "TRUE",
	// 	 "y", "Y", "yes", "Yes", "YES", "1"
	//
	// It considers the following values false:
	//   "f", "F", "false", "False", "FALSE",
	//   "n", "N", "no", "No", "NO", "0"
	fmt.Println(conv.Bool("T"), conv.Bool("False"))

	// Bool conversion from other supported types will return true unless it is
	// the zero value for the given type.
	fmt.Println(conv.Bool(int64(123)), conv.Bool(int64(0)))
	fmt.Println(conv.Bool(time.Duration(123)), conv.Bool(time.Duration(0)))
	fmt.Println(conv.Bool(time.Now()), conv.Bool(time.Time{}))

	// All other types will return false.
	fmt.Println(conv.Bool(struct{ string }{""}))

	// Output:
	// true false
	// true false
	// true false
	// true false
	// true false
	// false
}

// Duration conversion from other time.Duration values will be returned without modification.
func Example_durations() {
	fmt.Println(conv.Duration(time.Duration(time.Second))) // 1s
	fmt.Println(conv.Duration("1h30m"))                    // 1h30m0s
	fmt.Println(conv.Duration("12.15"))                    // 12.15s

	// Output:
	// 1s
	// 1h30m0s
	// 12.15s
}

// Numeric conversion from other numeric values of an identical type will be
// returned without modification. Numeric conversions deviate slighty from Go
// when dealing with under/over flow. When performing a conversion operation
// that would overflow, we instead assign the maximum value for the target type.
// Similarly, conversions that would underflow are assigned the minimun value
// for that type, meaning unsigned integers are given zero values isntead of
// spilling into large positive integers.
func Example_numerics() {

	// For more natural Float -> Integer when the underlying value is a string.
	// Conversion functions will always try to parse the value as the target type
	// first. If parsing fails float parsing with truncation will be attempted.
	fmt.Println(conv.Int(`-123.456`)) // -123

	// This does not apply for unsigned integers if the value is negative. Instead
	// performing a more intuitive (to the human) truncation to zero.
	fmt.Println(conv.Uint(`-123.456`)) // 0

	// Numeric conversions from Float -> time.Duration will separate the integer
	// and fractional portions into a more natural conversion.
	fmt.Println(conv.Duration(`12.34`)) // 12.34s

	// All other numeric conversions to durations assign the elapsed nanoseconds using Go
	// conversions.
	fmt.Println(conv.Duration(`123456`)) // 34h17m36s

	// Output:
	// -123
	// 0
	// 12.34s
	// 34h17m36s
}

func Example_slices() {

	// Slice does not need initialized.
	var into []int64

	// You must pass a pointer to a slice.
	err := conv.Slice(&into, []string{"123", "456", "6789"})
	if err != nil {
		log.Fatal("err:", err)
	}

	for _, v := range into {
		fmt.Println("v:", v)
	}

	// Output:
	// v: 123
	// v: 456
	// v: 6789
}

func Example_maps() {

	// Map must be initialized
	into := make(map[string]int64)

	// No need to pass a pointer
	err := conv.Map(into, []string{"123", "456", "6789"})
	if err != nil {
		log.Fatal("err:", err)
	}

	// This is just for testing determinism since keys are randomized.
	var keys []string
	for k := range into {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Print the keys
	for _, k := range keys {
		fmt.Println("k:", k, "v:", into[k])
	}

	// Output:
	// k: 0 v: 123
	// k: 1 v: 456
	// k: 2 v: 6789
}

// In short, panics should not occur within this library under any circumstance.
// This obviously excludes any oddities that may surface when the runtime is not
// in a healthy state, i.e. uderlying system instability, memory exhaustion. If
// you are able to create a reproducible panic please file a bug report.
func Example_panics() {
	fmt.Println(conv.Bool(nil))
	fmt.Println(conv.Bool([][]int{}))
	fmt.Println(conv.Bool((chan string)(nil)))
	fmt.Println(conv.Bool((*interface{})(nil)))
	fmt.Println(conv.Bool((*interface{})(nil)))
	fmt.Println(conv.Bool((**interface{})(nil)))

	// Output:
	// false
	// false
	// false
	// false
	// false
	// false
}
