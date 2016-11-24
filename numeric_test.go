package conv

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestFloat(t *testing.T) {
	var c Conv
	t.Run("Float32", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Float32, func(a *Assertion, e Expecter) {
			res, err := c.Float32(a.From)
			if res != Float32(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Float32 conversions")
		}
	})
	t.Run("Float64", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Float64, func(a *Assertion, e Expecter) {
			res, err := c.Float64(a.From)
			if res != Float64(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Float64 conversions")
		}
	})
}

func TestInt(t *testing.T) {
	var c Conv
	t.Run("Int", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Int, func(a *Assertion, e Expecter) {
			res, err := c.Int(a.From)
			if res != Int(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Int conversions")
		}
	})
	t.Run("Int8", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Int8, func(a *Assertion, e Expecter) {
			res, err := c.Int8(a.From)
			if res != Int8(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Int8 conversions")
		}
	})
	t.Run("Int16", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Int16, func(a *Assertion, e Expecter) {
			res, err := c.Int16(a.From)
			if res != Int16(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Int16 conversions")
		}
	})
	t.Run("Int32", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Int32, func(a *Assertion, e Expecter) {
			res, err := c.Int32(a.From)
			if res != Int32(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Int32 conversions")
		}
	})
	t.Run("Int64", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Int64, func(a *Assertion, e Expecter) {
			res, err := c.Int64(a.From)
			if res != Int64(a.From) {
				t.Fatalf("result drift between func and Conv")
			}
			if err = e.Expect(res, err); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Int64 conversions")
		}
	})
}

func BenchmarkInt(b *testing.B) {
	var c Conv
	b.Run("string to int64", func(b *testing.B) {
		l := len(strToInt64)
		b.Run("Conv", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for z := 0; z < l; z++ {
					v, err := c.Int64(strToInt64[z].from)
					if err != nil {
						b.Error(err)
					}
					if strToInt64[z].to != v {
						b.Errorf("(%T) %[1]v != %v (%[2]T)", strToInt64[z].to, v)
					}
				}
			}
		})
		b.Run("Stdlib", func(b *testing.B) {
			parseInt := func(from interface{}) (int64, error) {
				if T, ok := from.(string); ok {
					return strconv.ParseInt(T, 10, 0)
				}
				b.Fatal("expected string")
				return 0, nil
			}
			for i := 0; i < b.N; i++ {
				for z := 0; z < l; z++ {
					v, err := parseInt(strToInt64[z].from)
					if err != nil {
						b.Error(err)
					}
					if strToInt64[z].to != v {
						b.Errorf("(%T) %[1]v != %v (%[2]T)", strToInt64[z].to, v)
					}
				}
			}
		})
		b.Run("StdlibTyped", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for z := 0; z < l; z++ {
					v, err := strconv.ParseInt(strToInt64[z].from, 10, 0)
					if err != nil {
						b.Error(err)
					}
					if strToInt64[z].to != v {
						b.Errorf("(%T) %[1]v != %v (%[2]T)", strToInt64[z].to, v)
					}
				}
			}
		})
	})
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

var strToInt64 = [42]struct {
	from string
	to   int64
}{
	{"0", 0},
	{"-0", 0},
	{"1", 1},
	{"-1", -1},
	{"12", 12},
	{"-12", -12},
	{"123", 123},
	{"-123", -123},
	{"1234", 1234},
	{"-1234", -1234},
	{"12345", 12345},
	{"-12345", -12345},
	{"123456", 123456},
	{"-123456", -123456},
	{"1234567", 1234567},
	{"-1234567", -1234567},
	{"12345678", 12345678},
	{"-12345678", -12345678},
	{"123456789", 123456789},
	{"-123456789", -123456789},
	{"1234567890", 1234567890},
	{"-1234567890", -1234567890},
	{"12345678901", 12345678901},
	{"-12345678901", -12345678901},
	{"123456789012", 123456789012},
	{"-123456789012", -123456789012},
	{"1234567890123", 1234567890123},
	{"-1234567890123", -1234567890123},
	{"12345678901234", 12345678901234},
	{"-12345678901234", -12345678901234},
	{"123456789012345", 123456789012345},
	{"-123456789012345", -123456789012345},
	{"1234567890123456", 1234567890123456},
	{"-1234567890123456", -1234567890123456},
	{"12345678901234567", 12345678901234567},
	{"-12345678901234567", -12345678901234567},
	{"123456789012345678", 123456789012345678},
	{"-123456789012345678", -123456789012345678},
	{"1234567890123456789", 1234567890123456789},
	{"-1234567890123456789", -1234567890123456789},
	{"9223372036854775807", 1<<63 - 1},
	{"-9223372036854775808", -1 << 63},
}

