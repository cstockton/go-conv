package conv

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	generate = flag.Bool("generate", false, "generate files from templates.")
	nowrite  = flag.Bool("nowrite", false, "print to stdout instead of writing to files.")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestTypes(t *testing.T) {
	for _, assertion := range assertions {
		assertion := assertion
		key := fmt.Sprintf("%v/%v",
			strings.Join(assertion.Tags, "/"), assertion.Interface)

		t.Run(key, func(t *testing.T) {
			t.Run("T", func(t *testing.T) {
				if errs := assertion.Assert(Value{assertion.Interface}); len(errs) != 0 {
					for _, err := range errs {
						t.Errorf("%s", err)
					}
				}
			})
			t.Run("*T", func(t *testing.T) {
				p := &assertion.Interface
				if errs := assertion.Assert(Value{p}); len(errs) != 0 {
					for _, err := range errs {
						t.Errorf("%s", err)
					}
				}
			})
			t.Run("**T", func(t *testing.T) {
				p := &assertion.Interface
				if errs := assertion.Assert(Value{&p}); len(errs) != 0 {
					for _, err := range errs {
						t.Errorf("%s", err)
					}
				}
			})
		})
	}
}

type MyInt8 int8

func (m MyInt8) Uint8() uint8 { return 42 }

type MyValue struct {
	Value
}

func V(value interface{}) MyValue {
	return MyValue{}
}
func (m MyValue) Uint16() uint16 { return 42 }

