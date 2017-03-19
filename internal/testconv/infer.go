package testconv

import (
	"reflect"
	"testing"
)

func RunInferTests(t *testing.T, fn func(into, from interface{}) error) {
	// // Infer should work as a drop in replacement for Structs, Slices and Maps
	// t.Run("Structs", func(t *testing.T) {
	// 	RunStructTests(t, fn)
	// })
	// t.Run("Slices", func(t *testing.T) {
	// 	RunSliceTests(t, fn)
	// })
	// t.Run("Maps", func(t *testing.T) {
	// 	RunMapTests(t, fn)
	// })

	// Should work with all other assertions
	t.Run("Assertions", func(t *testing.T) {
		for _, a := range assertions {
			for _, e := range a.Exps {

				// Create a pointer from the exp value for inference
				intoVal := reflect.New(reflect.ValueOf(e.Value()).Type())

				// Infer conversion & dereference.
				err := fn(intoVal.Interface(), a.From)
				got := intoVal.Elem().Interface()

				if err = e.Expect(got, err); err != nil {
					t.Fatalf("(FAIL) %v %v\n%v", a, e, err)
				} else {
					t.Logf("(PASS) %v %v", a, e)
				}
			}
		}
	})

	// Touch the negative cases
	t.Run("Negative", func(t *testing.T) {
		type negativeTest struct {
			into, from interface{}
		}
		tests := []negativeTest{
			{nil, nil},                // kind invalid
			{nil, (interface{})(nil)}, // from kind
			{0, nil},                  // non ptr
			{(*complex128)(nil), nil},
			{(*interface{})(nil), nil},
			{(*interface{})(nil), (interface{})(nil)},
			{nil, new(int)},
			{new(int), nil},
		}
		for _, test := range tests {
			if err := fn(test.into, test.from); err == nil {
				t.Fatalf("(FAIL) exp non-nil error for %v -> %v", test.into, test.from)
			}
		}
	})
}
