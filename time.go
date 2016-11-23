package conv

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"time"
)

var (
	emptyTime      = time.Time{}
	typeOfTime     = reflect.TypeOf(emptyTime)
	typeOfDuration = reflect.TypeOf(time.Duration(0))
)

func (c Conv) convStrToDuration(v string) (time.Duration, error) {
	if parsed, err := time.ParseDuration(v); err == nil {
		return parsed, nil
	}
	if parsed, err := strconv.ParseInt(v, 10, 0); err == nil {
		// @TODO This feels more natural but maybe add a option to disable since
		// it breaks Duration() -> String() -> Duration()
		return time.Duration(parsed) * time.Second, nil
	}
	if parsed, err := strconv.ParseFloat(v, 64); err == nil {
		return time.Duration(1e9 * parsed), nil
	}
	return 0, fmt.Errorf("cannot parse %#v (type string) as time.Duration", v)
}

func (c Conv) convNumToDuration(
	k reflect.Kind, v reflect.Value) (time.Duration, bool) {
	switch {
	case isKindInt(k):
		return time.Duration(v.Int()), true
	case isKindUint(k):
		T := v.Uint()
		if T > math.MaxInt64 {
			T = math.MaxInt64
		}
		return time.Duration(T), true
	case isKindFloat(k):
		T := v.Float()
		if math.IsNaN(T) || math.IsInf(T, 0) {
			return 0, true
		}
		return time.Duration(1e9 * T), true
	case isKindComplex(k):
		T := v.Complex()
		if cmplx.IsNaN(T) || cmplx.IsInf(T) {
			return 0, true
		}
		return time.Duration(1e9 * real(T)), true
	}
	return 0, false
}

// Duration attempts to convert the given value to time.Duration, returns the
// zero value and an error on failure.
func (c Conv) Duration(from interface{}) (time.Duration, error) {
	if T, ok := from.(string); ok {
		return c.convStrToDuration(T)
	} else if T, ok := from.(time.Duration); ok {
		return T, nil
	} else if c, ok := from.(interface {
		Duration() (time.Duration, error)
	}); ok {
		return c.Duration()
	}

	value := reflect.ValueOf(indirect(from))
	kind := value.Kind()
	switch {
	case reflect.String == kind:
		return c.convStrToDuration(value.String())
	case isKindNumeric(kind):
		if parsed, ok := c.convNumToDuration(kind, value); ok {
			return parsed, nil
		}
	case reflect.Struct == kind && value.CanInterface():
		v := value.Interface()
		if T, ok := v.(time.Time); ok {
			// @TODO I find this useful but is it WTF?
			return time.Since(T), nil
		}
	}
	return 0, newConvErr(from, "time.Duration")
}

// Time attempts to convert the given value to time.Time, returns the zero value
// of time.Time and an error on failure.
func (c Conv) Time(from interface{}) (time.Time, error) {
	if T, ok := from.(time.Time); ok {
		return T, nil
	} else if T, ok := from.(*time.Time); ok {
		return *T, nil
	} else if c, ok := from.(interface {
		Time() (time.Time, error)
	}); ok {
		return c.Time()
	}

	value := reflect.ValueOf(indirect(from))
	kind := value.Kind()
	switch {
	case reflect.String == kind:
		if T, ok := timeFromString(value.String()); ok {
			return T, nil
		}
	case reflect.Struct == kind:
		if value.Type().ConvertibleTo(typeOfTime) {
			valueConv := value.Convert(typeOfTime)
			if valueConv.CanInterface() {
				return valueConv.Interface().(time.Time), nil
			}
		}
		field := value.FieldByName("Time")
		if field.IsValid() && field.CanInterface() {
			return c.Time(field.Interface())
		}
	}
	return emptyTime, newConvErr(from, "time.Time")
}

type formatInfo struct {
	format string
	needed string
}

var formats = []formatInfo{
	{time.RFC3339Nano, ""},
	{time.RFC3339, ""},
	{time.RFC850, ""},
	{time.RFC1123, ""},
	{time.RFC1123Z, ""},
	{"02 Jan 06 15:04:05", ""},
	{"02 Jan 06 15:04:05 +-0700", ""},
	{"02 Jan 06 15:4:5 MST", ""},
	{"02 Jan 2006 15:04:05", ""},
	{"2 Jan 2006 15:04:05", ""},
	{"2 Jan 2006 15:04:05 MST", ""},
	{"2 Jan 2006 15:04:05 -0700", ""},
	{"2 Jan 2006 15:04:05 -0700 (MST)", ""},
	{"02 January 2006 15:04", ""},
	{"02 Jan 2006 15:04 MST", ""},
	{"02 Jan 2006 15:04:05 MST", ""},
	{"02 Jan 2006 15:04:05 -0700", ""},
	{"02 Jan 2006 15:04:05 -0700 (MST)", ""},
	{"Mon, 2 Jan  15:04:05 MST 2006", ""},
	{"Mon, 2 Jan 15:04:05 MST 2006", ""},
	{"Mon, 02 Jan 2006 15:04:05", ""},
	{"Mon, 02 Jan 2006 15:04:05 (MST)", ""},
	{"Mon, 2 Jan 2006 15:04:05", ""},
	{"Mon, 2 Jan 2006 15:04:05 MST", ""},
	{"Mon, 2 Jan 2006 15:04:05 -0700", ""},
	{"Mon, 2 Jan 2006 15:04:05 -0700 (MST)", ""},
	{"Mon, 02 Jan 06 15:04:05 MST", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700 MST", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700 (MST)", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700 (MST-07:00)", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700 (MST MST)", ""},
	{"Mon, 02 Jan 2006 15:04 -0700", ""},
	{"Mon, 02 Jan 2006 15:04 -0700 (MST)", ""},
	{"Mon Jan 02 15:05:05 2006 MST", ""},
	{"Monday, 02 Jan 2006 15:04 -0700", ""},
	{"Monday, 02 Jan 2006 15:04:05 -0700", ""},
	{time.UnixDate, ""},
	{time.RubyDate, ""},
	{time.RFC822, ""},
	{time.RFC822Z, ""},
}

// Quick google yields no date parsing libraries, first thing that came to mind
// was trying all the formats in time package. This is reasonable enough until
// I can find a decent lexer or polish up my "timey" Go lib. I am using the
// table of dates politely released into public domain by github.com/tomarus:
//   https://github.com/tomarus/parsedate/blob/master/parsedate.go
func timeFromString(s string) (time.Time, bool) {
	if len(s) == 0 {
		return time.Time{}, false
	}
	for _, f := range formats {
		_, err := time.Parse(f.format, s)
		if err != nil {
			continue
		}
		if t, err := time.Parse(
			f.format+f.needed, s+time.Now().Format(f.needed)); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}
