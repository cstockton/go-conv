# Go Package: conv

  [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/cstockton/go-conv)
  [![Go Report Card](https://goreportcard.com/badge/github.com/cstockton/go-conv?style=flat-square)](https://goreportcard.com/report/github.com/cstockton/go-conv)
  [![Coverage Status](https://img.shields.io/codecov/c/github/cstockton/go-conv/master.svg?style=flat-square)](https://codecov.io/github/cstockton/go-conv?branch=master)
  [![Build Status](http://img.shields.io/travis/cstockton/go-conv.svg?style=flat-square)](https://travis-ci.org/cstockton/go-conv)
  [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/cstockton/go-conv/master/LICENSE)

  > Get:
  > ```bash
  > go get -u github.com/cstockton/go-conv
  > ```
  >
  > Example:
  > ```Go
  > // Basic types
  > if got, err := conv.Bool(`TRUE`); err == nil {
  > 	fmt.Printf("conv.Bool(`TRUE`)\n  -> %v\n", got)
  > }
  > if got, err := conv.Duration(`1m2s`); err == nil {
  > 	fmt.Printf("conv.Duration(`1m2s`)\n  -> %v\n", got)
  > }
  > var date time.Time
  > err := conv.Infer(&date, `Sat Mar 7 11:06:39 PST 2015`)
  > fmt.Printf("conv.Infer(&date, `Sat Mar 7 11:06:39 PST 2015`)\n  -> %v\n", got)
  > ```
  >
  > Output:
  > ```Go
  > conv.Bool(`TRUE`)
  >   -> true
  > conv.Duration(`1m2s`)
  >   -> 1m2s
  > conv.Infer(&date, `Sat Mar 7 11:06:39 PST 2015`)
  >   -> 2015-03-07 11:06:39 +0000 PST
  > ```


## Intro

**Notice:** If you begin getting compilation errors use the v1 import path `gopkg.in/cstockton/go-conv.v1` for an immediate fix and to future-proof.

Package conv provides fast and intuitive conversions across Go types. This library uses reflection to be robust but will bypass it for common conversions, for example string conversion to any type will never use reflection. All functions are safe for concurrent use by multiple Goroutines.

### Overview

  All conversion functions accept any type of value for conversion, if unable
  to find a reasonable conversion path they will return the target types zero
  value and an error.

  > Example:
  > ```Go
  > // The zero value and a non-nil error is returned on failure.
  > fmt.Println(conv.Int("Foo"))
  >
  > // Conversions are allowed as long as the underlying type is convertable, for
  > // example:
  > type MyString string
  > fmt.Println(conv.Int(MyString("42"))) // 42, nil
  >
  > // Pointers will be dereferenced when appropriate.
  > str := "42"
  > fmt.Println(conv.Int(&str)) // 42, nil
  >
  > // You may infer values from the base type of a pointer, giving you one
  > // function signature for all conversions. This may be convenient when the
  > // types are not known until runtime and reflection must be used.
  > var val int
  > err := conv.Infer(&val, `42`)
  > fmt.Println(val, err) // 42, nil
  > ```
  >
  > Output:
  > ```Go
  > 0 cannot convert "Foo" (type string) to int
  > 42 <nil>
  > 42 <nil>
  > 42 <nil>
  > ```


### Infer

  Infer will perform conversion by inferring the conversion operation from
  a pointer to a supported T of the `into` param. Since the value is assigned
  directly only a error value is returned, meaning no type assertions needed.

  > Example:
  > ```Go
  > // Infer requires a pointer to all types.
  > var into int
  > if err := conv.Infer(into, `42`); err != nil {
  > 	fmt.Println(err)
  > }
  > if err := conv.Infer(&into, `42`); err == nil {
  > 	fmt.Println(into)
  > }
  >
  > // Same as above but using new()
  > truth := new(bool)
  > if err := conv.Infer(truth, `TRUE`); err != nil {
  > 	fmt.Println("Failed!")
  > }
  > ```
  >
  > Output:
  > ```Go
  > cannot infer conversion for non-pointer 0 (type int)
  > 42
  > ```


### Strings

  String conversion from any values outside the cases below will simply be the
  result of calling fmt.Sprintf("%v", value), meaning it can not fail. An error
  is still provided and you should check it to be future proof.

  > Example:
  > ```Go
  > // String conversion from other string values will be returned without
  > // modification.
  > fmt.Println(conv.String("Foo"))
  >
  > // As a special case []byte will also be returned after a Go string conversion
  > // is applied.
  > fmt.Println(conv.String([]byte("Foo")))
  >
  > // String conversion from types that do not have a valid conversion path will
  > // still have sane string conversion for troubleshooting.
  > fmt.Println(conv.String(struct{ msg string }{"Foo"}))
  > ```
  >
  > Output:
  > ```Go
  > Foo <nil>
  > Foo <nil>
  > {Foo} <nil>
  > ```


### Bools

  Bool conversion supports all the paths provided by the standard libraries
  strconv.ParseBool when converting from a string, all other conversions are
  simply true when not the types zero value. As a special case zero length map
  and slice types are also false, even if initialized.

  > Example:
  > ```Go
  > // Bool conversion from other bool values will be returned without
  > // modification.
  > fmt.Println(conv.Bool(true))
  > fmt.Println(conv.Bool(false))
  >
  > // Bool conversion from strings consider the following values true:
  > //   "t", "T", "true", "True", "TRUE",
  > // 	 "y", "Y", "yes", "Yes", "YES", "1"
  > //
  > // It considers the following values false:
  > //   "f", "F", "false", "False", "FALSE",
  > //   "n", "N", "no", "No", "NO", "0"
  > fmt.Println(conv.Bool("T"))
  > fmt.Println(conv.Bool("False"))
  >
  > // Bool conversion from other supported types will return true unless it is
  > // the zero value for the given type.
  > fmt.Println(conv.Bool(int64(123)))
  > fmt.Println(conv.Bool(int64(0)))
  > fmt.Println(conv.Bool(time.Duration(123)))
  > fmt.Println(conv.Bool(time.Duration(0)))
  > fmt.Println(conv.Bool(time.Now()))
  > fmt.Println(conv.Bool(time.Time{}))
  >
  > // All other types will return false.
  > fmt.Println(conv.Bool(struct{ string }{""}))
  > ```
  >
  > Output:
  > ```Go
  > true <nil>
  > false <nil>
  > true <nil>
  > false <nil>
  > true <nil>
  > false <nil>
  > true <nil>
  > false <nil>
  > true <nil>
  > false <nil>
  > false cannot convert struct { string }{string:""} (type struct { string }) to bool
  > ```


### Numerics

  Numeric conversion from other numeric values of an identical type will be
  returned without modification. Numeric conversions deviate slightly from Go
  when dealing with under/over flow. When performing a conversion operation
  that would overflow, we instead assign the maximum value for the target type.
  Similarly, conversions that would underflow are assigned the minimun value
  for that type, meaning unsigned integers are given zero values instead of
  spilling into large positive integers.

  > Example:
  > ```Go
  > // For more natural Float -> Integer when the underlying value is a string.
  > // Conversion functions will always try to parse the value as the target type
  > // first. If parsing fails float parsing with truncation will be attempted.
  > fmt.Println(conv.Int("-123.456")) // -123
  >
  > // This does not apply for unsigned integers if the value is negative. Instead
  > // performing a more intuitive (to the human) truncation to zero.
  > fmt.Println(conv.Uint("-123.456")) // 0
  > ```
  >
  > Output:
  > ```Go
  > -123 <nil>
  > 0 <nil>
  > ```


### Durations

  Duration conversion supports all the paths provided by the standard libraries
  time.ParseDuration when converting from strings, with a couple enhancements
  outlined below.

  > Example:
  > ```Go
  > // Duration conversion from strings will first attempt to parse as a Go
  > // duration value using ParseDuration, then fall back to numeric conventions.
  > fmt.Println(conv.Duration("1h1m100ms"))     // 1h1m0.1s
  > fmt.Println(conv.Duration("3660100000000")) // 1h1m0.1s
  >
  > // Numeric conversions directly convert to time.Duration nanoseconds.
  > fmt.Println(conv.Duration(3660100000000)) // 1h1m0.1s
  >
  > // Floats deviate from the numeric conversion rules, instead
  > // separating the integer and fractional portions into seconds.
  > fmt.Println(conv.Duration("3660.10"))        // 1h1m0.1s
  > fmt.Println(conv.Duration(float64(3660.10))) // 1h1m0.1s
  >
  > // Complex numbers are Float conversions using the real number.
  > fmt.Println(conv.Duration(complex(3660.10, 0))) // 1h1m0.1s
  >
  > // Duration conversion from time.Duration and any numerical type will be
  > // converted using a standard Go conversion. This includes strings
  > fmt.Println(conv.Duration(time.Nanosecond)) // 1s
  > fmt.Println(conv.Duration(byte(1)))         // 1ns
  > ```
  >
  > Output:
  > ```Go
  > 1h1m0.1s <nil>
  > 1h1m0.1s <nil>
  > 1h1m0.1s <nil>
  > 1h1m0.1s <nil>
  > 1h1m0.1s <nil>
  > 1h1m0.1s <nil>
  > 1ns <nil>
  > 1ns <nil>
  > ```


### Panics

  In short, panics should not occur within this library under any circumstance.
  This obviously excludes any oddities that may surface when the runtime is not
  in a healthy state, i.e. uderlying system instability, memory exhaustion. If
  you are able to create a reproducible panic please file a bug report.

  > Example:
  > ```Go
  > // The zero value for the target type is always returned.
  > fmt.Println(conv.Bool(nil))
  > fmt.Println(conv.Bool([][]int{}))
  > fmt.Println(conv.Bool((chan string)(nil)))
  > fmt.Println(conv.Bool((*interface{})(nil)))
  > fmt.Println(conv.Bool((*interface{})(nil)))
  > fmt.Println(conv.Bool((**interface{})(nil)))
  > ```
  >
  > Output:
  > ```Go
  > false cannot convert <nil> (type <nil>) to bool
  > false <nil>
  > false <nil>
  > false cannot convert (*interface {})(nil) (type *interface {}) to bool
  > false cannot convert (*interface {})(nil) (type *interface {}) to bool
  > false cannot convert (**interface {})(nil) (type **interface {}) to bool
  > ```


## Contributing

Feel free to create issues for bugs, please ensure code coverage remains 100%
with any pull requests.


## Bugs and Patches

  Feel free to report bugs and submit pull requests.

  * bugs:
    <https://github.com/cstockton/go-conv/issues>
  * patches:
    <https://github.com/cstockton/go-conv/pulls>