func init() {
	timeParse := func(f, s string) time.Time {
		t, err := time.Parse(f, s)
		if err != nil {
			panic(err)
		}
		return t
	}
	DA := func(s string) DocAssertion { return DocAssertion(s) }
	myInt8 := MyInt8(-1 << 7)
	assert(myInt8, uint8(42))

	myValue := Value{"12345.6789"}
	assert(myValue.Int64(), int64(myValue.Int64()))
	assert(myValue.Float32(), float32(myValue.Float32()))

	summarize(`Package conv provides conversions without using reflection across
		most built-in Go types through type assertion switches. Conversions are
		defined in Assertion structures which are used to generate unit tests as
		well as documentation. This page was generated using "go generate" and each
		code example here is ran within a unit test before it is generated to ensure
		the accuracy of this page. In addition unit tests will fail when this document
		becomes out of date.`)
	section("String Conversions",
		group(`String conversion from other string values will be returned
			without modification. As a special case []byte will also be returned
			after a Go string conversion is applied.`,
			assert([]byte("Foo"), "Foo"),
			assert("Foo", "Foo"),
			assert("", ""),
		),
		group(`String conversion from any other value will simply be the
			 result of calling fmt.Sprintf("%v", value).`,
			assert(true, "true"),
			assert(false, "false"),
			assert(int64(123), "123"),
			assert(uint8(12), "12"),
			assert(time.Duration(3723000000000), "1h2m3s"),
			assert(DocAssertion("String(time.Time{sec:63393490862, nsec:3, loc: time.UTC})"),
				DocAssertion("2009-11-10 23:01:02.000000003 +0000 UTC")),
		),
		group(`String conversion from types that do not have a valid conversion path
			will still have sane string conversion for troubleshooting.`,
			assert(struct{ msg string }{"Foo"}, "{Foo}"),
		),
	)
	section("Bool Conversions",
		group(`Bool conversion from other bool values will be returned
			without modification.`,
			assert(true, true),
			assert(false, false),
		),
		group(`Bool conversion from strings accepts "1", "t", "T", "true", "True",
			"TRUE", "y", "Y", "yes", "Yes", "YES" for true. It returns false for "0",
			"f", "F", "false", "False", "FALSE", "n", "N", "no", "No", "NO".`,
			assert("true", true),
			assert("yes", true),
			assert("T", true),
			assert("0", false),
			assert("Foo", false),
		),
		group(`Bool conversion from all other types will return true unless
			it is the zero value for the given type.`,
			assert(int64(0), false),
			assert(int64(123), true),
			assert(time.Duration(0), false),
			assert(time.Duration(123), true),
			assert(time.Time{}, false),
			assert(timeParse(time.RFC3339Nano, "2016-11-03T15:04:05.12345-07:15"), true),
		),
	)
	section("Duration Conversions",
		group(`Duration conversion from other duration values will be returned
			without modification.`,
			assert(time.Duration(time.Second*3), time.Duration(time.Second*3)),
			assert(time.Duration(-(time.Second*3)), time.Duration(-(time.Second*3))),
		),
		group(`Duration conversion from strings attempts to use time.ParseDuration()
			as documented in the standard Go libraries time package. If parsing
			fails then it will be passed along to Int64() followed by a standard
			Go time.Duration conversion.`,
			assert("3s", time.Duration(time.Second*3)),
			assert("-3s", time.Duration(-(time.Second*3))),
			assert("3", time.Duration(3)),
		),
		group(`Duration conversion from time.Time is a special case that
			returns the time elapsed since time.Unix(). This behavior should be
			considered experimental.`,
			assert(timeParse(time.RFC3339Nano, "2016-11-03T15:04:05.12345-07:15"),
				time.Duration(1478211545)),
		),
		group(`Duration conversion from all other values will be the result
			of calling Int64() followed by a standard Go time.Duration conversion.
			The result of the Int64() conversion will be in nanoseconds.`,
			assert("3", time.Duration(3)),
			assert(3, time.Duration(3)),
			assert(true, time.Duration(1)),
			assert(false, time.Duration(0)),
		),
	)
	section("Time Conversions",
		group(`Time conversion from other time values will be returned
			without modification.`,
			assert(
				time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC),
				time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC)),
			assert(time.Time{}, time.Time{}),
		),
		group(`Time conversion from time.Duration is a special case that
			returns the current moment from time.Now() plus the duration. For example
			if the time was 2009-11-10 23:01:02, the string below would be a time.Time
			struct with an additional 3 seconds.`,
			assert(
				DocAssertion(`Time(time.Duration(time.Seconds*3))`),
				DocAssertion(`time.Time{"2009-11-10 23:01:05.000000003 +0000 UTC"}`)),
		),
		group(`Time conversion from strings will be passed through time.Parse
			using a variety of formats. Strings that could not be parsed along
			with all other values will return an empty time.Time{} struct.`,
			assert("Monday, 02 Jan 2006 15:04:05 -0700", TimeAssertion{
				Moment: timeParse("Monday, 02 Jan 2006 15:04:05 -0700", "Monday, 02 Jan 2006 15:04:05 -0700"),
			}),
			assert("1", time.Time{}),
			assert(true, time.Time{}),
			assert(1, time.Time{}),
		),
	)
	section("Numeric Conversions",
		group(`Numeric conversion from other numeric values of an identical
			type will be returned without modification. Conversions across different
			types will follow the rules in the official Go language spec under the
			heading "Conversions between numeric types"`,
			assert(int8(10), int64(10)),
			assert(float64(10), int8(10)),
			assert(complex128(127), int8(127)),
		),
		group(`This means overflow is identical to conversion within Go code
			through type conversion during runtime which do not panic.`,
			assert(int64(1<<7), int8(-1<<7)),
			assert(int64(-1<<7), uint8(1<<7)),
			assert(int64(1<<8), uint8(0)),
			assert(float64(12345.6789), int8(57)),
			assert(float64(12345.6789), float32(12345.679)),
			assert(float64(12345.6789), Uint64(12345)),
		),
		group(`Numeric conversion from strings uses the associated strconv.Parse*
			from the standard library. Overflow is handled like the cases above.`,
			assert("-123456789", int64(-123456789)),
			assert("123456789", uint64(123456789)),
			assert("12345.6789", float32(12345.679)),
			assert("true", int64(1)),
			assert("false", int64(0)),
			assert("abcde", int64(0)),
		),
		group(`For more natural Float -> Integer when the underlying value is a
			string. Conversion functions will always try to parse the value as the
			target type first. If parsing fails float parsing with truncation will be
			attempted. This deviates from the standard library but should be useful in
			common practice.`,
			assert(DA(`strconv.Atoi("-123.456")`), DA(`0 - err: invalid syntax`)),
			assert("-123.456", int64(-123)),
			assert("123.456", int64(123)),
			assert("123.456", uint64(123)),
		),
		group(`This does not apply for unsigned integers if the value is negative.
			Instead performing a more intuitive (to the human) truncation to zero.`,
			assert(DA(`strconv.Atoi("-123.456")`), DA(`0 - err: invalid syntax`)),
			assert("-123.456", int64(-123)),
			assert("-1.23", int64(-1)),
			assert("-1.23", uint64(0)),
		),
		group(`Numeric conversions from durations assign the elapsed
			nanoseconds using Go conversions.`,
			assert(time.Duration(time.Nanosecond*20), int8(20)),
			assert(-time.Duration(time.Nanosecond*20), int64(-20)),
		),
		group(`Numeric conversions to times are a special case that result
			in the time since the unix epoch as returned by Time.Unix(). This behavior
			is experimental and may change in the future.`,
			assert(time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC), 1257894062),
			assert(time.Time{}, -62135596800),
		),
		group(`Numeric conversions from bool are 1 for true, 0 for false.
			All other conversions that fail return 0.`,
			assert(true, int64(1)),
			assert(false, int64(0)),
			assert(struct{ msg string }{"Hello World"}, int64(0)),
		),
	)
	section("Pointer Conversions",
		group(`All conversions will allow up to two levels of pointer
			indirection if non-nil and the value pointer to is a convertable type.
			This is facilitated via the Indirect() function.`,
			assert(DocAssertion("(*string)(0x001)"),
				DocAssertion("Underlying value, i.e.: `Foo`")),
			assert(DocAssertion("(**string)(0x01)"),
				DocAssertion("Underlying value, i.e.: `Foo`")),
		),
	)
	section("Panics",
		group(`This library should not panic under any input for conversions.
			If you are able to produce a panic please file a bug report.`,
			assert(DocAssertion("Int64(nil)"), DocAssertion("0")),
			assert(DocAssertion("Int64([][]int{})"), DocAssertion("0")),
			assert(DocAssertion("Int64((chan string)(nil))"), DocAssertion("0")),
			assert(DocAssertion("Int64((*interface{})(nil))"), DocAssertion("0")),
			assert(DocAssertion("Int64((*interface {})(0x01))"), DocAssertion("0")),
			assert(DocAssertion("Int64((**interface {})(0x1))"), DocAssertion("0")),
		),
	)
	section("Value",
		group(`Value is a convenience struct for performing Conversion. It has a
			single field V of interface{} type which is passed to the associated
			conversion functions.`,
			assert(DA(`func (v Value) Bool() bool { return Bool(v.V) }`)),
		),
		group(`This means you may wrap any value with Value{...} for conversions.`,
			assert(DA(`v := Value{"12345.6789"}`)),
			assert(DA(`v.Int64()`), DA(String(int64(myValue.Int64())))),
			assert(DA(`v.Float64()`), DA(String(float64(myValue.Float64())))),
			assert(DA(`v.Float32()`), DA(String(float32(myValue.Float32())))),
		),
	)
	asserts("time",
		assert(time.Duration(3*time.Second), NowAssertion{Offset: time.Second * 3}),
	)
	asserts("numeric",
		asserts("float32",
			assert(float32(1.0), true, numerics(1),
				Float32Assertion{float32(1.0)},
				Float64Assertion{float64(1.0)}),
			assert(float32(0), false, numerics(0),
				Float32Assertion{float32(0)},
				Float64Assertion{float64(0)}),
			assert(float32(-1.0), true, numerics(-1),
				Float32Assertion{float32(-1.0)},
				Float64Assertion{float64(-1.0)}),
		),
		asserts("float64",
			assert(float64(1.0), true, numerics(1),
				Float32Assertion{float32(1.0)},
				Float64Assertion{float64(1.0)}),
			assert(float64(0), false, numerics(0),
				Float32Assertion{float32(0)},
				Float64Assertion{float64(0)}),
			assert(float64(-1.0), true, numerics(-1),
				Float32Assertion{float32(-1.0)},
				Float64Assertion{float64(-1.0)}),
		),
		asserts("Complex64",
			assert(complex64(-1), numerics(-1)),
			assert(complex64(0), numerics(0)),
			assert(complex64(1), numerics(1)),
		),
		asserts("Complex128",
			assert(complex128(-1), true, "(-1+0i)", numerics(-1), durations(-1)),
			assert(complex128(0), false, "(0+0i)", numerics(0), durations(0)),
			assert(complex128(1), true, "(1+0i)", numerics(1), durations(1)),
		),
		asserts("int",
			assert(int(-1), true, numerics(-1)),
			assert(int(0), false, numerics(0)),
			assert(int(1), true, numerics(1)),
		),
		asserts("int8",
			assert(int8(-1), true, numerics(-1)),
			assert(int8(0), false, numerics(0)),
			assert(int8(1), true, numerics(1)),
		),
		asserts("int16",
			assert(int16(-1), true, numerics(-1)),
			assert(int16(0), false, numerics(0)),
			assert(int16(1), true, numerics(1)),
		),
		asserts("int32",
			assert(int32(-1), true, numerics(-1)),
			assert(int32(0), false, numerics(0)),
			assert(int32(1), true, numerics(1)),
		),
		asserts("int64",
			assert(int64(-1), true, numerics(-1)),
			assert(int64(0), false, numerics(0)),
			assert(int64(1), true, numerics(1)),
		),
		asserts("uint",
			assert(uint(0), false, numerics(0)),
			assert(uint(1), true, numerics(1)),
		),
		asserts("uint8",
			assert(uint8(0), false, numerics(0)),
			assert(uint8(1), true, numerics(1)),
		),
		asserts("uint16",
			assert(uint16(0), false, numerics(0)),
			assert(uint16(1), true, numerics(1)),
		),
		asserts("uint32",
			assert(uint32(0), false, numerics(0)),
			assert(uint32(1), true, numerics(1)),
		),
		asserts("uint64",
			assert(uint64(0), false, numerics(0)),
			assert(uint64(1), true, numerics(1)),
		),
	)
}
