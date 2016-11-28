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
  > fmt.Printf("Basics:\n  `TRUE` -> %#v\n  `12.3` -> %#v\n  `1m2s` -> %#v\n\n",
  > 	conv.Bool("YES"), conv.Int64("12.3"), conv.Duration("1m2s"))
  > 
  > // Slice and map support
  > from := []string{"1.2", "34.5", "-678.9"}
  > var into []float64
  > conv.Slice(&into, from) // type inferred from the element of `into`
  > fmt.Printf("Slice:\n  %#v\n    -> %#v\n\n", from, into
  > ```
  >
  > Output:
  > ```Go
  > Basics:
  >   `TRUE` -> true
  >   `12.3` -> 12
  >   `1m2s` -> 62000000000
  > 
  > Slice:
  >   []string{"1.2", "34.5", "-678.9"}
  >     -> []float64{1.2, 34.5, -678.9}
  > ```


## Intro

Package conv provides fast and intuitive conversions across Go types. This library uses reflection to be robust but will bypass it for common conversions, for example string conversion to any type will never use reflection. In most cases this library is as fast or faster then the standard library for similar operations due to various aggressive (but safe) optimizations. The only external dependency ([iter](https://github.com/cstockton/go-iter)) has 100% test coverage and is maintained by me. It is used to walk the values given for map and slice conversion and will **never panic**. All methods and functions are **safe for concurrent use by multiple Goroutines**, with a single exception that Slice and Map conversion under certain circumstances may produce undefined results if they are mutated while being traversed.


### Overview

  All methods and functions accept any type of value for conversion, if unable
  to find a reasonable conversion path they will return the target types zero
  value. The Conv struct will also report an error on failure, while all the
  top level functions (conv.Bool(...), conv.Time(...), etc) will only return a
  single value for cases that you wish to leverage zero values. These functions
  are powered by the "DefaultConverter" variable so you may replace it with
  your own Converter or a Conv struct to adjust behavior.

  > Example:
  > ```Go
  > // All top level conversion functions discard errors, returning the types zero
  > // value instead.
  > fmt.Println(conv.Time("bad time string")) // time.Time{}
  > fmt.Println(conv.Time("Sat Mar 7 11:06:39 PST 2015"))
  > 
  > // Conversions are allowed as long as the underlying type is convertable, for
  > // example:
  > type MyString string
  > fmt.Println(conv.Int(MyString(`123`))) // 123
  > 
  > // Pointers will be dereferenced when appropriate.
  > s := `123`
  > fmt.Println(conv.Int(&s)) // 123
  > 
  > // If you would like to know if errors occur you may use a conv.Conv. It is
  > // safe for use by multiple Go routines and needs no initialization.
  > var c conv.Conv
  > i, err := c.Int(`Foo`)
  > 
  > // Got 0 because of err: cannot convert "Foo" (type string) to int
  > fmt.Printf("Got %v because of err: %v", i, err
  > ```
  >
  > Output:
  > ```Go
  > 0001-01-01 00:00:00 +0000 UTC
  > 2015-03-07 11:06:39 +0000 PST
  > 123
  > 123
  > Got 0 because of err: cannot convert "Foo" (type string) to int
  > ```


### Strings

  String conversion from any values outside the cases below will simply be the
  result of calling fmt.Sprintf("%v", value), meaning it can not fail. An error
  is still provided and you should check it to be future proof.

  > Example:
  > ```Go
  > // String conversion from other string values will be returned without
  > // modification.
  > fmt.Println(conv.String(`Foo`))
  > 
  > // As a special case []byte will also be returned after a Go string conversion
  > // is applied.
  > fmt.Println(conv.String([]byte(`Foo`)))
  > 
  > // String conversion from types that do not have a valid conversion path will
  > // still have sane string conversion for troubleshooting.
  > fmt.Println(conv.String(struct{ msg string }{"Foo"})
  > ```
  >
  > Output:
  > ```Go
  > Foo
  > Foo
  > {Foo}
  > ```


### Bools

  Bool Conversions supports all the paths provided by the standard libraries
  strconv.ParseBool when converting from a string, all other conversions are
  simply true when not the types zero value. As a special case zero length map
  and slice types are also false, even if initialized.

  > Example:
  > ```Go
  > // Bool conversion from other bool values will be returned without
  > // modification.
  > fmt.Println(conv.Bool(true), conv.Bool(false))
  > 
  > // Bool conversion from strings consider the following values true:
  > //   "t", "T", "true", "True", "TRUE",
  > // 	 "y", "Y", "yes", "Yes", "YES", "1"
  > //
  > // It considers the following values false:
  > //   "f", "F", "false", "False", "FALSE",
  > //   "n", "N", "no", "No", "NO", "0"
  > fmt.Println(conv.Bool("T"), conv.Bool("False"))
  > 
  > // Bool conversion from other supported types will return true unless it is
  > // the zero value for the given type.
  > fmt.Println(conv.Bool(int64(123)), conv.Bool(int64(0)))
  > fmt.Println(conv.Bool(time.Duration(123)), conv.Bool(time.Duration(0)))
  > fmt.Println(conv.Bool(time.Now()), conv.Bool(time.Time{}))
  > 
  > // All other types will return false.
  > fmt.Println(conv.Bool(struct{ string }{""})
  > ```
  >
  > Output:
  > ```Go
  > true false
  > true false
  > true false
  > true false
  > true false
  > false
  > ```


### Numerics

  Numeric conversion from other numeric values of an identical type will be
  returned without modification. Numeric conversions deviate slightly from Go
  when dealing with under/over flow. When performing a conversion operation
  that would overflow, we instead assign the maximum value for the target type.
  Similarly, conversions that would underflow are assigned the minimun value
  for that type, meaning unsigned integers are given zero values isntead of
  spilling into large positive integers.

  > Example:
  > ```Go
  > // For more natural Float -> Integer when the underlying value is a string.
  > // Conversion functions will always try to parse the value as the target type
  > // first. If parsing fails float parsing with truncation will be attempted.
  > fmt.Println(conv.Int(`-123.456`)) // -123
  > 
  > // This does not apply for unsigned integers if the value is negative. Instead
  > // performing a more intuitive (to the human) truncation to zero.
  > fmt.Println(conv.Uint(`-123.456`)) // 
  > ```
  >
  > Output:
  > ```Go
  > -123
  > 0
  > ```


### Durations

  Duration conversion supports all the paths provided by the standard libraries
  time.ParseDuration when converting from strings, with a couple enhancements
  outlined below.

  > Example:
  > ```Go
  > // Duration conversion from other time.Duration values will be returned
  > // without modification.
  > fmt.Println(conv.Duration(time.Duration(time.Second))) // 1s
  > fmt.Println(conv.Duration("1h30m"))                    // 1h30m0s
  > 
  > // Duration conversions from floats will separate the integer
  > // and fractional portions into a more natural conversion.
  > fmt.Println(conv.Duration("12.15"))
  > 
  > // All other duration conversions from numeric types assign the
  > // elapsed nanoseconds using Go conversions.
  > fmt.Println(conv.Duration(`123456`)) // 34h17m36
  > ```
  >
  > Output:
  > ```Go
  > 1s
  > 1h30m0s
  > 12.15s
  > 34h17m36s
  > ```


### Slices

  Slice conversion will infer the element type from the given slice, using the
  associated conversion function as the given structure is traversed
  recursively. The behavior if the value is mutated during iteration is
  undefined, though at worst an error will be returned as this library will
  never panic.
  
  An error is returned if the below restrictions are not met:
  
    - It must be a pointer to a slice, it does not have to be initialized
    - The element must be a T or *T of a type supported by this library

  > Example:
  > ```Go
  > // Slice does not need initialized.
  > var into []int64
  > 
  > // You must pass a pointer to a slice.
  > err := conv.Slice(&into, []string{"123", "456", "6789"})
  > if err != nil {
  > 	log.Fatal("err:", err)
  > }
  > 
  > for _, v := range into {
  > 	fmt.Println("v:", v)
  > 
  > ```
  >
  > Output:
  > ```Go
  > v: 123
  > v: 456
  > v: 6789
  > ```


### Maps

  Map conversion will infer the conversion functions to use from the key and
  element types of the given map. The second argument will be walked as
  described in the supporting package, go-iter.
  
  An error is returned if the below restrictions are not met:
  
    - It must be a non-pointer, non-nil initialized map
    - Both the key and element T must be supported by this library
    - The key must be a value T, the element may be a T or *T
  
  Excerpt from github.com/cstockton/go-iter iter.Walk:
  
  Walk will recursively walk the given interface value as long as an error does
  not occur. The pair func will be given a interface value for each value
  visited during walking and is expected to return an error if it thinks the
  traversal should end. A nil value and error is given to the walk func if an
  inaccessible value (can't reflect.Interface()) is found.
  
  Walk is called on each element of maps, slices and arrays. If the underlying
  iterator is configured for channels it receives until one fails. Channels
  should probably be avoided as ranging over them is more concise.

  > Example:
  > ```Go
  > // Map must be initialized
  > into := make(map[string]int64)
  > 
  > // No need to pass a pointer
  > err := conv.Map(into, []string{"123", "456", "6789"})
  > if err != nil {
  > 	log.Fatal("err:", err)
  > }
  > 
  > // This is just for testing determinism since keys are randomized.
  > var keys []string
  > for k := range into {
  > 	keys = append(keys, k)
  > }
  > sort.Strings(keys)
  > 
  > // Print the keys
  > for _, k := range keys {
  > 	fmt.Println("k:", k, "v:", into[k])
  > 
  > ```
  >
  > Output:
  > ```Go
  > k: 0 v: 123
  > k: 1 v: 456
  > k: 2 v: 6789
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
  > fmt.Println(conv.Bool((**interface{})(nil))
  > ```
  >
  > Output:
  > ```Go
  > false
  > false
  > false
  > false
  > false
  > false
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
