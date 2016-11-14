package conv

import (
	"math"
	"math/cmplx"
	"reflect"
	"testing"
	"time"
)

type testTimeConverter time.Time

func (t testTimeConverter) Time() (time.Time, error) {
	return time.Time(t).Add(time.Minute), nil
}

type testDurationConverter time.Duration

func (d testDurationConverter) Duration() (time.Duration, error) {
	return time.Duration(d) + time.Minute, nil
}

func init() {
	var (
		dZero       time.Duration
		t2006               = time.Date(2006, time.January, 2, 15, 4, 5, 999999999, time.UTC)
		tNow                = time.Now()
		tNowMinus1m         = tNow.Add(-time.Minute)
		tNowPlus1m          = tNow.Add(time.Minute)
		d42ns               = time.Nanosecond * 42
		d2m                 = time.Minute * 2
		d34s                = time.Second * 34
		d567ms              = time.Millisecond * 567
		d234567             = d2m + d34s + d567ms
		d154567f32  float32 = 154.567
		d154567f64          = 154.567
		dMAX                = time.Duration(math.MaxInt64)
	)

	// durations
	{

		// strings
		assert("2m34.567s", d234567)
		assert("-2m34.567s", -d234567)
		assert("154.567", d234567)
		assert("-154.567", -d234567)
		assert("120", d2m)

		// durations
		assert(d234567, d234567)
		assert(dZero, dZero)
		assert(new(time.Duration), dZero)

		// time
		assert(tNowPlus1m, DurationExp{-time.Minute, time.Second})
		assert(tNowMinus1m, DurationExp{time.Minute, time.Second})

		// underlying
		type ulyDuration time.Duration
		assert(ulyDuration(time.Second), time.Second)
		assert(ulyDuration(-time.Second), -time.Second)

		// implements converter
		assert(testDurationConverter(time.Second), time.Second+time.Minute)
		assert(testDurationConverter(-time.Second), time.Minute-time.Second)

		// numerics
		assert(int(42), d42ns)
		assert(int8(42), d42ns)
		assert(int16(42), d42ns)
		assert(int32(42), d42ns)
		assert(int64(42), d42ns)
		assert(uint(42), d42ns)
		assert(uint8(42), d42ns)
		assert(uint16(42), d42ns)
		assert(uint32(42), d42ns)
		assert(uint64(42), d42ns)

		// floats
		assert(d154567f32, DurationExp{d234567, time.Millisecond})
		assert(d154567f64, d234567)
		assert(math.NaN(), dZero)
		assert(math.Inf(1), dZero)
		assert(math.Inf(-1), dZero)

		// complex
		assert(complex(d154567f32, 0), DurationExp{d234567, time.Millisecond})
		assert(complex(d154567f64, 0), d234567)
		assert(cmplx.NaN(), dZero)
		assert(cmplx.Inf(), dZero)

		// overflow
		assert(uint64(math.MaxUint64), dMAX)

		// errors
		assert(nil, experr(dZero, `cannot convert <nil> (type <nil>) to time.Duration`))
		assert("foo", experr(dZero, `cannot convert "foo" (type string) to time.Duration`))
		assert("tooLong", experr(
			dZero, `cannot convert "tooLong" (type string) to time.Duration`))
		assert(struct{}{}, experr(
			dZero, `cannot convert struct {}{} (type struct {}) to `))
		assert([]string{"1s"}, experr(
			dZero, `cannot convert []string{"1s"} (type []string) to `))
		assert([]string{}, experr(
			dZero, `cannot convert []string{} (type []string) to `))
	}

	// times
	{

		// basic
		assert(time.Time{}, time.Time{})
		assert(new(time.Time), time.Time{})
		assert(t2006, t2006)

		// strings
		fmts := []string{
			"02 Jan 06 15:04:05",
			"2 Jan 2006 15:04:05",
			"2 Jan 2006 15:04:05 -0700 (UTC)",
			"02 Jan 2006 15:04 UTC",
			"02 Jan 2006 15:04:05 UTC",
			"02 Jan 2006 15:04:05 -0700 (UTC)",
			"Mon, 2 Jan  15:04:05 UTC 2006",
			"Mon, 2 Jan 15:04:05 UTC 2006",
			"Mon, 02 Jan 2006 15:04:05",
			"Mon, 02 Jan 2006 15:04:05 (UTC)",
			"Mon, 2 Jan 2006 15:04:05",
		}
		for _, s := range fmts {
			assert(s, TimeExp{Moment: t2006.Truncate(time.Minute), Truncate: time.Minute})
		}

		// underlying
		type ulyTime time.Time
		assert(ulyTime(t2006), t2006)
		assert(ulyTime(t2006), t2006)

		// implements converter
		assert(testTimeConverter(t2006), t2006.Add(time.Minute))

		// embedded time
		type embedTime struct{ time.Time }
		assert(embedTime{t2006}, t2006)

		// errors
		assert(nil, experr(emptyTime, `cannot convert <nil> (type <nil>) to time.Time`))
		assert("foo", experr(emptyTime, `cannot convert "foo" (type string) to time.Time`))
		assert("tooLong", experr(
			emptyTime, `cannot convert "tooLong" (type string) to time.Time`))
		assert(struct{}{}, experr(
			emptyTime, `cannot convert struct {}{} (type struct {}) to `))
		assert([]string{"1s"}, experr(
			emptyTime, `cannot convert []string{"1s"} (type []string) to `))
		assert([]string{}, experr(
			emptyTime, `cannot convert []string{} (type []string) to `))
	}
}

func TestTime(t *testing.T) {
	var c Conv
	t.Run("Time", func(t *testing.T) {
		if n := assertions.EachOf(TimeKind, func(a *Assertion, e Expecter) {
			if err := e.Expect(c.Time(a.From)); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Time conversions")
		}
	})
	t.Run("timeFromString", func(t *testing.T) {
		if _, ok := timeFromString(""); ok {
			t.Fatal("expected timeFromString to return false on 0 len str")
		}
	})
}

func TestDuration(t *testing.T) {
	var c Conv
	t.Run("Duration", func(t *testing.T) {
		if n := assertions.EachOf(DurationKind, func(a *Assertion, e Expecter) {
			if err := e.Expect(c.Duration(a.From)); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Duration conversions")
		}
	})
	t.Run("convNumToDuration", func(t *testing.T) {
		var val reflect.Value
		if _, ok := c.convNumToDuration(0, val); ok {
			t.Fatal("expected convNumToDuration to return false on invalid kind")
		}
	})
}
