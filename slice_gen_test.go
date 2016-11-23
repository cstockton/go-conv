package conv

import (
	"errors"
	"reflect"
	"testing"
	"time"

	iter "github.com/cstockton/go-iter"
)

type sliceGenTest struct {
	into interface{}
	from []string
	exp  interface{}
}

var sliceGenTests []sliceGenTest

func TestSliceGen(t *testing.T) {
	var c Conv

	triggerAt, count, triggerErr := 0, 0, errors.New("triggerErr")
	cErr := errHookConv{c, func(from interface{}, err error) error {
		count++
		if triggerAt == count {
			return triggerErr
		}
		return err
	}}

	for _, sliceGenTest := range sliceGenTests {
		err := c.Slice(sliceGenTest.into, sliceGenTest.from)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(sliceGenTest.exp, indirect(sliceGenTest.into)) {
			t.Logf("from (%T) --> %[1]v", sliceGenTest.from)
			t.Fatalf("\nexp (%T) --> %[1]v\ngot (%[2]T) --> %[2]v",
				sliceGenTest.exp, sliceGenTest.into)
		}

		val := reflect.ValueOf(sliceGenTest.from).Index(0).Interface()
		fn := sliceIterFn(cErr, sliceGenTest.into)
		el := iter.NewPair(nil, 0, val, nil)
		elErr := iter.NewPair(nil, 0, val, triggerErr)

		// Trigger element err check
		if err := fn(elErr); err != triggerErr {
			t.Errorf("expected element err in mapIterFn %T", sliceGenTest.into)
		}

		// Trigger key err check
		triggerAt, count = 1, 0
		if err := fn(el); err != triggerErr {
			t.Errorf("expected val err in mapIterFn %T", sliceGenTest.into)
		}
	}
}

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

	sliceGenTests = []sliceGenTest{
		{new([]bool),
			[]string{"Yes", "FALSE"},
			[]bool{expBoolVal1, expBoolVal2}},
		{new([]*bool),
			[]string{"Yes", "FALSE"},
			[]*bool{&expBoolVal1, &expBoolVal2}},
		{new([]time.Duration),
			[]string{"10ns", "20µs", "30ms"},
			[]time.Duration{expDurationVal1, expDurationVal2, expDurationVal3}},
		{new([]*time.Duration),
			[]string{"10ns", "20µs", "30ms"},
			[]*time.Duration{&expDurationVal1, &expDurationVal2, &expDurationVal3}},
		{new([]float32),
			[]string{"1.2", "3.45", "6.78"},
			[]float32{expFloat32Val1, expFloat32Val2, expFloat32Val3}},
		{new([]*float32),
			[]string{"1.2", "3.45", "6.78"},
			[]*float32{&expFloat32Val1, &expFloat32Val2, &expFloat32Val3}},
		{new([]float64),
			[]string{"1.2", "3.45", "6.78"},
			[]float64{expFloat64Val1, expFloat64Val2, expFloat64Val3}},
		{new([]*float64),
			[]string{"1.2", "3.45", "6.78"},
			[]*float64{&expFloat64Val1, &expFloat64Val2, &expFloat64Val3}},
		{new([]int),
			[]string{"12", "34", "56"},
			[]int{expIntVal1, expIntVal2, expIntVal3}},
		{new([]*int),
			[]string{"12", "34", "56"},
			[]*int{&expIntVal1, &expIntVal2, &expIntVal3}},
		{new([]int16),
			[]string{"12", "34", "56"},
			[]int16{expInt16Val1, expInt16Val2, expInt16Val3}},
		{new([]*int16),
			[]string{"12", "34", "56"},
			[]*int16{&expInt16Val1, &expInt16Val2, &expInt16Val3}},
		{new([]int32),
			[]string{"12", "34", "56"},
			[]int32{expInt32Val1, expInt32Val2, expInt32Val3}},
		{new([]*int32),
			[]string{"12", "34", "56"},
			[]*int32{&expInt32Val1, &expInt32Val2, &expInt32Val3}},
		{new([]int64),
			[]string{"12", "34", "56"},
			[]int64{expInt64Val1, expInt64Val2, expInt64Val3}},
		{new([]*int64),
			[]string{"12", "34", "56"},
			[]*int64{&expInt64Val1, &expInt64Val2, &expInt64Val3}},
		{new([]int8),
			[]string{"12", "34", "56"},
			[]int8{expInt8Val1, expInt8Val2, expInt8Val3}},
		{new([]*int8),
			[]string{"12", "34", "56"},
			[]*int8{&expInt8Val1, &expInt8Val2, &expInt8Val3}},
		{new([]string),
			[]string{"k1", "K2", "03"},
			[]string{expStringVal1, expStringVal2, expStringVal3}},
		{new([]*string),
			[]string{"k1", "K2", "03"},
			[]*string{&expStringVal1, &expStringVal2, &expStringVal3}},
		{new([]time.Time),
			[]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)",
				"Mon, 2 Jan 16:04:05 UTC 2006",
				"Mon, 02 Jan 2006 17:04:05 (UTC)"},
			[]time.Time{expTimeVal1, expTimeVal2, expTimeVal3}},
		{new([]*time.Time),
			[]string{
				"2 Jan 2006 15:04:05 -0700 (UTC)",
				"Mon, 2 Jan 16:04:05 UTC 2006",
				"Mon, 02 Jan 2006 17:04:05 (UTC)"},
			[]*time.Time{&expTimeVal1, &expTimeVal2, &expTimeVal3}},
		{new([]uint),
			[]string{"12", "34", "56"},
			[]uint{expUintVal1, expUintVal2, expUintVal3}},
		{new([]*uint),
			[]string{"12", "34", "56"},
			[]*uint{&expUintVal1, &expUintVal2, &expUintVal3}},
		{new([]uint16),
			[]string{"12", "34", "56"},
			[]uint16{expUint16Val1, expUint16Val2, expUint16Val3}},
		{new([]*uint16),
			[]string{"12", "34", "56"},
			[]*uint16{&expUint16Val1, &expUint16Val2, &expUint16Val3}},
		{new([]uint32),
			[]string{"12", "34", "56"},
			[]uint32{expUint32Val1, expUint32Val2, expUint32Val3}},
		{new([]*uint32),
			[]string{"12", "34", "56"},
			[]*uint32{&expUint32Val1, &expUint32Val2, &expUint32Val3}},
		{new([]uint64),
			[]string{"12", "34", "56"},
			[]uint64{expUint64Val1, expUint64Val2, expUint64Val3}},
		{new([]*uint64),
			[]string{"12", "34", "56"},
			[]*uint64{&expUint64Val1, &expUint64Val2, &expUint64Val3}},
		{new([]uint8),
			[]string{"12", "34", "56"},
			[]uint8{expUint8Val1, expUint8Val2, expUint8Val3}},
		{new([]*uint8),
			[]string{"12", "34", "56"},
			[]*uint8{&expUint8Val1, &expUint8Val2, &expUint8Val3}},
	}
}
