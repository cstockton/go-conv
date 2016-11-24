package conv

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	iter "github.com/cstockton/go-iter"
)

func TestMap(t *testing.T) {
	t.Run("FromSlice", func(t *testing.T) {
		t.Run("ToStringInt64", func(t *testing.T) {
			var c Conv
			into := make(map[string]int64)
			err := c.Map(into, []string{"12", "345", "6789"})
			if err != nil {
				t.Error(err)
			}

			exp := map[string]int64{"0": 12, "1": 345, "2": 6789}
			if !reflect.DeepEqual(exp, into) {
				t.Fatalf("exp (%T) --> %[1]v != %v <-- (%[2]T) got", exp, into)
			}
			into = make(map[string]int64)
			err = Map(into, []string{"12", "345", "6789"})
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(exp, into) {
				t.Fatalf("exp (%T) --> %[1]v != %v <-- (%[2]T) got", exp, into)
			}

			err = Map(nil, []string{"12", "345", "6789"})
			if err == nil {
				t.Error("expected non-nil err")
			}
		})
		t.Run("ToStringInt64P", func(t *testing.T) {
			var c Conv
			i1, i2, i3 := new(int64), new(int64), new(int64)
			*i1, *i2, *i3 = 12, 345, 6789
			exp := map[string]*int64{"0": i1, "1": i2, "2": i3}

			into := make(map[string]*int64)
			err := c.Map(into, []string{"12", "345", "6789"})
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(exp, into) {
				t.Fatalf("exp (%T) --> %[1]v != %v <-- (%[2]T) got", exp, into)
			}
			into = make(map[string]*int64)
			err = Map(into, []string{"12", "345", "6789"})
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(exp, into) {
				t.Fatalf("exp (%T) --> %[1]v != %v <-- (%[2]T) got", exp, into)
			}
		})
	})
}

func TestMapGen(t *testing.T) {
	var c Conv

	triggerAt, count, triggerErr := 0, 0, errors.New("triggerErr")
	cErr := errHookConv{c, func(from interface{}, err error) error {
		count++
		if triggerAt == count {
			return triggerErr
		}
		return err
	}}

	for _, mapGenTest := range mapGenTests {
		err := c.Map(mapGenTest.into, mapGenTest.from)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(mapGenTest.exp, mapGenTest.into) {
			t.Logf("from (%T) --> %[1]v", mapGenTest.from)
			t.Fatalf("\nexp (%T) --> %[1]v\ngot (%[2]T) --> %[2]v",
				mapGenTest.exp, mapGenTest.into)
		}

		refVal := reflect.ValueOf(mapGenTest.from)
		keys := refVal.MapKeys()
		key, val := keys[0].Interface(), refVal.MapIndex(keys[0]).Interface()
		fn := mapIterFn(cErr, mapGenTest.into)
		el := iter.NewPair(nil, key, val, nil)
		elErr := iter.NewPair(nil, key, val, triggerErr)

		// Trigger element err check
		if err := fn(elErr); err != triggerErr {
			t.Errorf("expected element err in mapIterFn %T", mapGenTest.into)
		}

		// Trigger key err check
		triggerAt, count = 1, 0
		if err := fn(el); err != triggerErr {
			t.Errorf("expected key err in mapIterFn %T", mapGenTest.into)
		}

		// Trigger val err check
		triggerAt, count = 2, 0
		if err := fn(el); err != triggerErr {
			t.Errorf("expected val err in mapIterFn %T", mapGenTest.into)
		}
	}
}

// Summary: Not much of a tax here, about 2x as slow.
//
// BenchmarkMap/<slice size>/<from> to <to>/Conv:
//   Measures the most convenient form of conversion using this library.
//
// BenchmarkSlice/<slice size>/<from> to <to>/Conv:
//   Measures using the library only for the conversion, looping for apending.
//
// BenchmarkSlice/<slice size>/<from> to <to>/Conv:
//   Measures not using this library at all, pure Go implementation.
//
// BenchmarkMap/Length(1024)/[]string_to_map[int]string/Conv-24    	    1000	   1321364 ns/op
// BenchmarkMap/Length(1024)/[]string_to_map[int]string/LoopConv-24         	    2000	    896001 ns/op
// BenchmarkMap/Length(1024)/[]string_to_map[int]string/LoopStdlib-24       	    2000	    652117 ns/op
// BenchmarkMap/Length(64)/[]string_to_map[int]string/Conv-24               	   20000	     74431 ns/op
// BenchmarkMap/Length(64)/[]string_to_map[int]string/LoopConv-24           	   20000	     56702 ns/op
// BenchmarkMap/Length(64)/[]string_to_map[int]string/LoopStdlib-24         	   30000	     44191 ns/op
// BenchmarkMap/Length(16)/[]string_to_map[int]string/Conv-24               	  100000	     18422 ns/op
// BenchmarkMap/Length(16)/[]string_to_map[int]string/LoopConv-24           	  100000	     14193 ns/op
// BenchmarkMap/Length(16)/[]string_to_map[int]string/LoopStdlib-24         	  200000	     10021 ns/op
// BenchmarkMap/Length(4)/[]string_to_map[int]string/Conv-24                	  300000	      4402 ns/op
// BenchmarkMap/Length(4)/[]string_to_map[int]string/LoopConv-24            	  500000	      2783 ns/op
// BenchmarkMap/Length(4)/[]string_to_map[int]string/LoopStdlib-24          	 1000000	      1986 ns/op
func BenchmarkMap(b *testing.B) {
	var c Conv

	for _, num := range []int{1024, 64, 16, 4} {
		num := num

		// slow down is really tolerable, only a factor of 1-3 tops
		b.Run(fmt.Sprintf("Length(%d)", num), func(b *testing.B) {

			b.Run("[]string to map[int]string", func(b *testing.B) {
				strs := make([]string, num)
				for n := 0; n < num; n++ {
					strs[n] = fmt.Sprintf("%v00", n)
				}
				b.ResetTimer()

				b.Run("Conv", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						into := make(map[string]int64)
						err := c.Map(into, strs)
						if err != nil {
							b.Error(err)
						}
						if len(into) != num {
							b.Error("bad impl")
						}
					}
				})
				b.Run("LoopConv", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						into := make(map[string]int64)

						for seq, s := range strs {
							k, err := c.String(seq)
							if err != nil {
								b.Error(err)
							}
							v, err := c.Int64(s)
							if err != nil {
								b.Error(err)
							}
							into[k] = v
						}
						if len(into) != num {
							b.Error("bad impl")
						}
					}
				})
				b.Run("LoopStdlib", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						into := make(map[string]int64)

						for seq, s := range strs {
							k := fmt.Sprintf("%v", seq)
							v, err := strconv.ParseInt(s, 10, 0)
							if err != nil {
								b.Error(err)
							}
							into[k] = v
						}
						if len(into) != num {
							b.Error("bad impl")
						}
					}
				})
			})
		})
	}
}

type mapGenTest struct {
	into interface{}
	from map[string]string
	exp  interface{}
}

var mapGenTests []mapGenTest

