# Go Package: conv

  [About](#about) | [Usage](#usage) | [Go Doc](https://godoc.org/github.com/cstockton/go-conv)

  > Get:
  > ```bash
  > go get -u github.com/cstockton/go-conv
  > ```
  >
  > Example:
  > ```Go
  > conv.Bool("true") // bool(true)
  > conv.Int64("-1.23") // -1
  > conv.Duration("123ns") // time.Duration(123)
  > ```


## About

Small library to make working with various types of values a bit easier. When
`go generate` was released I thought it would be useful to integrate into unit
tests and code generation for some of the more boiler plate type libraries
and operations I perform in Go, like conversions. All heavy unit testing types
as well as assertion definitions are within `*_test.go` files in the `init()`
function. This means they are not exported and you do not pay an allocation tax
at program startup for these features.


## Usage


Package conv provides conversions without using reflection across 	most built-in Go types through type assertion switches. Conversions are 	defined in Assertion structures which are used to generate unit tests as 	well as documentation. This page was generated using "go generate" and each 	code example here is ran within a unit test before it is generated to ensure 	the accuracy of this page. In addition unit tests will fail when this document 	becomes out of date.

### String Conversions

String conversion from other string values will be returned 	without modification. As a special case []byte will also be returned 	after a Go string conversion is applied.

  > Example:
  > ```Go
  > String([]byte{0x46, 0x6f, 0x6f}) // "Foo"
  > String("Foo") // "Foo"
  > String("") // ""
  > ```

String conversion from any other value will simply be the  result of calling fmt.Sprintf("%v", value).

  > Example:
  > ```Go
  > String(true) // "true"
  > String(false) // "false"
  > String(123) // "123"
  > String(0xc) // "12"
  > String(time.Duration(3723000000000)) // "1h2m3s"
  > String(time.Time{sec:63393490862, nsec:3, loc: time.UTC}) // 2009-11-10 23:01:02.000000003 +0000 UTC
  > ```

String conversion from types that do not have a valid conversion path 	will still have sane string conversion for troubleshooting.

  > Example:
  > ```Go
  > String(struct { msg string }{msg:"Foo"}) // "{Foo}"
  > ```


### Bool Conversions

Bool conversion from other bool values will be returned 	without modification.

  > Example:
  > ```Go
  > Bool(true) // true
  > Bool(false) // false
  > ```

Bool conversion from strings accepts "1", "t", "T", "true", "True", 	"TRUE", "y", "Y", "yes", "Yes", "YES" for true. It returns false for "0", 	"f", "F", "false", "False", "FALSE", "n", "N", "no", "No", "NO".

  > Example:
  > ```Go
  > Bool("true") // true
  > Bool("yes") // true
  > Bool("T") // true
  > Bool("0") // false
  > Bool("Foo") // false
  > ```

Bool conversion from all other types will return true unless 	it is the zero value for the given type.

  > Example:
  > ```Go
  > Bool(0) // false
  > Bool(123) // true
  > Bool(time.Duration(0)) // false
  > Bool(time.Duration(123)) // true
  > Bool(time.Time{sec:0, nsec:0, loc: time.UTC}) // false
  > Bool(time.Time{sec:63613808345, nsec:123450000, loc: time.UTC}) // true
  > ```


### Duration Conversions

Duration conversion from other duration values will be returned 	without modification.

  > Example:
  > ```Go
  > Duration(time.Duration(3000000000)) // time.Duration(3000000000)
  > Duration(time.Duration(-3000000000)) // time.Duration(-3000000000)
  > ```

Duration conversion from strings attempts to use time.ParseDuration() 	as documented in the standard Go libraries time package. If parsing 	fails then it will be passed along to Int64() followed by a standard 	Go time.Duration conversion.

  > Example:
  > ```Go
  > Duration("3s") // time.Duration(3000000000)
  > Duration("-3s") // time.Duration(-3000000000)
  > Duration("3") // time.Duration(3)
  > ```

Duration conversion from time.Time is a special case that 	returns the time elapsed since time.Unix(). This behavior should be 	considered experimental.

  > Example:
  > ```Go
  > Duration(time.Time{sec:63613808345, nsec:123450000, loc: time.UTC}) // time.Duration(1478211545)
  > ```

Duration conversion from all other values will be the result 	of calling Int64() followed by a standard Go time.Duration conversion. 	The result of the Int64() conversion will be in nanoseconds.

  > Example:
  > ```Go
  > Duration("3") // time.Duration(3)
  > Duration(3) // time.Duration(3)
  > Duration(true) // time.Duration(1)
  > Duration(false) // time.Duration(0)
  > ```


### Time Conversions

Time conversion from other time values will be returned 	without modification.

  > Example:
  > ```Go
  > Time(time.Time{sec:63393490862, nsec:3, loc: time.UTC}) // time.Time{sec:63393490862, nsec:3, loc: time.UTC}
  > Time(time.Time{sec:0, nsec:0, loc: time.UTC}) // time.Time{sec:0, nsec:0, loc: time.UTC}
  > ```

Time conversion from time.Duration is a special case that 	returns the current moment from time.Now() plus the duration. For example 	if the time was 2009-11-10 23:01:02, the string below would be a time.Time 	struct with an additional 3 seconds.

  > Example:
  > ```Go
  > Time(time.Duration(time.Seconds*3)) // time.Time{"2009-11-10 23:01:05.000000003 +0000 UTC"}
  > ```

Time conversion from strings will be passed through time.Parse 	using a variety of formats. Strings that could not be parsed along 	with all other values will return an empty time.Time{} struct.

  > Example:
  > ```Go
  > Time("Monday, 02 Jan 2006 15:04:05 -0700") // time.Time{sec:63271836245, nsec:0, loc: time.UTC}
  > Time("1") // time.Time{sec:0, nsec:0, loc: time.UTC}
  > Time(true) // time.Time{sec:0, nsec:0, loc: time.UTC}
  > Time(1) // time.Time{sec:0, nsec:0, loc: time.UTC}
  > ```


### Numeric Conversions

Numeric conversion from other numeric values of an identical 	type will be returned without modification. Conversions across different 	types will follow the rules in the official Go language spec under the 	heading "Conversions between numeric types"

  > Example:
  > ```Go
  > Int64(10) // 10
  > Int8(10) // 10
  > Int8((127+0i)) // 127
  > ```

This means overflow is identical to conversion within Go code 	through type conversion during runtime which do not panic.

  > Example:
  > ```Go
  > Int8(128) // -128
  > Uint8(-128) // 0x80
  > Uint8(256) // 0x0
  > Int8(12345.6789) // 57
  > Float32(12345.6789) // 12345.679
  > Uint64(12345.6789) // 0x3039
  > ```

Numeric conversion from strings uses the associated strconv.Parse* 	from the standard library. Overflow is handled like the cases above.

  > Example:
  > ```Go
  > Int64("-123456789") // -123456789
  > Uint64("123456789") // 0x75bcd15
  > Float32("12345.6789") // 12345.679
  > Int64("true") // 1
  > Int64("false") // 0
  > Int64("abcde") // 0
  > ```

For more natural Float -> Integer when the underlying value is a 	string. Conversion functions will always try to parse the value as the 	target type first. If parsing fails float parsing with truncation will be 	attempted. This deviates from the standard library but should be useful in 	common practice.

  > Example:
  > ```Go
  > strconv.Atoi("-123.456") // 0 - err: invalid syntax
  > Int64("-123.456") // -123
  > Int64("123.456") // 123
  > Uint64("123.456") // 0x7b
  > ```

This does not apply for unsigned integers if the value is negative. 	Instead performing a more intuitive (to the human) truncation to zero.

  > Example:
  > ```Go
  > strconv.Atoi("-123.456") // 0 - err: invalid syntax
  > Int64("-123.456") // -123
  > Int64("-1.23") // -1
  > Uint64("-1.23") // 0x0
  > ```

Numeric conversions from durations assign the elapsed 	nanoseconds using Go conversions.

  > Example:
  > ```Go
  > Int8(time.Duration(20)) // 20
  > Int64(time.Duration(-20)) // -20
  > ```

Numeric conversions to times are a special case that result 	in the time since the unix epoch as returned by Time.Unix(). This behavior 	is experimental and may change in the future.

  > Example:
  > ```Go
  > Int(time.Time{sec:63393490862, nsec:3, loc: time.UTC}) // 1257894062
  > Int(time.Time{sec:0, nsec:0, loc: time.UTC}) // -62135596800
  > ```

Numeric conversions from bool are 1 for true, 0 for false. 	All other conversions that fail return 0.

  > Example:
  > ```Go
  > Int64(true) // 1
  > Int64(false) // 0
  > Int64(struct { msg string }{msg:"Hello World"}) // 0
  > ```


### Pointer Conversions

All conversions will allow up to two levels of pointer 	indirection if non-nil and the value pointer to is a convertable type. 	This is facilitated via the Indirect() function.

  > Example:
  > ```Go
  > (*string)(0x001) // Underlying value, i.e.: `Foo`
  > (**string)(0x01) // Underlying value, i.e.: `Foo`
  > ```


### Panics

This library should not panic under any input for conversions. 	If you are able to produce a panic please file a bug report.

  > Example:
  > ```Go
  > Int64(nil) // 0
  > Int64([][]int{}) // 0
  > Int64((chan string)(nil)) // 0
  > Int64((*interface{})(nil)) // 0
  > Int64((*interface {})(0x01)) // 0
  > Int64((**interface {})(0x1)) // 0
  > ```


### Value

Value is a convenience struct for performing Conversion. It has a 	single field V of interface{} type which is passed to the associated 	conversion functions.

  > Example:
  > ```Go
  > func (v Value) Bool() bool { return Bool(v.V) }
  > ```

This means you may wrap any value with Value{...} for conversions.

  > Example:
  > ```Go
  > v := Value{"12345.6789"}
  > v.Int64() // 12345
  > v.Float64() // 12345.6789
  > v.Float32() // 12345.679
  > ```




## Contributing

Feel free to make a issues defining some conversion's that don't exist or
challenging the current ones. If you find a conversion you wish to change the
behavior of you MUST add an Assertion within a `*_test` file. A good place may
be [conv_test.go](conv_test.go).

  - Example assertions for a Complex128 type, the functions are helpers that are defined in [assert_test.go](assert_test.go).

    > Go:
    > ```Go
    > asserts("Complex128",
    >   assert(complex128(-1), true, "(-1+0i)", numerics(-1), durations(-1)),
    >   assert(complex128(0), false, "(0+0i)", numerics(0), durations(0)),
    >   assert(complex128(1), true, "(1+0i)", numerics(1), durations(1)),
    > )

This may cause other documentation to need updating with the `go generate` command.
Keep in mind the documentation through the `Section` and `Group` types is meant
to be consumable rather than comprehensive. Most conversion Assertions should live
outside of those. The template files are in the [testdata](testdata)
folder.

The source and target files are defined below:

  - [README.md](README.md) -> [README.md.tpl](testdata/README.md.tpl)
  - [doc.go](doc.go) -> [README.md.tpl](testdata/doc.go.tpl)
  - [numeric.go](numeric.go) -> [README.md.tpl](testdata/numeric.go.tpl)


## Bugs and Patches

  Feel free to report bugs and submit pull requests.

  * bugs:
    <https://github.com/cstockton/go-conv/issues>
  * patches:
    <https://github.com/cstockton/go-conv/pulls>



[Go Doc]: https://godoc.org/github.com/cstockton/go-value