type testFloat64Converter float64

func (t testFloat64Converter) Float64() (float64, error) {
	return float64(t) + 5, nil
}

type testInt64Converter int64

func (t testInt64Converter) Int64() (int64, error) {
	return int64(t) + 5, nil
}

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

func init() {
	type ulyInt int
	type ulyInt8 int8
	type ulyInt16 int8
	type ulyInt32 int8
	type ulyInt64 int64

	exp := func(e int, e8 int8, e16 int16, e32 int32, e64 int64) []Expecter {
		return []Expecter{Exp{e}, Exp{e8}, Exp{e16}, Exp{e32}, Exp{e64}}
	}
	experrs := func(s string) []Expecter {
		return []Expecter{
			experr(int(0), s), experr(int8(0), s), experr(int16(0), s),
			experr(int32(0), s), experr(int64(0), s)}
	}

	// basics
	assert(-1, exp(-1, -1, -1, -1, -1))
	assert(0, exp(0, 0, 0, 0, 0))
	assert(1, exp(1, 1, 1, 1, 1))
	assert(false, exp(0, 0, 0, 0, 0))
	assert(true, exp(1, 1, 1, 1, 1))
	assert("false", exp(0, 0, 0, 0, 0))
	assert("true", exp(1, 1, 1, 1, 1))

	// test length kinds
	assert([]string{"one", "two"}, 2, 2, 2, 2, 2)
	assert(map[int]string{1: "one", 2: "two"}, 2, 2, 2, 2, 2)

	// test implements Int64(int64, error)
	assert(testInt64Converter(5), 10, 10, 10, 10, 10)

	// overflow
	assert(uint64(math.MaxUint64), exp(int(mathMaxInt), math.MaxInt8,
		math.MaxInt16, math.MaxInt32, math.MaxInt64))

	// underflow
	assert(int64(math.MinInt64), exp(int(mathMinInt), math.MinInt8, math.MinInt16,
		math.MinInt32, math.MinInt64))

	// max bounds
	assert(math.MaxInt8, exp(math.MaxInt8, math.MaxInt8, math.MaxInt8,
		math.MaxInt8, math.MaxInt8))
	assert(math.MaxInt16, exp(math.MaxInt16, math.MaxInt8, math.MaxInt16,
		math.MaxInt16, math.MaxInt16))
	assert(math.MaxInt32, exp(math.MaxInt32, math.MaxInt8, math.MaxInt16,
		math.MaxInt32, math.MaxInt32))
	assert(math.MaxInt64, exp(int(mathMaxInt), math.MaxInt8, math.MaxInt16,
		math.MaxInt32, math.MaxInt64))

	// min bounds
	assert(math.MinInt8, exp(math.MinInt8, math.MinInt8, math.MinInt8,
		math.MinInt8, math.MinInt8))
	assert(math.MinInt16, exp(math.MinInt16, math.MinInt8, math.MinInt16,
		math.MinInt16, math.MinInt16))
	assert(math.MinInt32, exp(math.MinInt32, math.MinInt8, math.MinInt16,
		math.MinInt32, math.MinInt32))
	assert(int64(math.MinInt64), exp(int(mathMinInt), math.MinInt8, math.MinInt16,
		math.MinInt32, math.MinInt64))

	// perms of various type
	for i := math.MinInt8; i < math.MaxInt8; i += 0xB {

		// uints
		if i > 0 {
			assert(uint(i), int(i), int8(i), int16(i), int32(i), int64(i))
			assert(uint8(i), int(i), int8(i), int16(i), int32(i), int64(i))
			assert(uint16(i), int(i), int8(i), int16(i), int32(i), int64(i))
			assert(uint32(i), int(i), int8(i), int16(i), int32(i), int64(i))
			assert(uint64(i), int(i), int8(i), int16(i), int32(i), int64(i))
		}

		// underlying
		assert(ulyInt(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(ulyInt8(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(ulyInt16(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(ulyInt32(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(ulyInt64(i), int(i), int8(i), int16(i), int32(i), int64(i))

		// implements
		if i < math.MaxInt8-5 {
			assert(testInt64Converter(i),
				int(i+5), int8(i+5), int16(i+5), int32(i+5), int64(i+5))
			assert(testInt64Converter(ulyInt(i)),
				int(i+5), int8(i+5), int16(i+5), int32(i+5), int64(i+5))
		}

		// ints
		assert(int(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(int8(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(int16(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(int32(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(int64(i), int(i), int8(i), int16(i), int32(i), int64(i))

		// floats
		assert(float32(i), int(i), int8(i), int16(i), int32(i), int64(i))
		assert(float64(i), int(i), int8(i), int16(i), int32(i), int64(i))

		// complex
		assert(complex(float32(i), 0),
			int(i), int8(i), int16(i), int32(i), int64(i))
		assert(complex(float64(i), 0),
			int(i), int8(i), int16(i), int32(i), int64(i))

		// from string int
		assert(fmt.Sprintf("%d", i),
			int(i), int8(i), int16(i), int32(i), int64(i))
		assert(testStringConverter(fmt.Sprintf("%d", i)),
			int(i), int8(i), int16(i), int32(i), int64(i))

		// from string float form
		assert(fmt.Sprintf("%d.0", i),
			int(i), int8(i), int16(i), int32(i), int64(i))
	}

	assert("foo", experrs(`"foo" (type string) `))
	assert(struct{}{}, experrs(`cannot convert struct {}{} (type struct {}) to `))
	assert(nil, experrs(`cannot convert <nil> (type <nil>) to `))
}

func init() {
	type ulyFloat32 float32
	type ulyFloat64 float64

	exp := func(e32 float32, e64 float64) []Expecter {
		return []Expecter{Float32Exp{e32}, Float64Exp{e64}}
	}
	experrs := func(s string) []Expecter {
		return []Expecter{experr(float32(0), s), experr(float64(0), s)}
	}

	// basics
	assert(0, exp(0, 0))
	assert(1, exp(1, 1))
	assert(false, exp(0, 0))
	assert(true, exp(1, 1))
	assert("false", exp(0, 0))
	assert("true", exp(1, 1))

	// test length kinds
	assert([]string{"one", "two"}, exp(2, 2))
	assert(map[int]string{1: "one", 2: "two"}, exp(2, 2))

	// test implements Float64(float64, error)
	assert(testFloat64Converter(5), exp(10, 10))

	// max bounds
	assert(math.MaxFloat32, exp(math.MaxFloat32, math.MaxFloat32))
	assert(math.MaxFloat64, exp(math.MaxFloat32, math.MaxFloat64))

	// min bounds
	assert(-math.MaxFloat32, exp(-math.MaxFloat32, -math.MaxFloat32))
	assert(-math.MaxFloat64, exp(-math.MaxFloat32, -math.MaxFloat64))

	// ints
	assert(int(10), exp(10, float64(10)))
	assert(int8(10), exp(10, float64(10)))
	assert(int16(10), exp(10, float64(10)))
	assert(int32(10), exp(10, float64(10)))
	assert(int64(10), exp(10, float64(10)))

	// uints
	assert(uint(10), exp(10, float64(10)))
	assert(uint8(10), exp(10, float64(10)))
	assert(uint16(10), exp(10, float64(10)))
	assert(uint32(10), exp(10, float64(10)))
	assert(uint64(10), exp(10, float64(10)))

	// perms of various type
	for i := float32(-3.0); i < 3.0; i += .5 {

		// underlying
		assert(ulyFloat32(i), exp(i, float64(i)))
		assert(ulyFloat64(i), exp(i, float64(i)))

		// implements
		assert(testFloat64Converter(i), exp(i+5, float64(i+5)))
		assert(testFloat64Converter(ulyFloat64(i)), exp(i+5, float64(i+5)))

		// floats
		assert(float32(i), exp(i, float64(i)))
		assert(float64(i), exp(i, float64(i)))

		// complex
		assert(complex(float32(i), 0), exp(i, float64(i)))
		assert(complex(float64(i), 0), exp(i, float64(i)))

		// from string int
		assert(fmt.Sprintf("%#v", i), exp(i, float64(i)))
		assert(testStringConverter(fmt.Sprintf("%#v", i)), exp(i, float64(i)))

		// from string float form
		assert(fmt.Sprintf("%#v", i), exp(i, float64(i)))
	}

	assert("foo", experrs(`cannot convert "foo" (type string) to `))
	assert(struct{}{}, experrs(`cannot convert struct {}{} (type struct {}) to `))
	assert(nil, experrs(`cannot convert <nil> (type <nil>) to `))
}