func init() {
	var (
		expBoolVal1     = Bool("Yes")
		expBoolVal2     = Bool("FALSE")
		expDurationVal1 = Duration("10ns")
		expDurationVal2 = Duration("20µs")
		expDurationVal3 = Duration("30ms")
		expFloat32Val1  = Float32("1.2")
		expFloat32Val2  = Float32("3.45")
		expFloat32Val3  = Float32("6.78")
		expFloat64Val1  = Float64("1.2")
		expFloat64Val2  = Float64("3.45")
		expFloat64Val3  = Float64("6.78")
		expIntVal1      = Int("12")
		expIntVal2      = Int("34")
		expIntVal3      = Int("56")
		expInt16Val1    = Int16("12")
		expInt16Val2    = Int16("34")
		expInt16Val3    = Int16("56")
		expInt32Val1    = Int32("12")
		expInt32Val2    = Int32("34")
		expInt32Val3    = Int32("56")
		expInt64Val1    = Int64("12")
		expInt64Val2    = Int64("34")
		expInt64Val3    = Int64("56")
		expInt8Val1     = Int8("12")
		expInt8Val2     = Int8("34")
		expInt8Val3     = Int8("56")
		expStringVal1   = String("k1")
		expStringVal2   = String("K2")
		expStringVal3   = String("03")
		expTimeVal1     = Time("2 Jan 2006 15:04:05 -0700 (UTC)")
		expTimeVal2     = Time("Mon, 2 Jan 16:04:05 UTC 2006")
		expTimeVal3     = Time("Mon, 02 Jan 2006 17:04:05 (UTC)")
		expUintVal1     = Uint("12")
		expUintVal2     = Uint("34")
		expUintVal3     = Uint("56")
		expUint16Val1   = Uint16("12")
		expUint16Val2   = Uint16("34")
		expUint16Val3   = Uint16("56")
		expUint32Val1   = Uint32("12")
		expUint32Val2   = Uint32("34")
		expUint32Val3   = Uint32("56")
		expUint64Val1   = Uint64("12")
		expUint64Val2   = Uint64("34")
		expUint64Val3   = Uint64("56")
		expUint8Val1    = Uint8("12")
		expUint8Val2    = Uint8("34")
		expUint8Val3    = Uint8("56")
	)

	mapGenTests = []mapGenTest{
		{make(map[bool]bool),
			map[string]string{"Yes": "Yes", "FALSE": "FALSE"},
			map[bool]bool{expBoolVal1: expBoolVal1, expBoolVal2: expBoolVal2}},
		{make(map[bool]*bool),
			map[string]string{"Yes": "Yes", "FALSE": "FALSE"},
			map[bool]*bool{expBoolVal1: &expBoolVal1, expBoolVal2: &expBoolVal2}},
		{make(map[bool]time.Duration),
			map[string]string{"Yes": "10ns", "FALSE": "20µs"},
			map[bool]time.Duration{
				expBoolVal1: expDurationVal1, expBoolVal2: expDurationVal2}},
		{make(map[bool]*time.Duration),
			map[string]string{"Yes": "10ns", "FALSE": "20µs"},
			map[bool]*time.Duration{
				expBoolVal1: &expDurationVal1, expBoolVal2: &expDurationVal2}},
		{make(map[bool]float32),
			map[string]string{"Yes": "1.2", "FALSE": "3.45"},
			map[bool]float32{expBoolVal1: expFloat32Val1, expBoolVal2: expFloat32Val2}},
		{make(map[bool]*float32),
			map[string]string{"Yes": "1.2", "FALSE": "3.45"},
			map[bool]*float32{expBoolVal1: &expFloat32Val1, expBoolVal2: &expFloat32Val2}},
		{make(map[bool]float64),
			map[string]string{"Yes": "1.2", "FALSE": "3.45"},
			map[bool]float64{expBoolVal1: expFloat64Val1, expBoolVal2: expFloat64Val2}},
		{make(map[bool]*float64),
			map[string]string{"Yes": "1.2", "FALSE": "3.45"},
			map[bool]*float64{expBoolVal1: &expFloat64Val1, expBoolVal2: &expFloat64Val2}},
		{make(map[bool]int),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]int{expBoolVal1: expIntVal1, expBoolVal2: expIntVal2}},
		{make(map[bool]*int),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*int{expBoolVal1: &expIntVal1, expBoolVal2: &expIntVal2}},
		{make(map[bool]int16),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]int16{expBoolVal1: expInt16Val1, expBoolVal2: expInt16Val2}},
		{make(map[bool]*int16),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*int16{expBoolVal1: &expInt16Val1, expBoolVal2: &expInt16Val2}},
		{make(map[bool]int32),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]int32{expBoolVal1: expInt32Val1, expBoolVal2: expInt32Val2}},
		{make(map[bool]*int32),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*int32{expBoolVal1: &expInt32Val1, expBoolVal2: &expInt32Val2}},
		{make(map[bool]int64),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]int64{expBoolVal1: expInt64Val1, expBoolVal2: expInt64Val2}},
		{make(map[bool]*int64),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*int64{expBoolVal1: &expInt64Val1, expBoolVal2: &expInt64Val2}},
		{make(map[bool]int8),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]int8{expBoolVal1: expInt8Val1, expBoolVal2: expInt8Val2}},
		{make(map[bool]*int8),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*int8{expBoolVal1: &expInt8Val1, expBoolVal2: &expInt8Val2}},
		{make(map[bool]string),
			map[string]string{"Yes": "k1", "FALSE": "K2"},
			map[bool]string{expBoolVal1: expStringVal1, expBoolVal2: expStringVal2}},
		{make(map[bool]*string),
			map[string]string{"Yes": "k1", "FALSE": "K2"},
			map[bool]*string{expBoolVal1: &expStringVal1, expBoolVal2: &expStringVal2}},
		{make(map[bool]time.Time),
			map[string]string{
				"Yes":   "2 Jan 2006 15:04:05 -0700 (UTC)",
				"FALSE": "Mon, 2 Jan 16:04:05 UTC 2006"},
			map[bool]time.Time{expBoolVal1: expTimeVal1, expBoolVal2: expTimeVal2}},
		{make(map[bool]*time.Time),
			map[string]string{
				"Yes":   "2 Jan 2006 15:04:05 -0700 (UTC)",
				"FALSE": "Mon, 2 Jan 16:04:05 UTC 2006"},
			map[bool]*time.Time{expBoolVal1: &expTimeVal1, expBoolVal2: &expTimeVal2}},
		{make(map[bool]uint),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]uint{expBoolVal1: expUintVal1, expBoolVal2: expUintVal2}},
		{make(map[bool]*uint),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*uint{expBoolVal1: &expUintVal1, expBoolVal2: &expUintVal2}},
		{make(map[bool]uint16),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]uint16{expBoolVal1: expUint16Val1, expBoolVal2: expUint16Val2}},
		{make(map[bool]*uint16),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*uint16{expBoolVal1: &expUint16Val1, expBoolVal2: &expUint16Val2}},
		{make(map[bool]uint32),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]uint32{expBoolVal1: expUint32Val1, expBoolVal2: expUint32Val2}},
		{make(map[bool]*uint32),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*uint32{expBoolVal1: &expUint32Val1, expBoolVal2: &expUint32Val2}},
		{make(map[bool]uint64),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]uint64{expBoolVal1: expUint64Val1, expBoolVal2: expUint64Val2}},
		{make(map[bool]*uint64),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*uint64{expBoolVal1: &expUint64Val1, expBoolVal2: &expUint64Val2}},
		{make(map[bool]uint8),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]uint8{expBoolVal1: expUint8Val1, expBoolVal2: expUint8Val2}},
		{make(map[bool]*uint8),
			map[string]string{"Yes": "12", "FALSE": "34"},
			map[bool]*uint8{expBoolVal1: &expUint8Val1, expBoolVal2: &expUint8Val2}},
		{make(map[time.Duration]bool),
			map[string]string{"10ns": "Yes", "20µs": "FALSE"},
			map[time.Duration]bool{
				expDurationVal1: expBoolVal1, expDurationVal2: expBoolVal2}},
		{make(map[time.Duration]*bool),
			map[string]string{"10ns": "Yes", "20µs": "FALSE"},
			map[time.Duration]*bool{
				expDurationVal1: &expBoolVal1, expDurationVal2: &expBoolVal2}},
		{make(map[time.Duration]time.Duration),
			map[string]string{"10ns": "10ns", "20µs": "20µs", "30ms": "30ms"},
			map[time.Duration]time.Duration{
				expDurationVal1: expDurationVal1,
				expDurationVal2: expDurationVal2,
				expDurationVal3: expDurationVal3}},
		{make(map[time.Duration]*time.Duration),
			map[string]string{"10ns": "10ns", "20µs": "20µs", "30ms": "30ms"},
			map[time.Duration]*time.Duration{
				expDurationVal1: &expDurationVal1,
				expDurationVal2: &expDurationVal2,
				expDurationVal3: &expDurationVal3}},
		{make(map[time.Duration]float32),
			map[string]string{"10ns": "1.2", "20µs": "3.45", "30ms": "6.78"},
			map[time.Duration]float32{
				expDurationVal1: expFloat32Val1,
				expDurationVal2: expFloat32Val2,
				expDurationVal3: expFloat32Val3}},
		{make(map[time.Duration]*float32),
			map[string]string{"10ns": "1.2", "20µs": "3.45", "30ms": "6.78"},
			map[time.Duration]*float32{
				expDurationVal1: &expFloat32Val1,
				expDurationVal2: &expFloat32Val2,
				expDurationVal3: &expFloat32Val3}},
		{make(map[time.Duration]float64),
			map[string]string{"10ns": "1.2", "20µs": "3.45", "30ms": "6.78"},
			map[time.Duration]float64{
				expDurationVal1: expFloat64Val1,
				expDurationVal2: expFloat64Val2,
				expDurationVal3: expFloat64Val3}},
		{make(map[time.Duration]*float64),
			map[string]string{"10ns": "1.2", "20µs": "3.45", "30ms": "6.78"},
			map[time.Duration]*float64{
				expDurationVal1: &expFloat64Val1,
				expDurationVal2: &expFloat64Val2,
				expDurationVal3: &expFloat64Val3}},
		{make(map[time.Duration]int),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]int{
				expDurationVal1: expIntVal1,
				expDurationVal2: expIntVal2,
				expDurationVal3: expIntVal3}},
		{make(map[time.Duration]*int),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*int{
				expDurationVal1: &expIntVal1,
				expDurationVal2: &expIntVal2,
				expDurationVal3: &expIntVal3}},
		{make(map[time.Duration]int16),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]int16{
				expDurationVal1: expInt16Val1,
				expDurationVal2: expInt16Val2,
				expDurationVal3: expInt16Val3}},
		{make(map[time.Duration]*int16),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*int16{
				expDurationVal1: &expInt16Val1,
				expDurationVal2: &expInt16Val2,
				expDurationVal3: &expInt16Val3}},
		{make(map[time.Duration]int32),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]int32{
				expDurationVal1: expInt32Val1,
				expDurationVal2: expInt32Val2,
				expDurationVal3: expInt32Val3}},
		{make(map[time.Duration]*int32),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*int32{
				expDurationVal1: &expInt32Val1,
				expDurationVal2: &expInt32Val2,
				expDurationVal3: &expInt32Val3}},
		{make(map[time.Duration]int64),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]int64{
				expDurationVal1: expInt64Val1,
				expDurationVal2: expInt64Val2,
				expDurationVal3: expInt64Val3}},
		{make(map[time.Duration]*int64),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*int64{
				expDurationVal1: &expInt64Val1,
				expDurationVal2: &expInt64Val2,
				expDurationVal3: &expInt64Val3}},
		{make(map[time.Duration]int8),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]int8{
				expDurationVal1: expInt8Val1,
				expDurationVal2: expInt8Val2,
				expDurationVal3: expInt8Val3}},
		{make(map[time.Duration]*int8),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*int8{
				expDurationVal1: &expInt8Val1,
				expDurationVal2: &expInt8Val2,
				expDurationVal3: &expInt8Val3}},
		{make(map[time.Duration]string),
			map[string]string{"10ns": "k1", "20µs": "K2", "30ms": "03"},
			map[time.Duration]string{
				expDurationVal1: expStringVal1,
				expDurationVal2: expStringVal2,
				expDurationVal3: expStringVal3}},
		{make(map[time.Duration]*string),
			map[string]string{"10ns": "k1", "20µs": "K2", "30ms": "03"},
			map[time.Duration]*string{
				expDurationVal1: &expStringVal1,
				expDurationVal2: &expStringVal2,
				expDurationVal3: &expStringVal3}},
		{make(map[time.Duration]time.Time),
			map[string]string{
				"10ns": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"20µs": "Mon, 2 Jan 16:04:05 UTC 2006",
				"30ms": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[time.Duration]time.Time{
				expDurationVal1: expTimeVal1,
				expDurationVal2: expTimeVal2,
				expDurationVal3: expTimeVal3}},
		{make(map[time.Duration]*time.Time),
			map[string]string{
				"10ns": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"20µs": "Mon, 2 Jan 16:04:05 UTC 2006",
				"30ms": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[time.Duration]*time.Time{
				expDurationVal1: &expTimeVal1,
				expDurationVal2: &expTimeVal2,
				expDurationVal3: &expTimeVal3}},
		{make(map[time.Duration]uint),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]uint{
				expDurationVal1: expUintVal1,
				expDurationVal2: expUintVal2,
				expDurationVal3: expUintVal3}},
		{make(map[time.Duration]*uint),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*uint{
				expDurationVal1: &expUintVal1,
				expDurationVal2: &expUintVal2,
				expDurationVal3: &expUintVal3}},
		{make(map[time.Duration]uint16),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]uint16{
				expDurationVal1: expUint16Val1,
				expDurationVal2: expUint16Val2,
				expDurationVal3: expUint16Val3}},
		{make(map[time.Duration]*uint16),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*uint16{
				expDurationVal1: &expUint16Val1,
				expDurationVal2: &expUint16Val2,
				expDurationVal3: &expUint16Val3}},
		{make(map[time.Duration]uint32),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]uint32{
				expDurationVal1: expUint32Val1,
				expDurationVal2: expUint32Val2,
				expDurationVal3: expUint32Val3}},
		{make(map[time.Duration]*uint32),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*uint32{
				expDurationVal1: &expUint32Val1,
				expDurationVal2: &expUint32Val2,
				expDurationVal3: &expUint32Val3}},
		{make(map[time.Duration]uint64),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]uint64{
				expDurationVal1: expUint64Val1,
				expDurationVal2: expUint64Val2,
				expDurationVal3: expUint64Val3}},
		{make(map[time.Duration]*uint64),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*uint64{
				expDurationVal1: &expUint64Val1,
				expDurationVal2: &expUint64Val2,
				expDurationVal3: &expUint64Val3}},
		{make(map[time.Duration]uint8),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]uint8{
				expDurationVal1: expUint8Val1,
				expDurationVal2: expUint8Val2,
				expDurationVal3: expUint8Val3}},
		{make(map[time.Duration]*uint8),
			map[string]string{"10ns": "12", "20µs": "34", "30ms": "56"},
			map[time.Duration]*uint8{
				expDurationVal1: &expUint8Val1,
				expDurationVal2: &expUint8Val2,
				expDurationVal3: &expUint8Val3}},
		{make(map[float32]bool),
			map[string]string{"1.2": "Yes", "3.45": "FALSE"},
			map[float32]bool{expFloat32Val1: expBoolVal1, expFloat32Val2: expBoolVal2}},
		{make(map[float32]*bool),
			map[string]string{"1.2": "Yes", "3.45": "FALSE"},
			map[float32]*bool{expFloat32Val1: &expBoolVal1, expFloat32Val2: &expBoolVal2}},
		{make(map[float32]time.Duration),
			map[string]string{"1.2": "10ns", "3.45": "20µs", "6.78": "30ms"},
			map[float32]time.Duration{
				expFloat32Val1: expDurationVal1,
				expFloat32Val2: expDurationVal2,
				expFloat32Val3: expDurationVal3}},
		{make(map[float32]*time.Duration),
			map[string]string{"1.2": "10ns", "3.45": "20µs", "6.78": "30ms"},
			map[float32]*time.Duration{
				expFloat32Val1: &expDurationVal1,
				expFloat32Val2: &expDurationVal2,
				expFloat32Val3: &expDurationVal3}},
		{make(map[float32]float32),
			map[string]string{"1.2": "1.2", "3.45": "3.45", "6.78": "6.78"},
			map[float32]float32{
				expFloat32Val1: expFloat32Val1,
				expFloat32Val2: expFloat32Val2,
				expFloat32Val3: expFloat32Val3}},
		{make(map[float32]*float32),
			map[string]string{"1.2": "1.2", "3.45": "3.45", "6.78": "6.78"},
			map[float32]*float32{
				expFloat32Val1: &expFloat32Val1,
				expFloat32Val2: &expFloat32Val2,
				expFloat32Val3: &expFloat32Val3}},
		{make(map[float32]float64),
			map[string]string{"1.2": "1.2", "3.45": "3.45", "6.78": "6.78"},
			map[float32]float64{
				expFloat32Val1: expFloat64Val1,
				expFloat32Val2: expFloat64Val2,
				expFloat32Val3: expFloat64Val3}},
		{make(map[float32]*float64),
			map[string]string{"1.2": "1.2", "3.45": "3.45", "6.78": "6.78"},
			map[float32]*float64{
				expFloat32Val1: &expFloat64Val1,
				expFloat32Val2: &expFloat64Val2,
				expFloat32Val3: &expFloat64Val3}},
		{make(map[float32]int),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]int{
				expFloat32Val1: expIntVal1,
				expFloat32Val2: expIntVal2,
				expFloat32Val3: expIntVal3}},
		{make(map[float32]*int),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*int{
				expFloat32Val1: &expIntVal1,
				expFloat32Val2: &expIntVal2,
				expFloat32Val3: &expIntVal3}},
		{make(map[float32]int16),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]int16{
				expFloat32Val1: expInt16Val1,
				expFloat32Val2: expInt16Val2,
				expFloat32Val3: expInt16Val3}},
		{make(map[float32]*int16),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*int16{
				expFloat32Val1: &expInt16Val1,
				expFloat32Val2: &expInt16Val2,
				expFloat32Val3: &expInt16Val3}},
		{make(map[float32]int32),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]int32{
				expFloat32Val1: expInt32Val1,
				expFloat32Val2: expInt32Val2,
				expFloat32Val3: expInt32Val3}},
		{make(map[float32]*int32),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*int32{
				expFloat32Val1: &expInt32Val1,
				expFloat32Val2: &expInt32Val2,
				expFloat32Val3: &expInt32Val3}},
		{make(map[float32]int64),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]int64{
				expFloat32Val1: expInt64Val1,
				expFloat32Val2: expInt64Val2,
				expFloat32Val3: expInt64Val3}},
		{make(map[float32]*int64),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*int64{
				expFloat32Val1: &expInt64Val1,
				expFloat32Val2: &expInt64Val2,
				expFloat32Val3: &expInt64Val3}},
		{make(map[float32]int8),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]int8{
				expFloat32Val1: expInt8Val1,
				expFloat32Val2: expInt8Val2,
				expFloat32Val3: expInt8Val3}},
		{make(map[float32]*int8),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*int8{
				expFloat32Val1: &expInt8Val1,
				expFloat32Val2: &expInt8Val2,
				expFloat32Val3: &expInt8Val3}},
		{make(map[float32]string),
			map[string]string{"1.2": "k1", "3.45": "K2", "6.78": "03"},
			map[float32]string{
				expFloat32Val1: expStringVal1,
				expFloat32Val2: expStringVal2,
				expFloat32Val3: expStringVal3}},
		{make(map[float32]*string),
			map[string]string{"1.2": "k1", "3.45": "K2", "6.78": "03"},
			map[float32]*string{
				expFloat32Val1: &expStringVal1,
				expFloat32Val2: &expStringVal2,
				expFloat32Val3: &expStringVal3}},
		{make(map[float32]time.Time),
			map[string]string{
				"1.2":  "2 Jan 2006 15:04:05 -0700 (UTC)",
				"3.45": "Mon, 2 Jan 16:04:05 UTC 2006",
				"6.78": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[float32]time.Time{
				expFloat32Val1: expTimeVal1,
				expFloat32Val2: expTimeVal2,
				expFloat32Val3: expTimeVal3}},
		{make(map[float32]*time.Time),
			map[string]string{
				"1.2":  "2 Jan 2006 15:04:05 -0700 (UTC)",
				"3.45": "Mon, 2 Jan 16:04:05 UTC 2006",
				"6.78": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[float32]*time.Time{
				expFloat32Val1: &expTimeVal1,
				expFloat32Val2: &expTimeVal2,
				expFloat32Val3: &expTimeVal3}},
		{make(map[float32]uint),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]uint{
				expFloat32Val1: expUintVal1,
				expFloat32Val2: expUintVal2,
				expFloat32Val3: expUintVal3}},
		{make(map[float32]*uint),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*uint{
				expFloat32Val1: &expUintVal1,
				expFloat32Val2: &expUintVal2,
				expFloat32Val3: &expUintVal3}},
		{make(map[float32]uint16),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]uint16{
				expFloat32Val1: expUint16Val1,
				expFloat32Val2: expUint16Val2,
				expFloat32Val3: expUint16Val3}},
		{make(map[float32]*uint16),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*uint16{
				expFloat32Val1: &expUint16Val1,
				expFloat32Val2: &expUint16Val2,
				expFloat32Val3: &expUint16Val3}},
		{make(map[float32]uint32),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]uint32{
				expFloat32Val1: expUint32Val1,
				expFloat32Val2: expUint32Val2,
				expFloat32Val3: expUint32Val3}},
		{make(map[float32]*uint32),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*uint32{
				expFloat32Val1: &expUint32Val1,
				expFloat32Val2: &expUint32Val2,
				expFloat32Val3: &expUint32Val3}},
		{make(map[float32]uint64),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]uint64{
				expFloat32Val1: expUint64Val1,
				expFloat32Val2: expUint64Val2,
				expFloat32Val3: expUint64Val3}},
		{make(map[float32]*uint64),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*uint64{
				expFloat32Val1: &expUint64Val1,
				expFloat32Val2: &expUint64Val2,
				expFloat32Val3: &expUint64Val3}},
		{make(map[float32]uint8),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]uint8{
				expFloat32Val1: expUint8Val1,
				expFloat32Val2: expUint8Val2,
				expFloat32Val3: expUint8Val3}},
		{make(map[float32]*uint8),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float32]*uint8{
				expFloat32Val1: &expUint8Val1,
				expFloat32Val2: &expUint8Val2,
				expFloat32Val3: &expUint8Val3}},
		{make(map[float64]bool),
			map[string]string{"1.2": "Yes", "3.45": "FALSE"},
			map[float64]bool{expFloat64Val1: expBoolVal1, expFloat64Val2: expBoolVal2}},
		{make(map[float64]*bool),
			map[string]string{"1.2": "Yes", "3.45": "FALSE"},
			map[float64]*bool{expFloat64Val1: &expBoolVal1, expFloat64Val2: &expBoolVal2}},
		{make(map[float64]time.Duration),
			map[string]string{"1.2": "10ns", "3.45": "20µs", "6.78": "30ms"},
			map[float64]time.Duration{
				expFloat64Val1: expDurationVal1,
				expFloat64Val2: expDurationVal2,
				expFloat64Val3: expDurationVal3}},
		{make(map[float64]*time.Duration),
			map[string]string{"1.2": "10ns", "3.45": "20µs", "6.78": "30ms"},
			map[float64]*time.Duration{
				expFloat64Val1: &expDurationVal1,
				expFloat64Val2: &expDurationVal2,
				expFloat64Val3: &expDurationVal3}},
		{make(map[float64]float32),
			map[string]string{"1.2": "1.2", "3.45": "3.45", "6.78": "6.78"},
			map[float64]float32{
				expFloat64Val1: expFloat32Val1,
				expFloat64Val2: expFloat32Val2,
				expFloat64Val3: expFloat32Val3}},
		{make(map[float64]*float32),
			map[string]string{"1.2": "1.2", "3.45": "3.45", "6.78": "6.78"},
			map[float64]*float32{
				expFloat64Val1: &expFloat32Val1,
				expFloat64Val2: &expFloat32Val2,
				expFloat64Val3: &expFloat32Val3}},
		{make(map[float64]float64),
			map[string]string{"1.2": "1.2", "3.45": "3.45", "6.78": "6.78"},
			map[float64]float64{
				expFloat64Val1: expFloat64Val1,
				expFloat64Val2: expFloat64Val2,
				expFloat64Val3: expFloat64Val3}},
		{make(map[float64]*float64),
			map[string]string{"1.2": "1.2", "3.45": "3.45", "6.78": "6.78"},
			map[float64]*float64{
				expFloat64Val1: &expFloat64Val1,
				expFloat64Val2: &expFloat64Val2,
				expFloat64Val3: &expFloat64Val3}},
		{make(map[float64]int),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]int{
				expFloat64Val1: expIntVal1,
				expFloat64Val2: expIntVal2,
				expFloat64Val3: expIntVal3}},
		{make(map[float64]*int),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*int{
				expFloat64Val1: &expIntVal1,
				expFloat64Val2: &expIntVal2,
				expFloat64Val3: &expIntVal3}},
		{make(map[float64]int16),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]int16{
				expFloat64Val1: expInt16Val1,
				expFloat64Val2: expInt16Val2,
				expFloat64Val3: expInt16Val3}},
		{make(map[float64]*int16),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*int16{
				expFloat64Val1: &expInt16Val1,
				expFloat64Val2: &expInt16Val2,
				expFloat64Val3: &expInt16Val3}},
		{make(map[float64]int32),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]int32{
				expFloat64Val1: expInt32Val1,
				expFloat64Val2: expInt32Val2,
				expFloat64Val3: expInt32Val3}},
		{make(map[float64]*int32),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*int32{
				expFloat64Val1: &expInt32Val1,
				expFloat64Val2: &expInt32Val2,
				expFloat64Val3: &expInt32Val3}},
		{make(map[float64]int64),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]int64{
				expFloat64Val1: expInt64Val1,
				expFloat64Val2: expInt64Val2,
				expFloat64Val3: expInt64Val3}},
		{make(map[float64]*int64),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*int64{
				expFloat64Val1: &expInt64Val1,
				expFloat64Val2: &expInt64Val2,
				expFloat64Val3: &expInt64Val3}},
		{make(map[float64]int8),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]int8{
				expFloat64Val1: expInt8Val1,
				expFloat64Val2: expInt8Val2,
				expFloat64Val3: expInt8Val3}},
		{make(map[float64]*int8),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*int8{
				expFloat64Val1: &expInt8Val1,
				expFloat64Val2: &expInt8Val2,
				expFloat64Val3: &expInt8Val3}},
		{make(map[float64]string),
			map[string]string{"1.2": "k1", "3.45": "K2", "6.78": "03"},
			map[float64]string{
				expFloat64Val1: expStringVal1,
				expFloat64Val2: expStringVal2,
				expFloat64Val3: expStringVal3}},
		{make(map[float64]*string),
			map[string]string{"1.2": "k1", "3.45": "K2", "6.78": "03"},
			map[float64]*string{
				expFloat64Val1: &expStringVal1,
				expFloat64Val2: &expStringVal2,
				expFloat64Val3: &expStringVal3}},
		{make(map[float64]time.Time),
			map[string]string{
				"1.2":  "2 Jan 2006 15:04:05 -0700 (UTC)",
				"3.45": "Mon, 2 Jan 16:04:05 UTC 2006",
				"6.78": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[float64]time.Time{
				expFloat64Val1: expTimeVal1,
				expFloat64Val2: expTimeVal2,
				expFloat64Val3: expTimeVal3}},
		{make(map[float64]*time.Time),
			map[string]string{
				"1.2":  "2 Jan 2006 15:04:05 -0700 (UTC)",
				"3.45": "Mon, 2 Jan 16:04:05 UTC 2006",
				"6.78": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[float64]*time.Time{
				expFloat64Val1: &expTimeVal1,
				expFloat64Val2: &expTimeVal2,
				expFloat64Val3: &expTimeVal3}},
		{make(map[float64]uint),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]uint{
				expFloat64Val1: expUintVal1,
				expFloat64Val2: expUintVal2,
				expFloat64Val3: expUintVal3}},
		{make(map[float64]*uint),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*uint{
				expFloat64Val1: &expUintVal1,
				expFloat64Val2: &expUintVal2,
				expFloat64Val3: &expUintVal3}},
		{make(map[float64]uint16),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]uint16{
				expFloat64Val1: expUint16Val1,
				expFloat64Val2: expUint16Val2,
				expFloat64Val3: expUint16Val3}},
		{make(map[float64]*uint16),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*uint16{
				expFloat64Val1: &expUint16Val1,
				expFloat64Val2: &expUint16Val2,
				expFloat64Val3: &expUint16Val3}},
		{make(map[float64]uint32),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]uint32{
				expFloat64Val1: expUint32Val1,
				expFloat64Val2: expUint32Val2,
				expFloat64Val3: expUint32Val3}},
		{make(map[float64]*uint32),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*uint32{
				expFloat64Val1: &expUint32Val1,
				expFloat64Val2: &expUint32Val2,
				expFloat64Val3: &expUint32Val3}},
		{make(map[float64]uint64),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]uint64{
				expFloat64Val1: expUint64Val1,
				expFloat64Val2: expUint64Val2,
				expFloat64Val3: expUint64Val3}},
		{make(map[float64]*uint64),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*uint64{
				expFloat64Val1: &expUint64Val1,
				expFloat64Val2: &expUint64Val2,
				expFloat64Val3: &expUint64Val3}},
		{make(map[float64]uint8),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]uint8{
				expFloat64Val1: expUint8Val1,
				expFloat64Val2: expUint8Val2,
				expFloat64Val3: expUint8Val3}},
		{make(map[float64]*uint8),
			map[string]string{"1.2": "12", "3.45": "34", "6.78": "56"},
			map[float64]*uint8{
				expFloat64Val1: &expUint8Val1,
				expFloat64Val2: &expUint8Val2,
				expFloat64Val3: &expUint8Val3}},
		{make(map[int]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int]bool{expIntVal1: expBoolVal1, expIntVal2: expBoolVal2}},
		{make(map[int]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int]*bool{expIntVal1: &expBoolVal1, expIntVal2: &expBoolVal2}},
		{make(map[int]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int]time.Duration{
				expIntVal1: expDurationVal1,
				expIntVal2: expDurationVal2,
				expIntVal3: expDurationVal3}},
		{make(map[int]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int]*time.Duration{
				expIntVal1: &expDurationVal1,
				expIntVal2: &expDurationVal2,
				expIntVal3: &expDurationVal3}},
		{make(map[int]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int]float32{
				expIntVal1: expFloat32Val1,
				expIntVal2: expFloat32Val2,
				expIntVal3: expFloat32Val3}},
		{make(map[int]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int]*float32{
				expIntVal1: &expFloat32Val1,
				expIntVal2: &expFloat32Val2,
				expIntVal3: &expFloat32Val3}},
		{make(map[int]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int]float64{
				expIntVal1: expFloat64Val1,
				expIntVal2: expFloat64Val2,
				expIntVal3: expFloat64Val3}},
		{make(map[int]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int]*float64{
				expIntVal1: &expFloat64Val1,
				expIntVal2: &expFloat64Val2,
				expIntVal3: &expFloat64Val3}},
		{make(map[int]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]int{expIntVal1: expIntVal1, expIntVal2: expIntVal2, expIntVal3: expIntVal3}},
		{make(map[int]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*int{
				expIntVal1: &expIntVal1, expIntVal2: &expIntVal2, expIntVal3: &expIntVal3}},
		{make(map[int]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]int16{
				expIntVal1: expInt16Val1, expIntVal2: expInt16Val2, expIntVal3: expInt16Val3}},
		{make(map[int]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*int16{
				expIntVal1: &expInt16Val1,
				expIntVal2: &expInt16Val2,
				expIntVal3: &expInt16Val3}},
		{make(map[int]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]int32{
				expIntVal1: expInt32Val1, expIntVal2: expInt32Val2, expIntVal3: expInt32Val3}},
		{make(map[int]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*int32{
				expIntVal1: &expInt32Val1,
				expIntVal2: &expInt32Val2,
				expIntVal3: &expInt32Val3}},
		{make(map[int]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]int64{
				expIntVal1: expInt64Val1, expIntVal2: expInt64Val2, expIntVal3: expInt64Val3}},
		{make(map[int]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*int64{
				expIntVal1: &expInt64Val1,
				expIntVal2: &expInt64Val2,
				expIntVal3: &expInt64Val3}},
		{make(map[int]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]int8{
				expIntVal1: expInt8Val1, expIntVal2: expInt8Val2, expIntVal3: expInt8Val3}},
		{make(map[int]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*int8{
				expIntVal1: &expInt8Val1, expIntVal2: &expInt8Val2, expIntVal3: &expInt8Val3}},
		{make(map[int]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int]string{
				expIntVal1: expStringVal1,
				expIntVal2: expStringVal2,
				expIntVal3: expStringVal3}},
		{make(map[int]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int]*string{
				expIntVal1: &expStringVal1,
				expIntVal2: &expStringVal2,
				expIntVal3: &expStringVal3}},
		{make(map[int]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int]time.Time{
				expIntVal1: expTimeVal1, expIntVal2: expTimeVal2, expIntVal3: expTimeVal3}},
		{make(map[int]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int]*time.Time{
				expIntVal1: &expTimeVal1, expIntVal2: &expTimeVal2, expIntVal3: &expTimeVal3}},
		{make(map[int]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]uint{
				expIntVal1: expUintVal1, expIntVal2: expUintVal2, expIntVal3: expUintVal3}},
		{make(map[int]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*uint{
				expIntVal1: &expUintVal1, expIntVal2: &expUintVal2, expIntVal3: &expUintVal3}},
		{make(map[int]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]uint16{
				expIntVal1: expUint16Val1,
				expIntVal2: expUint16Val2,
				expIntVal3: expUint16Val3}},
		{make(map[int]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*uint16{
				expIntVal1: &expUint16Val1,
				expIntVal2: &expUint16Val2,
				expIntVal3: &expUint16Val3}},
		{make(map[int]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]uint32{
				expIntVal1: expUint32Val1,
				expIntVal2: expUint32Val2,
				expIntVal3: expUint32Val3}},
		{make(map[int]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*uint32{
				expIntVal1: &expUint32Val1,
				expIntVal2: &expUint32Val2,
				expIntVal3: &expUint32Val3}},
		{make(map[int]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]uint64{
				expIntVal1: expUint64Val1,
				expIntVal2: expUint64Val2,
				expIntVal3: expUint64Val3}},
		{make(map[int]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*uint64{
				expIntVal1: &expUint64Val1,
				expIntVal2: &expUint64Val2,
				expIntVal3: &expUint64Val3}},
		{make(map[int]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]uint8{
				expIntVal1: expUint8Val1, expIntVal2: expUint8Val2, expIntVal3: expUint8Val3}},
		{make(map[int]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int]*uint8{
				expIntVal1: &expUint8Val1,
				expIntVal2: &expUint8Val2,
				expIntVal3: &expUint8Val3}},
		{make(map[int16]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int16]bool{expInt16Val1: expBoolVal1, expInt16Val2: expBoolVal2}},
		{make(map[int16]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int16]*bool{expInt16Val1: &expBoolVal1, expInt16Val2: &expBoolVal2}},
		{make(map[int16]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int16]time.Duration{
				expInt16Val1: expDurationVal1,
				expInt16Val2: expDurationVal2,
				expInt16Val3: expDurationVal3}},
		{make(map[int16]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int16]*time.Duration{
				expInt16Val1: &expDurationVal1,
				expInt16Val2: &expDurationVal2,
				expInt16Val3: &expDurationVal3}},
		{make(map[int16]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int16]float32{
				expInt16Val1: expFloat32Val1,
				expInt16Val2: expFloat32Val2,
				expInt16Val3: expFloat32Val3}},
		{make(map[int16]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int16]*float32{
				expInt16Val1: &expFloat32Val1,
				expInt16Val2: &expFloat32Val2,
				expInt16Val3: &expFloat32Val3}},
		{make(map[int16]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int16]float64{
				expInt16Val1: expFloat64Val1,
				expInt16Val2: expFloat64Val2,
				expInt16Val3: expFloat64Val3}},
		{make(map[int16]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int16]*float64{
				expInt16Val1: &expFloat64Val1,
				expInt16Val2: &expFloat64Val2,
				expInt16Val3: &expFloat64Val3}},
		{make(map[int16]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]int{
				expInt16Val1: expIntVal1, expInt16Val2: expIntVal2, expInt16Val3: expIntVal3}},
		{make(map[int16]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*int{
				expInt16Val1: &expIntVal1,
				expInt16Val2: &expIntVal2,
				expInt16Val3: &expIntVal3}},
		{make(map[int16]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]int16{
				expInt16Val1: expInt16Val1,
				expInt16Val2: expInt16Val2,
				expInt16Val3: expInt16Val3}},
		{make(map[int16]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*int16{
				expInt16Val1: &expInt16Val1,
				expInt16Val2: &expInt16Val2,
				expInt16Val3: &expInt16Val3}},
		{make(map[int16]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]int32{
				expInt16Val1: expInt32Val1,
				expInt16Val2: expInt32Val2,
				expInt16Val3: expInt32Val3}},
		{make(map[int16]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*int32{
				expInt16Val1: &expInt32Val1,
				expInt16Val2: &expInt32Val2,
				expInt16Val3: &expInt32Val3}},
		{make(map[int16]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]int64{
				expInt16Val1: expInt64Val1,
				expInt16Val2: expInt64Val2,
				expInt16Val3: expInt64Val3}},
		{make(map[int16]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*int64{
				expInt16Val1: &expInt64Val1,
				expInt16Val2: &expInt64Val2,
				expInt16Val3: &expInt64Val3}},
		{make(map[int16]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]int8{
				expInt16Val1: expInt8Val1,
				expInt16Val2: expInt8Val2,
				expInt16Val3: expInt8Val3}},
		{make(map[int16]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*int8{
				expInt16Val1: &expInt8Val1,
				expInt16Val2: &expInt8Val2,
				expInt16Val3: &expInt8Val3}},
		{make(map[int16]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int16]string{
				expInt16Val1: expStringVal1,
				expInt16Val2: expStringVal2,
				expInt16Val3: expStringVal3}},
		{make(map[int16]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int16]*string{
				expInt16Val1: &expStringVal1,
				expInt16Val2: &expStringVal2,
				expInt16Val3: &expStringVal3}},
		{make(map[int16]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int16]time.Time{
				expInt16Val1: expTimeVal1,
				expInt16Val2: expTimeVal2,
				expInt16Val3: expTimeVal3}},
		{make(map[int16]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int16]*time.Time{
				expInt16Val1: &expTimeVal1,
				expInt16Val2: &expTimeVal2,
				expInt16Val3: &expTimeVal3}},
		{make(map[int16]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]uint{
				expInt16Val1: expUintVal1,
				expInt16Val2: expUintVal2,
				expInt16Val3: expUintVal3}},
		{make(map[int16]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*uint{
				expInt16Val1: &expUintVal1,
				expInt16Val2: &expUintVal2,
				expInt16Val3: &expUintVal3}},
		{make(map[int16]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]uint16{
				expInt16Val1: expUint16Val1,
				expInt16Val2: expUint16Val2,
				expInt16Val3: expUint16Val3}},
		{make(map[int16]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*uint16{
				expInt16Val1: &expUint16Val1,
				expInt16Val2: &expUint16Val2,
				expInt16Val3: &expUint16Val3}},
		{make(map[int16]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]uint32{
				expInt16Val1: expUint32Val1,
				expInt16Val2: expUint32Val2,
				expInt16Val3: expUint32Val3}},
		{make(map[int16]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*uint32{
				expInt16Val1: &expUint32Val1,
				expInt16Val2: &expUint32Val2,
				expInt16Val3: &expUint32Val3}},
		{make(map[int16]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]uint64{
				expInt16Val1: expUint64Val1,
				expInt16Val2: expUint64Val2,
				expInt16Val3: expUint64Val3}},
		{make(map[int16]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*uint64{
				expInt16Val1: &expUint64Val1,
				expInt16Val2: &expUint64Val2,
				expInt16Val3: &expUint64Val3}},
		{make(map[int16]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]uint8{
				expInt16Val1: expUint8Val1,
				expInt16Val2: expUint8Val2,
				expInt16Val3: expUint8Val3}},
		{make(map[int16]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int16]*uint8{
				expInt16Val1: &expUint8Val1,
				expInt16Val2: &expUint8Val2,
				expInt16Val3: &expUint8Val3}},
		{make(map[int32]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int32]bool{expInt32Val1: expBoolVal1, expInt32Val2: expBoolVal2}},
		{make(map[int32]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int32]*bool{expInt32Val1: &expBoolVal1, expInt32Val2: &expBoolVal2}},
		{make(map[int32]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int32]time.Duration{
				expInt32Val1: expDurationVal1,
				expInt32Val2: expDurationVal2,
				expInt32Val3: expDurationVal3}},
		{make(map[int32]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int32]*time.Duration{
				expInt32Val1: &expDurationVal1,
				expInt32Val2: &expDurationVal2,
				expInt32Val3: &expDurationVal3}},
		{make(map[int32]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int32]float32{
				expInt32Val1: expFloat32Val1,
				expInt32Val2: expFloat32Val2,
				expInt32Val3: expFloat32Val3}},
		{make(map[int32]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int32]*float32{
				expInt32Val1: &expFloat32Val1,
				expInt32Val2: &expFloat32Val2,
				expInt32Val3: &expFloat32Val3}},
		{make(map[int32]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int32]float64{
				expInt32Val1: expFloat64Val1,
				expInt32Val2: expFloat64Val2,
				expInt32Val3: expFloat64Val3}},
		{make(map[int32]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int32]*float64{
				expInt32Val1: &expFloat64Val1,
				expInt32Val2: &expFloat64Val2,
				expInt32Val3: &expFloat64Val3}},
		{make(map[int32]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]int{
				expInt32Val1: expIntVal1, expInt32Val2: expIntVal2, expInt32Val3: expIntVal3}},
		{make(map[int32]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*int{
				expInt32Val1: &expIntVal1,
				expInt32Val2: &expIntVal2,
				expInt32Val3: &expIntVal3}},
		{make(map[int32]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]int16{
				expInt32Val1: expInt16Val1,
				expInt32Val2: expInt16Val2,
				expInt32Val3: expInt16Val3}},
		{make(map[int32]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*int16{
				expInt32Val1: &expInt16Val1,
				expInt32Val2: &expInt16Val2,
				expInt32Val3: &expInt16Val3}},
		{make(map[int32]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]int32{
				expInt32Val1: expInt32Val1,
				expInt32Val2: expInt32Val2,
				expInt32Val3: expInt32Val3}},
		{make(map[int32]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*int32{
				expInt32Val1: &expInt32Val1,
				expInt32Val2: &expInt32Val2,
				expInt32Val3: &expInt32Val3}},
		{make(map[int32]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]int64{
				expInt32Val1: expInt64Val1,
				expInt32Val2: expInt64Val2,
				expInt32Val3: expInt64Val3}},
		{make(map[int32]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*int64{
				expInt32Val1: &expInt64Val1,
				expInt32Val2: &expInt64Val2,
				expInt32Val3: &expInt64Val3}},
		{make(map[int32]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]int8{
				expInt32Val1: expInt8Val1,
				expInt32Val2: expInt8Val2,
				expInt32Val3: expInt8Val3}},
		{make(map[int32]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*int8{
				expInt32Val1: &expInt8Val1,
				expInt32Val2: &expInt8Val2,
				expInt32Val3: &expInt8Val3}},
		{make(map[int32]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int32]string{
				expInt32Val1: expStringVal1,
				expInt32Val2: expStringVal2,
				expInt32Val3: expStringVal3}},
		{make(map[int32]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int32]*string{
				expInt32Val1: &expStringVal1,
				expInt32Val2: &expStringVal2,
				expInt32Val3: &expStringVal3}},
		{make(map[int32]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int32]time.Time{
				expInt32Val1: expTimeVal1,
				expInt32Val2: expTimeVal2,
				expInt32Val3: expTimeVal3}},
		{make(map[int32]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int32]*time.Time{
				expInt32Val1: &expTimeVal1,
				expInt32Val2: &expTimeVal2,
				expInt32Val3: &expTimeVal3}},
		{make(map[int32]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]uint{
				expInt32Val1: expUintVal1,
				expInt32Val2: expUintVal2,
				expInt32Val3: expUintVal3}},
		{make(map[int32]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*uint{
				expInt32Val1: &expUintVal1,
				expInt32Val2: &expUintVal2,
				expInt32Val3: &expUintVal3}},
		{make(map[int32]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]uint16{
				expInt32Val1: expUint16Val1,
				expInt32Val2: expUint16Val2,
				expInt32Val3: expUint16Val3}},
		{make(map[int32]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*uint16{
				expInt32Val1: &expUint16Val1,
				expInt32Val2: &expUint16Val2,
				expInt32Val3: &expUint16Val3}},
		{make(map[int32]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]uint32{
				expInt32Val1: expUint32Val1,
				expInt32Val2: expUint32Val2,
				expInt32Val3: expUint32Val3}},
		{make(map[int32]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*uint32{
				expInt32Val1: &expUint32Val1,
				expInt32Val2: &expUint32Val2,
				expInt32Val3: &expUint32Val3}},
		{make(map[int32]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]uint64{
				expInt32Val1: expUint64Val1,
				expInt32Val2: expUint64Val2,
				expInt32Val3: expUint64Val3}},
		{make(map[int32]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*uint64{
				expInt32Val1: &expUint64Val1,
				expInt32Val2: &expUint64Val2,
				expInt32Val3: &expUint64Val3}},
		{make(map[int32]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]uint8{
				expInt32Val1: expUint8Val1,
				expInt32Val2: expUint8Val2,
				expInt32Val3: expUint8Val3}},
		{make(map[int32]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int32]*uint8{
				expInt32Val1: &expUint8Val1,
				expInt32Val2: &expUint8Val2,
				expInt32Val3: &expUint8Val3}},
		{make(map[int64]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int64]bool{expInt64Val1: expBoolVal1, expInt64Val2: expBoolVal2}},
		{make(map[int64]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int64]*bool{expInt64Val1: &expBoolVal1, expInt64Val2: &expBoolVal2}},
		{make(map[int64]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int64]time.Duration{
				expInt64Val1: expDurationVal1,
				expInt64Val2: expDurationVal2,
				expInt64Val3: expDurationVal3}},
		{make(map[int64]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int64]*time.Duration{
				expInt64Val1: &expDurationVal1,
				expInt64Val2: &expDurationVal2,
				expInt64Val3: &expDurationVal3}},
		{make(map[int64]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int64]float32{
				expInt64Val1: expFloat32Val1,
				expInt64Val2: expFloat32Val2,
				expInt64Val3: expFloat32Val3}},
		{make(map[int64]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int64]*float32{
				expInt64Val1: &expFloat32Val1,
				expInt64Val2: &expFloat32Val2,
				expInt64Val3: &expFloat32Val3}},
		{make(map[int64]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int64]float64{
				expInt64Val1: expFloat64Val1,
				expInt64Val2: expFloat64Val2,
				expInt64Val3: expFloat64Val3}},
		{make(map[int64]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int64]*float64{
				expInt64Val1: &expFloat64Val1,
				expInt64Val2: &expFloat64Val2,
				expInt64Val3: &expFloat64Val3}},
		{make(map[int64]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]int{
				expInt64Val1: expIntVal1, expInt64Val2: expIntVal2, expInt64Val3: expIntVal3}},
		{make(map[int64]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*int{
				expInt64Val1: &expIntVal1,
				expInt64Val2: &expIntVal2,
				expInt64Val3: &expIntVal3}},
		{make(map[int64]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]int16{
				expInt64Val1: expInt16Val1,
				expInt64Val2: expInt16Val2,
				expInt64Val3: expInt16Val3}},
		{make(map[int64]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*int16{
				expInt64Val1: &expInt16Val1,
				expInt64Val2: &expInt16Val2,
				expInt64Val3: &expInt16Val3}},
		{make(map[int64]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]int32{
				expInt64Val1: expInt32Val1,
				expInt64Val2: expInt32Val2,
				expInt64Val3: expInt32Val3}},
		{make(map[int64]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*int32{
				expInt64Val1: &expInt32Val1,
				expInt64Val2: &expInt32Val2,
				expInt64Val3: &expInt32Val3}},
		{make(map[int64]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]int64{
				expInt64Val1: expInt64Val1,
				expInt64Val2: expInt64Val2,
				expInt64Val3: expInt64Val3}},
		{make(map[int64]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*int64{
				expInt64Val1: &expInt64Val1,
				expInt64Val2: &expInt64Val2,
				expInt64Val3: &expInt64Val3}},
		{make(map[int64]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]int8{
				expInt64Val1: expInt8Val1,
				expInt64Val2: expInt8Val2,
				expInt64Val3: expInt8Val3}},
		{make(map[int64]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*int8{
				expInt64Val1: &expInt8Val1,
				expInt64Val2: &expInt8Val2,
				expInt64Val3: &expInt8Val3}},
		{make(map[int64]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int64]string{
				expInt64Val1: expStringVal1,
				expInt64Val2: expStringVal2,
				expInt64Val3: expStringVal3}},
		{make(map[int64]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int64]*string{
				expInt64Val1: &expStringVal1,
				expInt64Val2: &expStringVal2,
				expInt64Val3: &expStringVal3}},
		{make(map[int64]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int64]time.Time{
				expInt64Val1: expTimeVal1,
				expInt64Val2: expTimeVal2,
				expInt64Val3: expTimeVal3}},
		{make(map[int64]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int64]*time.Time{
				expInt64Val1: &expTimeVal1,
				expInt64Val2: &expTimeVal2,
				expInt64Val3: &expTimeVal3}},
		{make(map[int64]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]uint{
				expInt64Val1: expUintVal1,
				expInt64Val2: expUintVal2,
				expInt64Val3: expUintVal3}},
		{make(map[int64]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*uint{
				expInt64Val1: &expUintVal1,
				expInt64Val2: &expUintVal2,
				expInt64Val3: &expUintVal3}},
		{make(map[int64]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]uint16{
				expInt64Val1: expUint16Val1,
				expInt64Val2: expUint16Val2,
				expInt64Val3: expUint16Val3}},
		{make(map[int64]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*uint16{
				expInt64Val1: &expUint16Val1,
				expInt64Val2: &expUint16Val2,
				expInt64Val3: &expUint16Val3}},
		{make(map[int64]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]uint32{
				expInt64Val1: expUint32Val1,
				expInt64Val2: expUint32Val2,
				expInt64Val3: expUint32Val3}},
		{make(map[int64]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*uint32{
				expInt64Val1: &expUint32Val1,
				expInt64Val2: &expUint32Val2,
				expInt64Val3: &expUint32Val3}},
		{make(map[int64]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]uint64{
				expInt64Val1: expUint64Val1,
				expInt64Val2: expUint64Val2,
				expInt64Val3: expUint64Val3}},
		{make(map[int64]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*uint64{
				expInt64Val1: &expUint64Val1,
				expInt64Val2: &expUint64Val2,
				expInt64Val3: &expUint64Val3}},
		{make(map[int64]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]uint8{
				expInt64Val1: expUint8Val1,
				expInt64Val2: expUint8Val2,
				expInt64Val3: expUint8Val3}},
		{make(map[int64]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int64]*uint8{
				expInt64Val1: &expUint8Val1,
				expInt64Val2: &expUint8Val2,
				expInt64Val3: &expUint8Val3}},
		{make(map[int8]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int8]bool{expInt8Val1: expBoolVal1, expInt8Val2: expBoolVal2}},
		{make(map[int8]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[int8]*bool{expInt8Val1: &expBoolVal1, expInt8Val2: &expBoolVal2}},
		{make(map[int8]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int8]time.Duration{
				expInt8Val1: expDurationVal1,
				expInt8Val2: expDurationVal2,
				expInt8Val3: expDurationVal3}},
		{make(map[int8]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[int8]*time.Duration{
				expInt8Val1: &expDurationVal1,
				expInt8Val2: &expDurationVal2,
				expInt8Val3: &expDurationVal3}},
		{make(map[int8]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int8]float32{
				expInt8Val1: expFloat32Val1,
				expInt8Val2: expFloat32Val2,
				expInt8Val3: expFloat32Val3}},
		{make(map[int8]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int8]*float32{
				expInt8Val1: &expFloat32Val1,
				expInt8Val2: &expFloat32Val2,
				expInt8Val3: &expFloat32Val3}},
		{make(map[int8]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int8]float64{
				expInt8Val1: expFloat64Val1,
				expInt8Val2: expFloat64Val2,
				expInt8Val3: expFloat64Val3}},
		{make(map[int8]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[int8]*float64{
				expInt8Val1: &expFloat64Val1,
				expInt8Val2: &expFloat64Val2,
				expInt8Val3: &expFloat64Val3}},
		{make(map[int8]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]int{
				expInt8Val1: expIntVal1, expInt8Val2: expIntVal2, expInt8Val3: expIntVal3}},
		{make(map[int8]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*int{
				expInt8Val1: &expIntVal1, expInt8Val2: &expIntVal2, expInt8Val3: &expIntVal3}},
		{make(map[int8]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]int16{
				expInt8Val1: expInt16Val1,
				expInt8Val2: expInt16Val2,
				expInt8Val3: expInt16Val3}},
		{make(map[int8]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*int16{
				expInt8Val1: &expInt16Val1,
				expInt8Val2: &expInt16Val2,
				expInt8Val3: &expInt16Val3}},
		{make(map[int8]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]int32{
				expInt8Val1: expInt32Val1,
				expInt8Val2: expInt32Val2,
				expInt8Val3: expInt32Val3}},
		{make(map[int8]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*int32{
				expInt8Val1: &expInt32Val1,
				expInt8Val2: &expInt32Val2,
				expInt8Val3: &expInt32Val3}},
		{make(map[int8]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]int64{
				expInt8Val1: expInt64Val1,
				expInt8Val2: expInt64Val2,
				expInt8Val3: expInt64Val3}},
		{make(map[int8]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*int64{
				expInt8Val1: &expInt64Val1,
				expInt8Val2: &expInt64Val2,
				expInt8Val3: &expInt64Val3}},
		{make(map[int8]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]int8{
				expInt8Val1: expInt8Val1, expInt8Val2: expInt8Val2, expInt8Val3: expInt8Val3}},
		{make(map[int8]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*int8{
				expInt8Val1: &expInt8Val1,
				expInt8Val2: &expInt8Val2,
				expInt8Val3: &expInt8Val3}},
		{make(map[int8]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int8]string{
				expInt8Val1: expStringVal1,
				expInt8Val2: expStringVal2,
				expInt8Val3: expStringVal3}},
		{make(map[int8]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[int8]*string{
				expInt8Val1: &expStringVal1,
				expInt8Val2: &expStringVal2,
				expInt8Val3: &expStringVal3}},
		{make(map[int8]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int8]time.Time{
				expInt8Val1: expTimeVal1, expInt8Val2: expTimeVal2, expInt8Val3: expTimeVal3}},
		{make(map[int8]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[int8]*time.Time{
				expInt8Val1: &expTimeVal1,
				expInt8Val2: &expTimeVal2,
				expInt8Val3: &expTimeVal3}},
		{make(map[int8]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]uint{
				expInt8Val1: expUintVal1, expInt8Val2: expUintVal2, expInt8Val3: expUintVal3}},
		{make(map[int8]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*uint{
				expInt8Val1: &expUintVal1,
				expInt8Val2: &expUintVal2,
				expInt8Val3: &expUintVal3}},
		{make(map[int8]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]uint16{
				expInt8Val1: expUint16Val1,
				expInt8Val2: expUint16Val2,
				expInt8Val3: expUint16Val3}},
		{make(map[int8]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*uint16{
				expInt8Val1: &expUint16Val1,
				expInt8Val2: &expUint16Val2,
				expInt8Val3: &expUint16Val3}},
		{make(map[int8]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]uint32{
				expInt8Val1: expUint32Val1,
				expInt8Val2: expUint32Val2,
				expInt8Val3: expUint32Val3}},
		{make(map[int8]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*uint32{
				expInt8Val1: &expUint32Val1,
				expInt8Val2: &expUint32Val2,
				expInt8Val3: &expUint32Val3}},
		{make(map[int8]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]uint64{
				expInt8Val1: expUint64Val1,
				expInt8Val2: expUint64Val2,
				expInt8Val3: expUint64Val3}},
		{make(map[int8]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*uint64{
				expInt8Val1: &expUint64Val1,
				expInt8Val2: &expUint64Val2,
				expInt8Val3: &expUint64Val3}},
		{make(map[int8]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]uint8{
				expInt8Val1: expUint8Val1,
				expInt8Val2: expUint8Val2,
				expInt8Val3: expUint8Val3}},
		{make(map[int8]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[int8]*uint8{
				expInt8Val1: &expUint8Val1,
				expInt8Val2: &expUint8Val2,
				expInt8Val3: &expUint8Val3}},
		{make(map[string]bool),
			map[string]string{"k1": "Yes", "K2": "FALSE"},
			map[string]bool{expStringVal1: expBoolVal1, expStringVal2: expBoolVal2}},
		{make(map[string]*bool),
			map[string]string{"k1": "Yes", "K2": "FALSE"},
			map[string]*bool{expStringVal1: &expBoolVal1, expStringVal2: &expBoolVal2}},
		{make(map[string]time.Duration),
			map[string]string{"k1": "10ns", "K2": "20µs", "03": "30ms"},
			map[string]time.Duration{
				expStringVal1: expDurationVal1,
				expStringVal2: expDurationVal2,
				expStringVal3: expDurationVal3}},
		{make(map[string]*time.Duration),
			map[string]string{"k1": "10ns", "K2": "20µs", "03": "30ms"},
			map[string]*time.Duration{
				expStringVal1: &expDurationVal1,
				expStringVal2: &expDurationVal2,
				expStringVal3: &expDurationVal3}},
		{make(map[string]float32),
			map[string]string{"k1": "1.2", "K2": "3.45", "03": "6.78"},
			map[string]float32{
				expStringVal1: expFloat32Val1,
				expStringVal2: expFloat32Val2,
				expStringVal3: expFloat32Val3}},
		{make(map[string]*float32),
			map[string]string{"k1": "1.2", "K2": "3.45", "03": "6.78"},
			map[string]*float32{
				expStringVal1: &expFloat32Val1,
				expStringVal2: &expFloat32Val2,
				expStringVal3: &expFloat32Val3}},
		{make(map[string]float64),
			map[string]string{"k1": "1.2", "K2": "3.45", "03": "6.78"},
			map[string]float64{
				expStringVal1: expFloat64Val1,
				expStringVal2: expFloat64Val2,
				expStringVal3: expFloat64Val3}},
		{make(map[string]*float64),
			map[string]string{"k1": "1.2", "K2": "3.45", "03": "6.78"},
			map[string]*float64{
				expStringVal1: &expFloat64Val1,
				expStringVal2: &expFloat64Val2,
				expStringVal3: &expFloat64Val3}},
		{make(map[string]int),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]int{
				expStringVal1: expIntVal1,
				expStringVal2: expIntVal2,
				expStringVal3: expIntVal3}},
		{make(map[string]*int),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*int{
				expStringVal1: &expIntVal1,
				expStringVal2: &expIntVal2,
				expStringVal3: &expIntVal3}},
		{make(map[string]int16),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]int16{
				expStringVal1: expInt16Val1,
				expStringVal2: expInt16Val2,
				expStringVal3: expInt16Val3}},
		{make(map[string]*int16),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*int16{
				expStringVal1: &expInt16Val1,
				expStringVal2: &expInt16Val2,
				expStringVal3: &expInt16Val3}},
		{make(map[string]int32),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]int32{
				expStringVal1: expInt32Val1,
				expStringVal2: expInt32Val2,
				expStringVal3: expInt32Val3}},
		{make(map[string]*int32),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*int32{
				expStringVal1: &expInt32Val1,
				expStringVal2: &expInt32Val2,
				expStringVal3: &expInt32Val3}},
		{make(map[string]int64),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]int64{
				expStringVal1: expInt64Val1,
				expStringVal2: expInt64Val2,
				expStringVal3: expInt64Val3}},
		{make(map[string]*int64),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*int64{
				expStringVal1: &expInt64Val1,
				expStringVal2: &expInt64Val2,
				expStringVal3: &expInt64Val3}},
		{make(map[string]int8),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]int8{
				expStringVal1: expInt8Val1,
				expStringVal2: expInt8Val2,
				expStringVal3: expInt8Val3}},
		{make(map[string]*int8),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*int8{
				expStringVal1: &expInt8Val1,
				expStringVal2: &expInt8Val2,
				expStringVal3: &expInt8Val3}},
		{make(map[string]string),
			map[string]string{"k1": "k1", "K2": "K2", "03": "03"},
			map[string]string{
				expStringVal1: expStringVal1,
				expStringVal2: expStringVal2,
				expStringVal3: expStringVal3}},
		{make(map[string]*string),
			map[string]string{"k1": "k1", "K2": "K2", "03": "03"},
			map[string]*string{
				expStringVal1: &expStringVal1,
				expStringVal2: &expStringVal2,
				expStringVal3: &expStringVal3}},
		{make(map[string]time.Time),
			map[string]string{
				"k1": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"K2": "Mon, 2 Jan 16:04:05 UTC 2006",
				"03": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[string]time.Time{
				expStringVal1: expTimeVal1,
				expStringVal2: expTimeVal2,
				expStringVal3: expTimeVal3}},
		{make(map[string]*time.Time),
			map[string]string{
				"k1": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"K2": "Mon, 2 Jan 16:04:05 UTC 2006",
				"03": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[string]*time.Time{
				expStringVal1: &expTimeVal1,
				expStringVal2: &expTimeVal2,
				expStringVal3: &expTimeVal3}},
		{make(map[string]uint),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]uint{
				expStringVal1: expUintVal1,
				expStringVal2: expUintVal2,
				expStringVal3: expUintVal3}},
		{make(map[string]*uint),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*uint{
				expStringVal1: &expUintVal1,
				expStringVal2: &expUintVal2,
				expStringVal3: &expUintVal3}},
		{make(map[string]uint16),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]uint16{
				expStringVal1: expUint16Val1,
				expStringVal2: expUint16Val2,
				expStringVal3: expUint16Val3}},
		{make(map[string]*uint16),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*uint16{
				expStringVal1: &expUint16Val1,
				expStringVal2: &expUint16Val2,
				expStringVal3: &expUint16Val3}},
		{make(map[string]uint32),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]uint32{
				expStringVal1: expUint32Val1,
				expStringVal2: expUint32Val2,
				expStringVal3: expUint32Val3}},
		{make(map[string]*uint32),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*uint32{
				expStringVal1: &expUint32Val1,
				expStringVal2: &expUint32Val2,
				expStringVal3: &expUint32Val3}},
		{make(map[string]uint64),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]uint64{
				expStringVal1: expUint64Val1,
				expStringVal2: expUint64Val2,
				expStringVal3: expUint64Val3}},
		{make(map[string]*uint64),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*uint64{
				expStringVal1: &expUint64Val1,
				expStringVal2: &expUint64Val2,
				expStringVal3: &expUint64Val3}},
		{make(map[string]uint8),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]uint8{
				expStringVal1: expUint8Val1,
				expStringVal2: expUint8Val2,
				expStringVal3: expUint8Val3}},
		{make(map[string]*uint8),
			map[string]string{"k1": "12", "K2": "34", "03": "56"},
			map[string]*uint8{
				expStringVal1: &expUint8Val1,
				expStringVal2: &expUint8Val2,
				expStringVal3: &expUint8Val3}},
		{make(map[time.Time]bool),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "Yes",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "FALSE"},
			map[time.Time]bool{expTimeVal1: expBoolVal1, expTimeVal2: expBoolVal2}},
		{make(map[time.Time]*bool),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "Yes",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "FALSE"},
			map[time.Time]*bool{expTimeVal1: &expBoolVal1, expTimeVal2: &expBoolVal2}},
		{make(map[time.Time]time.Duration),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "10ns",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "20µs",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "30ms"},
			map[time.Time]time.Duration{
				expTimeVal1: expDurationVal1,
				expTimeVal2: expDurationVal2,
				expTimeVal3: expDurationVal3}},
		{make(map[time.Time]*time.Duration),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "10ns",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "20µs",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "30ms"},
			map[time.Time]*time.Duration{
				expTimeVal1: &expDurationVal1,
				expTimeVal2: &expDurationVal2,
				expTimeVal3: &expDurationVal3}},
		{make(map[time.Time]float32),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "1.2",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "3.45",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "6.78"},
			map[time.Time]float32{
				expTimeVal1: expFloat32Val1,
				expTimeVal2: expFloat32Val2,
				expTimeVal3: expFloat32Val3}},
		{make(map[time.Time]*float32),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "1.2",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "3.45",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "6.78"},
			map[time.Time]*float32{
				expTimeVal1: &expFloat32Val1,
				expTimeVal2: &expFloat32Val2,
				expTimeVal3: &expFloat32Val3}},
		{make(map[time.Time]float64),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "1.2",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "3.45",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "6.78"},
			map[time.Time]float64{
				expTimeVal1: expFloat64Val1,
				expTimeVal2: expFloat64Val2,
				expTimeVal3: expFloat64Val3}},
		{make(map[time.Time]*float64),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "1.2",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "3.45",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "6.78"},
			map[time.Time]*float64{
				expTimeVal1: &expFloat64Val1,
				expTimeVal2: &expFloat64Val2,
				expTimeVal3: &expFloat64Val3}},
		{make(map[time.Time]int),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]int{
				expTimeVal1: expIntVal1, expTimeVal2: expIntVal2, expTimeVal3: expIntVal3}},
		{make(map[time.Time]*int),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*int{
				expTimeVal1: &expIntVal1, expTimeVal2: &expIntVal2, expTimeVal3: &expIntVal3}},
		{make(map[time.Time]int16),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]int16{
				expTimeVal1: expInt16Val1,
				expTimeVal2: expInt16Val2,
				expTimeVal3: expInt16Val3}},
		{make(map[time.Time]*int16),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*int16{
				expTimeVal1: &expInt16Val1,
				expTimeVal2: &expInt16Val2,
				expTimeVal3: &expInt16Val3}},
		{make(map[time.Time]int32),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]int32{
				expTimeVal1: expInt32Val1,
				expTimeVal2: expInt32Val2,
				expTimeVal3: expInt32Val3}},
		{make(map[time.Time]*int32),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*int32{
				expTimeVal1: &expInt32Val1,
				expTimeVal2: &expInt32Val2,
				expTimeVal3: &expInt32Val3}},
		{make(map[time.Time]int64),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]int64{
				expTimeVal1: expInt64Val1,
				expTimeVal2: expInt64Val2,
				expTimeVal3: expInt64Val3}},
		{make(map[time.Time]*int64),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*int64{
				expTimeVal1: &expInt64Val1,
				expTimeVal2: &expInt64Val2,
				expTimeVal3: &expInt64Val3}},
		{make(map[time.Time]int8),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]int8{
				expTimeVal1: expInt8Val1, expTimeVal2: expInt8Val2, expTimeVal3: expInt8Val3}},
		{make(map[time.Time]*int8),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*int8{
				expTimeVal1: &expInt8Val1,
				expTimeVal2: &expInt8Val2,
				expTimeVal3: &expInt8Val3}},
		{make(map[time.Time]string),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "k1",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "K2",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "03"},
			map[time.Time]string{
				expTimeVal1: expStringVal1,
				expTimeVal2: expStringVal2,
				expTimeVal3: expStringVal3}},
		{make(map[time.Time]*string),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "k1",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "K2",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "03"},
			map[time.Time]*string{
				expTimeVal1: &expStringVal1,
				expTimeVal2: &expStringVal2,
				expTimeVal3: &expStringVal3}},
		{make(map[time.Time]time.Time),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "Mon, 2 Jan 16:04:05 UTC 2006",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[time.Time]time.Time{
				expTimeVal1: expTimeVal1, expTimeVal2: expTimeVal2, expTimeVal3: expTimeVal3}},
		{make(map[time.Time]*time.Time),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "Mon, 2 Jan 16:04:05 UTC 2006",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[time.Time]*time.Time{
				expTimeVal1: &expTimeVal1,
				expTimeVal2: &expTimeVal2,
				expTimeVal3: &expTimeVal3}},
		{make(map[time.Time]uint),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]uint{
				expTimeVal1: expUintVal1, expTimeVal2: expUintVal2, expTimeVal3: expUintVal3}},
		{make(map[time.Time]*uint),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*uint{
				expTimeVal1: &expUintVal1,
				expTimeVal2: &expUintVal2,
				expTimeVal3: &expUintVal3}},
		{make(map[time.Time]uint16),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]uint16{
				expTimeVal1: expUint16Val1,
				expTimeVal2: expUint16Val2,
				expTimeVal3: expUint16Val3}},
		{make(map[time.Time]*uint16),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*uint16{
				expTimeVal1: &expUint16Val1,
				expTimeVal2: &expUint16Val2,
				expTimeVal3: &expUint16Val3}},
		{make(map[time.Time]uint32),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]uint32{
				expTimeVal1: expUint32Val1,
				expTimeVal2: expUint32Val2,
				expTimeVal3: expUint32Val3}},
		{make(map[time.Time]*uint32),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*uint32{
				expTimeVal1: &expUint32Val1,
				expTimeVal2: &expUint32Val2,
				expTimeVal3: &expUint32Val3}},
		{make(map[time.Time]uint64),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]uint64{
				expTimeVal1: expUint64Val1,
				expTimeVal2: expUint64Val2,
				expTimeVal3: expUint64Val3}},
		{make(map[time.Time]*uint64),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*uint64{
				expTimeVal1: &expUint64Val1,
				expTimeVal2: &expUint64Val2,
				expTimeVal3: &expUint64Val3}},
		{make(map[time.Time]uint8),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]uint8{
				expTimeVal1: expUint8Val1,
				expTimeVal2: expUint8Val2,
				expTimeVal3: expUint8Val3}},
		{make(map[time.Time]*uint8),
			map[string]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)": "12",
				"Mon, 2 Jan 16:04:05 UTC 2006":    "34",
				"Mon, 02 Jan 2006 17:04:05 (UTC)": "56"},
			map[time.Time]*uint8{
				expTimeVal1: &expUint8Val1,
				expTimeVal2: &expUint8Val2,
				expTimeVal3: &expUint8Val3}},
		{make(map[uint]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint]bool{expUintVal1: expBoolVal1, expUintVal2: expBoolVal2}},
		{make(map[uint]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint]*bool{expUintVal1: &expBoolVal1, expUintVal2: &expBoolVal2}},
		{make(map[uint]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint]time.Duration{
				expUintVal1: expDurationVal1,
				expUintVal2: expDurationVal2,
				expUintVal3: expDurationVal3}},
		{make(map[uint]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint]*time.Duration{
				expUintVal1: &expDurationVal1,
				expUintVal2: &expDurationVal2,
				expUintVal3: &expDurationVal3}},
		{make(map[uint]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint]float32{
				expUintVal1: expFloat32Val1,
				expUintVal2: expFloat32Val2,
				expUintVal3: expFloat32Val3}},
		{make(map[uint]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint]*float32{
				expUintVal1: &expFloat32Val1,
				expUintVal2: &expFloat32Val2,
				expUintVal3: &expFloat32Val3}},
		{make(map[uint]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint]float64{
				expUintVal1: expFloat64Val1,
				expUintVal2: expFloat64Val2,
				expUintVal3: expFloat64Val3}},
		{make(map[uint]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint]*float64{
				expUintVal1: &expFloat64Val1,
				expUintVal2: &expFloat64Val2,
				expUintVal3: &expFloat64Val3}},
		{make(map[uint]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]int{
				expUintVal1: expIntVal1, expUintVal2: expIntVal2, expUintVal3: expIntVal3}},
		{make(map[uint]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*int{
				expUintVal1: &expIntVal1, expUintVal2: &expIntVal2, expUintVal3: &expIntVal3}},
		{make(map[uint]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]int16{
				expUintVal1: expInt16Val1,
				expUintVal2: expInt16Val2,
				expUintVal3: expInt16Val3}},
		{make(map[uint]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*int16{
				expUintVal1: &expInt16Val1,
				expUintVal2: &expInt16Val2,
				expUintVal3: &expInt16Val3}},
		{make(map[uint]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]int32{
				expUintVal1: expInt32Val1,
				expUintVal2: expInt32Val2,
				expUintVal3: expInt32Val3}},
		{make(map[uint]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*int32{
				expUintVal1: &expInt32Val1,
				expUintVal2: &expInt32Val2,
				expUintVal3: &expInt32Val3}},
		{make(map[uint]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]int64{
				expUintVal1: expInt64Val1,
				expUintVal2: expInt64Val2,
				expUintVal3: expInt64Val3}},
		{make(map[uint]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*int64{
				expUintVal1: &expInt64Val1,
				expUintVal2: &expInt64Val2,
				expUintVal3: &expInt64Val3}},
		{make(map[uint]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]int8{
				expUintVal1: expInt8Val1, expUintVal2: expInt8Val2, expUintVal3: expInt8Val3}},
		{make(map[uint]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*int8{
				expUintVal1: &expInt8Val1,
				expUintVal2: &expInt8Val2,
				expUintVal3: &expInt8Val3}},
		{make(map[uint]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint]string{
				expUintVal1: expStringVal1,
				expUintVal2: expStringVal2,
				expUintVal3: expStringVal3}},
		{make(map[uint]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint]*string{
				expUintVal1: &expStringVal1,
				expUintVal2: &expStringVal2,
				expUintVal3: &expStringVal3}},
		{make(map[uint]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint]time.Time{
				expUintVal1: expTimeVal1, expUintVal2: expTimeVal2, expUintVal3: expTimeVal3}},
		{make(map[uint]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint]*time.Time{
				expUintVal1: &expTimeVal1,
				expUintVal2: &expTimeVal2,
				expUintVal3: &expTimeVal3}},
		{make(map[uint]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]uint{
				expUintVal1: expUintVal1, expUintVal2: expUintVal2, expUintVal3: expUintVal3}},
		{make(map[uint]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*uint{
				expUintVal1: &expUintVal1,
				expUintVal2: &expUintVal2,
				expUintVal3: &expUintVal3}},
		{make(map[uint]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]uint16{
				expUintVal1: expUint16Val1,
				expUintVal2: expUint16Val2,
				expUintVal3: expUint16Val3}},
		{make(map[uint]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*uint16{
				expUintVal1: &expUint16Val1,
				expUintVal2: &expUint16Val2,
				expUintVal3: &expUint16Val3}},
		{make(map[uint]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]uint32{
				expUintVal1: expUint32Val1,
				expUintVal2: expUint32Val2,
				expUintVal3: expUint32Val3}},
		{make(map[uint]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*uint32{
				expUintVal1: &expUint32Val1,
				expUintVal2: &expUint32Val2,
				expUintVal3: &expUint32Val3}},
		{make(map[uint]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]uint64{
				expUintVal1: expUint64Val1,
				expUintVal2: expUint64Val2,
				expUintVal3: expUint64Val3}},
		{make(map[uint]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*uint64{
				expUintVal1: &expUint64Val1,
				expUintVal2: &expUint64Val2,
				expUintVal3: &expUint64Val3}},
		{make(map[uint]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]uint8{
				expUintVal1: expUint8Val1,
				expUintVal2: expUint8Val2,
				expUintVal3: expUint8Val3}},
		{make(map[uint]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint]*uint8{
				expUintVal1: &expUint8Val1,
				expUintVal2: &expUint8Val2,
				expUintVal3: &expUint8Val3}},
		{make(map[uint16]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint16]bool{expUint16Val1: expBoolVal1, expUint16Val2: expBoolVal2}},
		{make(map[uint16]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint16]*bool{expUint16Val1: &expBoolVal1, expUint16Val2: &expBoolVal2}},
		{make(map[uint16]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint16]time.Duration{
				expUint16Val1: expDurationVal1,
				expUint16Val2: expDurationVal2,
				expUint16Val3: expDurationVal3}},
		{make(map[uint16]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint16]*time.Duration{
				expUint16Val1: &expDurationVal1,
				expUint16Val2: &expDurationVal2,
				expUint16Val3: &expDurationVal3}},
		{make(map[uint16]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint16]float32{
				expUint16Val1: expFloat32Val1,
				expUint16Val2: expFloat32Val2,
				expUint16Val3: expFloat32Val3}},
		{make(map[uint16]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint16]*float32{
				expUint16Val1: &expFloat32Val1,
				expUint16Val2: &expFloat32Val2,
				expUint16Val3: &expFloat32Val3}},
		{make(map[uint16]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint16]float64{
				expUint16Val1: expFloat64Val1,
				expUint16Val2: expFloat64Val2,
				expUint16Val3: expFloat64Val3}},
		{make(map[uint16]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint16]*float64{
				expUint16Val1: &expFloat64Val1,
				expUint16Val2: &expFloat64Val2,
				expUint16Val3: &expFloat64Val3}},
		{make(map[uint16]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]int{
				expUint16Val1: expIntVal1,
				expUint16Val2: expIntVal2,
				expUint16Val3: expIntVal3}},
		{make(map[uint16]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*int{
				expUint16Val1: &expIntVal1,
				expUint16Val2: &expIntVal2,
				expUint16Val3: &expIntVal3}},
		{make(map[uint16]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]int16{
				expUint16Val1: expInt16Val1,
				expUint16Val2: expInt16Val2,
				expUint16Val3: expInt16Val3}},
		{make(map[uint16]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*int16{
				expUint16Val1: &expInt16Val1,
				expUint16Val2: &expInt16Val2,
				expUint16Val3: &expInt16Val3}},
		{make(map[uint16]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]int32{
				expUint16Val1: expInt32Val1,
				expUint16Val2: expInt32Val2,
				expUint16Val3: expInt32Val3}},
		{make(map[uint16]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*int32{
				expUint16Val1: &expInt32Val1,
				expUint16Val2: &expInt32Val2,
				expUint16Val3: &expInt32Val3}},
		{make(map[uint16]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]int64{
				expUint16Val1: expInt64Val1,
				expUint16Val2: expInt64Val2,
				expUint16Val3: expInt64Val3}},
		{make(map[uint16]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*int64{
				expUint16Val1: &expInt64Val1,
				expUint16Val2: &expInt64Val2,
				expUint16Val3: &expInt64Val3}},
		{make(map[uint16]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]int8{
				expUint16Val1: expInt8Val1,
				expUint16Val2: expInt8Val2,
				expUint16Val3: expInt8Val3}},
		{make(map[uint16]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*int8{
				expUint16Val1: &expInt8Val1,
				expUint16Val2: &expInt8Val2,
				expUint16Val3: &expInt8Val3}},
		{make(map[uint16]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint16]string{
				expUint16Val1: expStringVal1,
				expUint16Val2: expStringVal2,
				expUint16Val3: expStringVal3}},
		{make(map[uint16]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint16]*string{
				expUint16Val1: &expStringVal1,
				expUint16Val2: &expStringVal2,
				expUint16Val3: &expStringVal3}},
		{make(map[uint16]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint16]time.Time{
				expUint16Val1: expTimeVal1,
				expUint16Val2: expTimeVal2,
				expUint16Val3: expTimeVal3}},
		{make(map[uint16]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint16]*time.Time{
				expUint16Val1: &expTimeVal1,
				expUint16Val2: &expTimeVal2,
				expUint16Val3: &expTimeVal3}},
		{make(map[uint16]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]uint{
				expUint16Val1: expUintVal1,
				expUint16Val2: expUintVal2,
				expUint16Val3: expUintVal3}},
		{make(map[uint16]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*uint{
				expUint16Val1: &expUintVal1,
				expUint16Val2: &expUintVal2,
				expUint16Val3: &expUintVal3}},
		{make(map[uint16]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]uint16{
				expUint16Val1: expUint16Val1,
				expUint16Val2: expUint16Val2,
				expUint16Val3: expUint16Val3}},
		{make(map[uint16]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*uint16{
				expUint16Val1: &expUint16Val1,
				expUint16Val2: &expUint16Val2,
				expUint16Val3: &expUint16Val3}},
		{make(map[uint16]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]uint32{
				expUint16Val1: expUint32Val1,
				expUint16Val2: expUint32Val2,
				expUint16Val3: expUint32Val3}},
		{make(map[uint16]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*uint32{
				expUint16Val1: &expUint32Val1,
				expUint16Val2: &expUint32Val2,
				expUint16Val3: &expUint32Val3}},
		{make(map[uint16]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]uint64{
				expUint16Val1: expUint64Val1,
				expUint16Val2: expUint64Val2,
				expUint16Val3: expUint64Val3}},
		{make(map[uint16]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*uint64{
				expUint16Val1: &expUint64Val1,
				expUint16Val2: &expUint64Val2,
				expUint16Val3: &expUint64Val3}},
		{make(map[uint16]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]uint8{
				expUint16Val1: expUint8Val1,
				expUint16Val2: expUint8Val2,
				expUint16Val3: expUint8Val3}},
		{make(map[uint16]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint16]*uint8{
				expUint16Val1: &expUint8Val1,
				expUint16Val2: &expUint8Val2,
				expUint16Val3: &expUint8Val3}},
		{make(map[uint32]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint32]bool{expUint32Val1: expBoolVal1, expUint32Val2: expBoolVal2}},
		{make(map[uint32]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint32]*bool{expUint32Val1: &expBoolVal1, expUint32Val2: &expBoolVal2}},
		{make(map[uint32]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint32]time.Duration{
				expUint32Val1: expDurationVal1,
				expUint32Val2: expDurationVal2,
				expUint32Val3: expDurationVal3}},
		{make(map[uint32]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint32]*time.Duration{
				expUint32Val1: &expDurationVal1,
				expUint32Val2: &expDurationVal2,
				expUint32Val3: &expDurationVal3}},
		{make(map[uint32]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint32]float32{
				expUint32Val1: expFloat32Val1,
				expUint32Val2: expFloat32Val2,
				expUint32Val3: expFloat32Val3}},
		{make(map[uint32]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint32]*float32{
				expUint32Val1: &expFloat32Val1,
				expUint32Val2: &expFloat32Val2,
				expUint32Val3: &expFloat32Val3}},
		{make(map[uint32]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint32]float64{
				expUint32Val1: expFloat64Val1,
				expUint32Val2: expFloat64Val2,
				expUint32Val3: expFloat64Val3}},
		{make(map[uint32]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint32]*float64{
				expUint32Val1: &expFloat64Val1,
				expUint32Val2: &expFloat64Val2,
				expUint32Val3: &expFloat64Val3}},
		{make(map[uint32]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]int{
				expUint32Val1: expIntVal1,
				expUint32Val2: expIntVal2,
				expUint32Val3: expIntVal3}},
		{make(map[uint32]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*int{
				expUint32Val1: &expIntVal1,
				expUint32Val2: &expIntVal2,
				expUint32Val3: &expIntVal3}},
		{make(map[uint32]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]int16{
				expUint32Val1: expInt16Val1,
				expUint32Val2: expInt16Val2,
				expUint32Val3: expInt16Val3}},
		{make(map[uint32]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*int16{
				expUint32Val1: &expInt16Val1,
				expUint32Val2: &expInt16Val2,
				expUint32Val3: &expInt16Val3}},
		{make(map[uint32]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]int32{
				expUint32Val1: expInt32Val1,
				expUint32Val2: expInt32Val2,
				expUint32Val3: expInt32Val3}},
		{make(map[uint32]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*int32{
				expUint32Val1: &expInt32Val1,
				expUint32Val2: &expInt32Val2,
				expUint32Val3: &expInt32Val3}},
		{make(map[uint32]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]int64{
				expUint32Val1: expInt64Val1,
				expUint32Val2: expInt64Val2,
				expUint32Val3: expInt64Val3}},
		{make(map[uint32]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*int64{
				expUint32Val1: &expInt64Val1,
				expUint32Val2: &expInt64Val2,
				expUint32Val3: &expInt64Val3}},
		{make(map[uint32]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]int8{
				expUint32Val1: expInt8Val1,
				expUint32Val2: expInt8Val2,
				expUint32Val3: expInt8Val3}},
		{make(map[uint32]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*int8{
				expUint32Val1: &expInt8Val1,
				expUint32Val2: &expInt8Val2,
				expUint32Val3: &expInt8Val3}},
		{make(map[uint32]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint32]string{
				expUint32Val1: expStringVal1,
				expUint32Val2: expStringVal2,
				expUint32Val3: expStringVal3}},
		{make(map[uint32]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint32]*string{
				expUint32Val1: &expStringVal1,
				expUint32Val2: &expStringVal2,
				expUint32Val3: &expStringVal3}},
		{make(map[uint32]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint32]time.Time{
				expUint32Val1: expTimeVal1,
				expUint32Val2: expTimeVal2,
				expUint32Val3: expTimeVal3}},
		{make(map[uint32]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint32]*time.Time{
				expUint32Val1: &expTimeVal1,
				expUint32Val2: &expTimeVal2,
				expUint32Val3: &expTimeVal3}},
		{make(map[uint32]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]uint{
				expUint32Val1: expUintVal1,
				expUint32Val2: expUintVal2,
				expUint32Val3: expUintVal3}},
		{make(map[uint32]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*uint{
				expUint32Val1: &expUintVal1,
				expUint32Val2: &expUintVal2,
				expUint32Val3: &expUintVal3}},
		{make(map[uint32]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]uint16{
				expUint32Val1: expUint16Val1,
				expUint32Val2: expUint16Val2,
				expUint32Val3: expUint16Val3}},
		{make(map[uint32]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*uint16{
				expUint32Val1: &expUint16Val1,
				expUint32Val2: &expUint16Val2,
				expUint32Val3: &expUint16Val3}},
		{make(map[uint32]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]uint32{
				expUint32Val1: expUint32Val1,
				expUint32Val2: expUint32Val2,
				expUint32Val3: expUint32Val3}},
		{make(map[uint32]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*uint32{
				expUint32Val1: &expUint32Val1,
				expUint32Val2: &expUint32Val2,
				expUint32Val3: &expUint32Val3}},
		{make(map[uint32]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]uint64{
				expUint32Val1: expUint64Val1,
				expUint32Val2: expUint64Val2,
				expUint32Val3: expUint64Val3}},
		{make(map[uint32]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*uint64{
				expUint32Val1: &expUint64Val1,
				expUint32Val2: &expUint64Val2,
				expUint32Val3: &expUint64Val3}},
		{make(map[uint32]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]uint8{
				expUint32Val1: expUint8Val1,
				expUint32Val2: expUint8Val2,
				expUint32Val3: expUint8Val3}},
		{make(map[uint32]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint32]*uint8{
				expUint32Val1: &expUint8Val1,
				expUint32Val2: &expUint8Val2,
				expUint32Val3: &expUint8Val3}},
		{make(map[uint64]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint64]bool{expUint64Val1: expBoolVal1, expUint64Val2: expBoolVal2}},
		{make(map[uint64]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint64]*bool{expUint64Val1: &expBoolVal1, expUint64Val2: &expBoolVal2}},
		{make(map[uint64]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint64]time.Duration{
				expUint64Val1: expDurationVal1,
				expUint64Val2: expDurationVal2,
				expUint64Val3: expDurationVal3}},
		{make(map[uint64]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint64]*time.Duration{
				expUint64Val1: &expDurationVal1,
				expUint64Val2: &expDurationVal2,
				expUint64Val3: &expDurationVal3}},
		{make(map[uint64]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint64]float32{
				expUint64Val1: expFloat32Val1,
				expUint64Val2: expFloat32Val2,
				expUint64Val3: expFloat32Val3}},
		{make(map[uint64]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint64]*float32{
				expUint64Val1: &expFloat32Val1,
				expUint64Val2: &expFloat32Val2,
				expUint64Val3: &expFloat32Val3}},
		{make(map[uint64]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint64]float64{
				expUint64Val1: expFloat64Val1,
				expUint64Val2: expFloat64Val2,
				expUint64Val3: expFloat64Val3}},
		{make(map[uint64]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint64]*float64{
				expUint64Val1: &expFloat64Val1,
				expUint64Val2: &expFloat64Val2,
				expUint64Val3: &expFloat64Val3}},
		{make(map[uint64]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]int{
				expUint64Val1: expIntVal1,
				expUint64Val2: expIntVal2,
				expUint64Val3: expIntVal3}},
		{make(map[uint64]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*int{
				expUint64Val1: &expIntVal1,
				expUint64Val2: &expIntVal2,
				expUint64Val3: &expIntVal3}},
		{make(map[uint64]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]int16{
				expUint64Val1: expInt16Val1,
				expUint64Val2: expInt16Val2,
				expUint64Val3: expInt16Val3}},
		{make(map[uint64]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*int16{
				expUint64Val1: &expInt16Val1,
				expUint64Val2: &expInt16Val2,
				expUint64Val3: &expInt16Val3}},
		{make(map[uint64]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]int32{
				expUint64Val1: expInt32Val1,
				expUint64Val2: expInt32Val2,
				expUint64Val3: expInt32Val3}},
		{make(map[uint64]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*int32{
				expUint64Val1: &expInt32Val1,
				expUint64Val2: &expInt32Val2,
				expUint64Val3: &expInt32Val3}},
		{make(map[uint64]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]int64{
				expUint64Val1: expInt64Val1,
				expUint64Val2: expInt64Val2,
				expUint64Val3: expInt64Val3}},
		{make(map[uint64]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*int64{
				expUint64Val1: &expInt64Val1,
				expUint64Val2: &expInt64Val2,
				expUint64Val3: &expInt64Val3}},
		{make(map[uint64]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]int8{
				expUint64Val1: expInt8Val1,
				expUint64Val2: expInt8Val2,
				expUint64Val3: expInt8Val3}},
		{make(map[uint64]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*int8{
				expUint64Val1: &expInt8Val1,
				expUint64Val2: &expInt8Val2,
				expUint64Val3: &expInt8Val3}},
		{make(map[uint64]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint64]string{
				expUint64Val1: expStringVal1,
				expUint64Val2: expStringVal2,
				expUint64Val3: expStringVal3}},
		{make(map[uint64]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint64]*string{
				expUint64Val1: &expStringVal1,
				expUint64Val2: &expStringVal2,
				expUint64Val3: &expStringVal3}},
		{make(map[uint64]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint64]time.Time{
				expUint64Val1: expTimeVal1,
				expUint64Val2: expTimeVal2,
				expUint64Val3: expTimeVal3}},
		{make(map[uint64]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint64]*time.Time{
				expUint64Val1: &expTimeVal1,
				expUint64Val2: &expTimeVal2,
				expUint64Val3: &expTimeVal3}},
		{make(map[uint64]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]uint{
				expUint64Val1: expUintVal1,
				expUint64Val2: expUintVal2,
				expUint64Val3: expUintVal3}},
		{make(map[uint64]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*uint{
				expUint64Val1: &expUintVal1,
				expUint64Val2: &expUintVal2,
				expUint64Val3: &expUintVal3}},
		{make(map[uint64]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]uint16{
				expUint64Val1: expUint16Val1,
				expUint64Val2: expUint16Val2,
				expUint64Val3: expUint16Val3}},
		{make(map[uint64]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*uint16{
				expUint64Val1: &expUint16Val1,
				expUint64Val2: &expUint16Val2,
				expUint64Val3: &expUint16Val3}},
		{make(map[uint64]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]uint32{
				expUint64Val1: expUint32Val1,
				expUint64Val2: expUint32Val2,
				expUint64Val3: expUint32Val3}},
		{make(map[uint64]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*uint32{
				expUint64Val1: &expUint32Val1,
				expUint64Val2: &expUint32Val2,
				expUint64Val3: &expUint32Val3}},
		{make(map[uint64]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]uint64{
				expUint64Val1: expUint64Val1,
				expUint64Val2: expUint64Val2,
				expUint64Val3: expUint64Val3}},
		{make(map[uint64]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*uint64{
				expUint64Val1: &expUint64Val1,
				expUint64Val2: &expUint64Val2,
				expUint64Val3: &expUint64Val3}},
		{make(map[uint64]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]uint8{
				expUint64Val1: expUint8Val1,
				expUint64Val2: expUint8Val2,
				expUint64Val3: expUint8Val3}},
		{make(map[uint64]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint64]*uint8{
				expUint64Val1: &expUint8Val1,
				expUint64Val2: &expUint8Val2,
				expUint64Val3: &expUint8Val3}},
		{make(map[uint8]bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint8]bool{expUint8Val1: expBoolVal1, expUint8Val2: expBoolVal2}},
		{make(map[uint8]*bool),
			map[string]string{"12": "Yes", "34": "FALSE"},
			map[uint8]*bool{expUint8Val1: &expBoolVal1, expUint8Val2: &expBoolVal2}},
		{make(map[uint8]time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint8]time.Duration{
				expUint8Val1: expDurationVal1,
				expUint8Val2: expDurationVal2,
				expUint8Val3: expDurationVal3}},
		{make(map[uint8]*time.Duration),
			map[string]string{"12": "10ns", "34": "20µs", "56": "30ms"},
			map[uint8]*time.Duration{
				expUint8Val1: &expDurationVal1,
				expUint8Val2: &expDurationVal2,
				expUint8Val3: &expDurationVal3}},
		{make(map[uint8]float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint8]float32{
				expUint8Val1: expFloat32Val1,
				expUint8Val2: expFloat32Val2,
				expUint8Val3: expFloat32Val3}},
		{make(map[uint8]*float32),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint8]*float32{
				expUint8Val1: &expFloat32Val1,
				expUint8Val2: &expFloat32Val2,
				expUint8Val3: &expFloat32Val3}},
		{make(map[uint8]float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint8]float64{
				expUint8Val1: expFloat64Val1,
				expUint8Val2: expFloat64Val2,
				expUint8Val3: expFloat64Val3}},
		{make(map[uint8]*float64),
			map[string]string{"12": "1.2", "34": "3.45", "56": "6.78"},
			map[uint8]*float64{
				expUint8Val1: &expFloat64Val1,
				expUint8Val2: &expFloat64Val2,
				expUint8Val3: &expFloat64Val3}},
		{make(map[uint8]int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]int{
				expUint8Val1: expIntVal1, expUint8Val2: expIntVal2, expUint8Val3: expIntVal3}},
		{make(map[uint8]*int),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*int{
				expUint8Val1: &expIntVal1,
				expUint8Val2: &expIntVal2,
				expUint8Val3: &expIntVal3}},
		{make(map[uint8]int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]int16{
				expUint8Val1: expInt16Val1,
				expUint8Val2: expInt16Val2,
				expUint8Val3: expInt16Val3}},
		{make(map[uint8]*int16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*int16{
				expUint8Val1: &expInt16Val1,
				expUint8Val2: &expInt16Val2,
				expUint8Val3: &expInt16Val3}},
		{make(map[uint8]int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]int32{
				expUint8Val1: expInt32Val1,
				expUint8Val2: expInt32Val2,
				expUint8Val3: expInt32Val3}},
		{make(map[uint8]*int32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*int32{
				expUint8Val1: &expInt32Val1,
				expUint8Val2: &expInt32Val2,
				expUint8Val3: &expInt32Val3}},
		{make(map[uint8]int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]int64{
				expUint8Val1: expInt64Val1,
				expUint8Val2: expInt64Val2,
				expUint8Val3: expInt64Val3}},
		{make(map[uint8]*int64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*int64{
				expUint8Val1: &expInt64Val1,
				expUint8Val2: &expInt64Val2,
				expUint8Val3: &expInt64Val3}},
		{make(map[uint8]int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]int8{
				expUint8Val1: expInt8Val1,
				expUint8Val2: expInt8Val2,
				expUint8Val3: expInt8Val3}},
		{make(map[uint8]*int8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*int8{
				expUint8Val1: &expInt8Val1,
				expUint8Val2: &expInt8Val2,
				expUint8Val3: &expInt8Val3}},
		{make(map[uint8]string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint8]string{
				expUint8Val1: expStringVal1,
				expUint8Val2: expStringVal2,
				expUint8Val3: expStringVal3}},
		{make(map[uint8]*string),
			map[string]string{"12": "k1", "34": "K2", "56": "03"},
			map[uint8]*string{
				expUint8Val1: &expStringVal1,
				expUint8Val2: &expStringVal2,
				expUint8Val3: &expStringVal3}},
		{make(map[uint8]time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint8]time.Time{
				expUint8Val1: expTimeVal1,
				expUint8Val2: expTimeVal2,
				expUint8Val3: expTimeVal3}},
		{make(map[uint8]*time.Time),
			map[string]string{
				"12": "2 Jan 2006 15:04:05 -0700 (UTC)",
				"34": "Mon, 2 Jan 16:04:05 UTC 2006",
				"56": "Mon, 02 Jan 2006 17:04:05 (UTC)"},
			map[uint8]*time.Time{
				expUint8Val1: &expTimeVal1,
				expUint8Val2: &expTimeVal2,
				expUint8Val3: &expTimeVal3}},
		{make(map[uint8]uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]uint{
				expUint8Val1: expUintVal1,
				expUint8Val2: expUintVal2,
				expUint8Val3: expUintVal3}},
		{make(map[uint8]*uint),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*uint{
				expUint8Val1: &expUintVal1,
				expUint8Val2: &expUintVal2,
				expUint8Val3: &expUintVal3}},
		{make(map[uint8]uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]uint16{
				expUint8Val1: expUint16Val1,
				expUint8Val2: expUint16Val2,
				expUint8Val3: expUint16Val3}},
		{make(map[uint8]*uint16),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*uint16{
				expUint8Val1: &expUint16Val1,
				expUint8Val2: &expUint16Val2,
				expUint8Val3: &expUint16Val3}},
		{make(map[uint8]uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]uint32{
				expUint8Val1: expUint32Val1,
				expUint8Val2: expUint32Val2,
				expUint8Val3: expUint32Val3}},
		{make(map[uint8]*uint32),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*uint32{
				expUint8Val1: &expUint32Val1,
				expUint8Val2: &expUint32Val2,
				expUint8Val3: &expUint32Val3}},
		{make(map[uint8]uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]uint64{
				expUint8Val1: expUint64Val1,
				expUint8Val2: expUint64Val2,
				expUint8Val3: expUint64Val3}},
		{make(map[uint8]*uint64),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*uint64{
				expUint8Val1: &expUint64Val1,
				expUint8Val2: &expUint64Val2,
				expUint8Val3: &expUint64Val3}},
		{make(map[uint8]uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]uint8{
				expUint8Val1: expUint8Val1,
				expUint8Val2: expUint8Val2,
				expUint8Val3: expUint8Val3}},
		{make(map[uint8]*uint8),
			map[string]string{"12": "12", "34": "34", "56": "56"},
			map[uint8]*uint8{
				expUint8Val1: &expUint8Val1,
				expUint8Val2: &expUint8Val2,
				expUint8Val3: &expUint8Val3}},
	}
}
