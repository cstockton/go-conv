package conv

import (
	"fmt"
	"time"
)

// TimeConverter interface allows a value to be converted to a time.Time.
type TimeConverter interface {
	Time() time.Time
}

// Time converts the given value to a time.Time.
func Time(value interface{}) time.Time {
	value = Indirect(value)

	switch T := value.(type) {
	case TimeConverter:
		if T != nil {
			return T.Time()
		}
	case time.Duration:
		return time.Now().Add(T)
	case time.Time:
		return T
	case string:
		if t, err := timeFromString(T); err == nil {
			return t
		}
	}
	return time.Time{}
}

type formatInfo struct {
	format string
	needed string
}

var formats = []formatInfo{
	{time.ANSIC, ""},
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
func timeFromString(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, fmt.Errorf("Failed to parse time: %s", s)
	}
	for _, f := range formats {
		t, err := time.Parse(f.format, s)
		if err != nil {
			continue
		}
		if t, err = time.Parse(f.format+f.needed, s+time.Now().Format(f.needed)); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("Failed to parse time: %s", s)
}
