package conv

import (
	"math"
	"math/cmplx"
	"reflect"
	"testing"
	"time"
)

type testBoolConverter bool

func (t testBoolConverter) Bool() (bool, error) {
	return !bool(t), nil
}

type testStringConverter string

func (t testStringConverter) String() (string, error) {
	return string(t) + "Tested", nil
}

func init() {

	// bools
	{

		// strings: truthy
		trueStrings := []string{
			"1", "t", "T", "true", "True", "TRUE", "y", "Y", "yes", "Yes", "YES"}
		for _, truthy := range trueStrings {
			assert(truthy, true)
			assert(testStringConverter(truthy), true)
		}

		// strings: falsy
		falseStrings := []string{
			"0", "f", "F", "false", "False", "FALSE", "n", "N", "no", "No", "NO"}
		for _, falsy := range falseStrings {
			assert(falsy, false)
			assert(testStringConverter(falsy), false)
		}

		// numerics: true
		for _, i := range []int{-1, 1} {
			assert(int(i), true)
			assert(int8(i), true)
			assert(int16(i), true)
			assert(int32(i), true)
			assert(int64(i), true)
			assert(uint(i), true)
			assert(uint8(i), true)
			assert(uint16(i), true)
			assert(uint32(i), true)
			assert(uint64(i), true)
			assert(float32(i), true)
			assert(float64(i), true)
			assert(complex(float32(i), 0), true)
			assert(complex(float64(i), 0), true)
		}

		// int/uint: false
		assert(int(0), false)
		assert(int8(0), false)
		assert(int16(0), false)
		assert(int32(0), false)
		assert(int64(0), false)
		assert(uint(0), false)
		assert(uint8(0), false)
		assert(uint16(0), false)
		assert(uint32(0), false)
		assert(uint64(0), false)

		// float: NaN and 0 are false.
		assert(float32(math.NaN()), false)
		assert(math.NaN(), false)
		assert(float32(0), false)
		assert(float64(0), false)

		// complex: NaN and 0 are false.
		assert(complex64(cmplx.NaN()), false)
		assert(cmplx.NaN(), false)
		assert(complex(float32(0), 0), false)
		assert(complex(float64(0), 0), false)

		// time
		assert(time.Time{}, false)
		assert(time.Now(), true)

		// bool
		assert(false, false)
		assert(true, true)

		// underlying bool
		type ulyBool bool
		assert(ulyBool(false), false)
		assert(ulyBool(true), true)

		// implements bool converter
		assert(testBoolConverter(false), true)
		assert(testBoolConverter(true), false)

		// test length kinds
		assert([]string{"one", "two"}, true)
		assert(map[int]string{1: "one", 2: "two"}, true)
		assert([]string{}, false)
		assert([]string(nil), false)

		// errors
		assert(nil, experr(false, `cannot convert <nil> (type <nil>) to bool`))
		assert("foo", experr(false, `cannot parse "foo" (type string) as bool`))
		assert("tooLong", experr(
			false, `cannot parse type string with len 7 as bool`))
		assert(struct{}{}, experr(
			false, `cannot convert struct {}{} (type struct {}) to `))
	}

	// strings
	{

		// basic
		assert(`hello`, `hello`)
		assert(``, ``)
		assert([]byte(`hello`), `hello`)
		assert([]byte(``), ``)

		// ptr indirection
		assert(new(string), ``)
		assert(new([]byte), ``)

		// underlying string
		type ulyString string
		assert(ulyString(`hello`), `hello`)
		assert(ulyString(``), ``)

		// implements string converter
		assert(testStringConverter(`hello`), `helloTested`)
		assert(testStringConverter(`hello`), `helloTested`)

		// errors
		assert(nil, experr(false, `cannot convert <nil> (type <nil>) to bool`))
		assert("foo", experr(false, `cannot parse "foo" (type string) as bool`))
		assert("tooLong", experr(
			false, `cannot parse type string with len 7 as bool`))
		assert(struct{}{}, experr(
			false, `cannot convert struct {}{} (type struct {}) to `))
	}
}

func TestString(t *testing.T) {
	var c Conv
	t.Run("String", func(t *testing.T) {
		if n := assertions.EachOf(reflect.String, func(a *Assertion, e Expecter) {
			if err := e.Expect(c.String(a.From)); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for String conversions")
		}
	})
}

func TestBool(t *testing.T) {
	var c Conv
	t.Run("Bool", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Bool, func(a *Assertion, e Expecter) {
			if err := e.Expect(c.Bool(a.From)); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Bool conversions")
		}
	})
	t.Run("convNumToBool", func(t *testing.T) {
		var val reflect.Value
		if got, ok := c.convNumToBool(0, val); ok || got {
			t.Fatal("expected failure")
		}
	})
}
