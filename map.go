package conv

import (
	"fmt"
	"time"

	iter "github.com/cstockton/go-iter"
)

// Map will perform conversion by inferring the key and element types from the
// given map and taking values from the given interface.
func (c Conv) Map(into, from interface{}) error {
	return recoverFn(func() error {
		fn := mapIterFn(c, into)
		if fn == nil {
			return fmt.Errorf("%T is not a supported map type", into)
		}
		return iter.Walk(from, fn)
	})
}

func mapIterFn(c Converter, into interface{}) func(iter.Pair) error {
	switch T := into.(type) {
	case map[bool]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[bool]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[bool]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Bool(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Duration]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Duration]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Duration(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float32]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float32]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[float64]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[float64]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Float64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int16]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int16]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int32]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int32]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int64]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int64]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[int8]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[int8]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Int8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[string]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[string]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.String(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[time.Time]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[time.Time]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Time(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint16]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint16]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint16(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint32]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint32]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint32(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint64]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint64]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint64(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	case map[uint8]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = v
			return nil
		}
	case map[uint8]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			k, err := c.Uint8(el.Key())
			if err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			T[k] = &v
			return nil
		}
	}
	return nil
}
