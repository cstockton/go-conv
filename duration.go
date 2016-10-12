package conv

import (
	"strconv"
	"time"
)

// DurationConverter interface allows a value to be converted to a time.Duration.
type DurationConverter interface {
	Duration() time.Duration
}

// Duration converts the given value to a time.Duration.
func Duration(value interface{}) time.Duration {
	value = Indirect(value)

	switch T := value.(type) {
	case DurationConverter:
		if nil != T {
			return T.Duration()
		}
	case time.Duration:
		return T
	case string:
		if d, err := time.ParseDuration(T); err == nil {
			return time.Duration(d)
		}
		if d, err := strconv.ParseFloat(T, 64); err == nil {
			return time.Duration(d)
		}
		if d, err := strconv.ParseInt(T, 10, 64); err == nil {
			return time.Duration(d)
		}
	}
	return time.Duration(Int64(value))
}
