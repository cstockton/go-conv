package conv

import (
	"fmt"
	"time"

	iter "github.com/cstockton/go-iter"
)

// Slice will perform conversion by inferring the element type from the given
// slice and taking values from the given interface.
func (c Conv) Slice(into interface{}, from interface{}) error {
	return recoverFn(func() error {
		fn := sliceIterFn(c, into)
		if fn == nil {
			return fmt.Errorf("%T is not a pointer to a supported slice type", into)
		}
		return iter.Walk(from, fn)
	})
}

func sliceIterFn(c Converter, into interface{}) func(iter.Pair) error {
	switch T := into.(type) {
	case *[]bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*bool:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Bool(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*time.Duration:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Duration(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*float32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Float32(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*float64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Float64(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*int:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*int16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int16(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*int32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int32(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*int64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int64(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*int8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Int8(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*string:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.String(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*time.Time:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Time(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*uint:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*uint16:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint16(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*uint32:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint32(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*uint64:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint64(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	case *[]uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, v)
			return nil
		}
	case *[]*uint8:
		return func(el iter.Pair) error {
			if err := el.Err(); err != nil {
				return err
			}
			v, err := c.Uint8(el.Val())
			if err != nil {
				return err
			}
			*T = append(*T, &v)
			return nil
		}
	}
	return nil
}
