package conv

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestSlice(t *testing.T) {
	var c Conv
	t.Run("Slice", func(t *testing.T) {
		t.Run("Int64", func(t *testing.T) {
			var into []int64
			exp := []int64{12, 345, 6789}

			err := c.Slice(&into, []string{"12", "345", "6789"})
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(exp, into) {
				t.Fatalf("exp (%T) --> %[1]v != %v <-- (%[2]T) got", exp, into)
			}

			into = []int64{}
			err = Slice(&into, []string{"12", "345", "6789"})
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(exp, into) {
				t.Fatalf("exp --> (%T) %#[1]v != %T %#[2]v <-- got", exp, into)
			}

			err = Slice(nil, []string{"12", "345", "6789"})
			if err == nil {
				t.Error("expected non-nil err")
			}
		})
		t.Run("Int64Pointer", func(t *testing.T) {
			var into []*int64
			i1, i2, i3 := new(int64), new(int64), new(int64)
			*i1, *i2, *i3 = 12, 345, 6789
			exp := []*int64{i1, i2, i3}

			err := c.Slice(&into, []string{"12", "345", "6789"})
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(exp, into) {
				t.Fatalf("exp --> (%T) %#[1]v != %T %#[2]v <-- got", exp, into)
			}

			into = []*int64{}
			err = Slice(&into, []string{"12", "345", "6789"})
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(exp, into) {
				t.Fatalf("exp --> (%T) %#[1]v != %T %#[2]v <-- got", exp, into)
			}
		})
	})
}

// Summary:
//
// BenchmarkSlice/<slice size>/<from> to <to>/Conv:
//   Measures the most convenient form of conversion using this library.
//
// BenchmarkSlice/<slice size>/<from> to <to>/Conv:
//   Measures using the library only for the conversion, looping for apending.
//
// BenchmarkSlice/<slice size>/<from> to <to>/Conv:
//   Measures not using this library at all, pure Go implementation.
//
func BenchmarkSlice(b *testing.B) {
	var c Conv

	for _, num := range []int{1024, 64, 16, 4} {
		num := num

		// slow down is really tolerable, only a factor of 1-3 tops
		b.Run(fmt.Sprintf("Length(%d)", num), func(b *testing.B) {

			b.Run("[]string to []int64", func(b *testing.B) {
				strs := make([]string, num)
				for n := 0; n < num; n++ {
					strs[n] = fmt.Sprintf("%v00", n)
				}
				b.ResetTimer()

				b.Run("Conv", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						var into []int64
						err := c.Slice(&into, strs)
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
						var into []int64

						for _, s := range strs {
							v, err := c.Int64(s)
							if err != nil {
								b.Error(err)
							}
							into = append(into, v)
						}
						if len(into) != num {
							b.Error("bad impl")
						}
					}
				})
				b.Run("LoopStdlib", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						var into []int64

						for _, s := range strs {
							v, err := strconv.ParseInt(s, 10, 0)
							if err != nil {
								b.Error(err)
							}
							into = append(into, v)
						}
						if len(into) != num {
							b.Error("bad impl")
						}
					}
				})
			})

			b.Run("[]string to []*int64", func(b *testing.B) {
				strs := make([]string, num)
				for n := 0; n < num; n++ {
					strs[n] = fmt.Sprintf("%v00", n)
				}
				b.ResetTimer()

				b.Run("Library", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						var into []*int64
						err := c.Slice(&into, strs)
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
						into := new([]*int64)

						for _, s := range strs {
							v, err := c.Int64(s)
							if err != nil {
								b.Error(err)
							}
							*into = append(*into, &v)
						}
						if len(*into) != num {
							b.Error("bad impl")
						}
					}
				})
				b.Run("LoopStdlib", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						into := new([]*int64)

						for _, s := range strs {
							v, err := strconv.ParseInt(s, 10, 0)
							if err != nil {
								b.Error(err)
							}
							*into = append(*into, &v)
						}
						if len(*into) != num {
							b.Error("bad impl")
						}
					}
				})
			})
		})
	}
}

// This was just way, way, too slow to be practical.
// func (c Conv) testReflectSlice(into interface{}, from interface{}) error {
// 	intoVal := reflect.ValueOf(into)
// 	if !intoVal.IsValid() {
// 		return fmt.Errorf("%T is not a valid value", into)
// 	}
//
// 	intoKind := intoVal.Kind()
// 	if intoKind != reflect.Ptr {
// 		return fmt.Errorf("%T is not a pointer", into)
// 	}
//
// 	intoVal = intoVal.Elem()
// 	if !intoVal.IsValid() {
// 		return fmt.Errorf("%T is not a valid value", into)
// 	}
//
// 	intoKind = intoVal.Kind()
// 	if intoKind != reflect.Slice {
// 		return fmt.Errorf("%T is not a pointer to a slice", into)
// 	}
//
// 	elemTyp, ok := lutKindFromType(intoVal.Type().Elem())
// 	if !ok {
// 		return fmt.Errorf("%T is not a supported slice type2", into)
// 	}
//
// 	convFn := lutLookup(elemTyp)
// 	if convFn == nil {
// 		return fmt.Errorf("%T is not a supported slice type", into)
// 	}
// 	convFn = lutFuncInt64P
//
// 	var convVals []reflect.Value
// 	err := iter.Walk(from, func(el iter.Pair) error {
// 		// // values
// 		// res, err := convFn(c, el.Val())
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// convVals = append(convVals, reflect.ValueOf(res))
// 		// return nil
//
// 		// // slices of pointers
// 		// 	res, err := convFn(c, el.Val())
// 		// 	if err != nil {
// 		// 		return err
// 		// 	}
// 		//
// 		// 	val := reflect.New(intoVal.Type().Elem().Elem())
// 		// 	refval := reflect.ValueOf(res)
// 		// 	val.Elem().Set(refval)
// 		// 	convVals = append(convVals, val)
// 		// 	return nil
// 		return nil
// 	})
// 	if err != nil {
// 		return err
// 	}
//
// 	intoVal.Set(reflect.Append(intoVal, convVals...))
// 	return nil
// }
