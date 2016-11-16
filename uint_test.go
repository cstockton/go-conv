package conv

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

type testUint64Converter uint64

func (t testUint64Converter) Uint64() (uint64, error) {
	return uint64(t) + 5, nil
}

func init() {
	type ulyUint uint
	type ulyUint8 uint8
	type ulyUint16 uint8
	type ulyUint32 uint8
	type ulyUint64 uint64

	exp := func(e uint, e8 uint8, e16 uint16, e32 uint32, e64 uint64) []Expecter {
		return []Expecter{Exp{e}, Exp{e8}, Exp{e16}, Exp{e32}, Exp{e64}}
	}
	experrs := func(s string) []Expecter {
		return []Expecter{
			experr(uint(0), s), experr(uint8(0), s), experr(uint16(0), s),
			experr(uint32(0), s), experr(uint64(0), s)}
	}

	// basics
	assert(0, exp(0, 0, 0, 0, 0))
	assert(1, exp(1, 1, 1, 1, 1))
	assert(false, exp(0, 0, 0, 0, 0))
	assert(true, exp(1, 1, 1, 1, 1))
	assert("false", exp(0, 0, 0, 0, 0))
	assert("true", exp(1, 1, 1, 1, 1))

	// test length kinds
	assert([]string{"one", "two"}, exp(2, 2, 2, 2, 2))
	assert(map[int]string{1: "one", 2: "two"}, exp(2, 2, 2, 2, 2))

	// test implements Uint64(uint64, error)
	assert(testUint64Converter(5), exp(10, 10, 10, 10, 10))

	// max bounds
	assert(math.MaxUint8, exp(math.MaxUint8, math.MaxUint8, math.MaxUint8,
		math.MaxUint8, math.MaxUint8))
	assert(math.MaxUint16, exp(math.MaxUint16, math.MaxUint8, math.MaxUint16,
		math.MaxUint16, math.MaxUint16))
	assert(math.MaxUint32, exp(math.MaxUint32, math.MaxUint8, math.MaxUint16,
		math.MaxUint32, math.MaxUint32))
	assert(uint64(math.MaxUint64), exp(uint(mathMaxUint), math.MaxUint8,
		math.MaxUint16, math.MaxUint32, uint64(math.MaxUint64)))

	// min bounds
	assert(math.MinInt8, exp(0, 0, 0, 0, 0))
	assert(math.MinInt16, exp(0, 0, 0, 0, 0))
	assert(math.MinInt32, exp(0, 0, 0, 0, 0))
	assert(int64(math.MinInt64), exp(0, 0, 0, 0, 0))

	// perms of various type
	for n := uint8(0); n < math.MaxUint8; n += 0xB {
		i := n

		// uints
		if n < 1 {
			i = 0
		} else {
			assert(uint(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
			assert(uint8(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
			assert(uint16(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
			assert(uint32(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
			assert(uint64(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		}

		// underlying
		assert(ulyUint(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		assert(ulyUint8(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		assert(ulyUint16(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		assert(ulyUint32(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		assert(ulyUint64(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))

		// implements
		if i < math.MaxUint8-5 {
			assert(testUint64Converter(i),
				uint(i+5), uint8(i+5), uint16(i+5), uint32(i+5), uint64(i+5))
			assert(testUint64Converter(ulyUint(i)),
				uint(i+5), uint8(i+5), uint16(i+5), uint32(i+5), uint64(i+5))
		}

		// ints
		if i < math.MaxInt8 {
			assert(int(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
			assert(int8(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
			assert(int16(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
			assert(int32(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
			assert(int64(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		}

		// floats
		assert(float32(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		assert(float64(i), uint(i), uint8(i), uint16(i), uint32(i), uint64(i))

		// complex
		assert(complex(float32(i), 0),
			uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		assert(complex(float64(i), 0),
			uint(i), uint8(i), uint16(i), uint32(i), uint64(i))

		// from string int
		assert(fmt.Sprintf("%d", i),
			uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
		assert(testStringConverter(fmt.Sprintf("%d", i)),
			uint(i), uint8(i), uint16(i), uint32(i), uint64(i))

		// from string float form
		assert(fmt.Sprintf("%d.0", i),
			uint(i), uint8(i), uint16(i), uint32(i), uint64(i))
	}

	assert(nil, experrs(`cannot convert <nil> (type <nil>) to `))
	assert("foo", experrs(` "foo" (type string) `))
	assert(struct{}{}, experrs(`cannot convert struct {}{} (type struct {}) to `))
}

func TestUint(t *testing.T) {
	var c Conv
	t.Run("Uint", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Uint, func(a *Assertion, e Expecter) {
			res, err := c.Uint(a.From)
			if res != Uint(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Uint conversions")
		}
	})
	t.Run("Uint8", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Uint8, func(a *Assertion, e Expecter) {
			res, err := c.Uint8(a.From)
			if res != Uint8(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Uint8 conversions")
		}
	})
	t.Run("Uint16", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Uint16, func(a *Assertion, e Expecter) {
			res, err := c.Uint16(a.From)
			if res != Uint16(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Uint16 conversions")
		}
	})
	t.Run("Uint32", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Uint32, func(a *Assertion, e Expecter) {
			res, err := c.Uint32(a.From)
			if res != Uint32(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Uint32 conversions")
		}
	})
	t.Run("Uint64", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Uint64, func(a *Assertion, e Expecter) {
			res, err := c.Uint64(a.From)
			if res != Uint64(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Uint64 conversions")
		}
	})
}
