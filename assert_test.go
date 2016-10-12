package conv

import (
	"bytes"
	"fmt"
	"math"
	"path"
	"runtime"
	"strings"
	"time"
)

// assertions is a slice of every assertion in this library.
var (
	summary    string
	assertions Assertions
	sections   []Section
)

// asserter is the interface this package uses to validate a Converter properly
// implements the rules declared in this package. It should return an error for
// each invalid conversion.
type Asserter interface {
	Assert(c Converter) []error
}

// Assertions is slice of assertion.
type Assertions []Assertion

// asserts takes a list of interfaces and associates tags to them before
// appending to our global assertions list.
func asserts(args ...interface{}) Assertions {
	var (
		tags []string
		d    string
		s    Assertions
	)
	for _, arg := range args {
		switch T := arg.(type) {
		case []string:
			for _, s := range T {
				if len(s) > 0 {
					tags = append(tags, s)
				}
			}
		case string:
			if len(T) > 0 {
				tags = append(tags, T)
			}
		case Assertions:
			s = append(s, T...)
		case Assertion:
			s = append(s, T)
		}
	}
	for i := range s {
		s[i].Tags = append(s[i].Tags, tags...)
		if len(d) > 0 {
			s[i].Desc = d
		}
		assertions = append(assertions, s[i])
	}
	return s
}

// Tags will return all assertions that contain all the given tags.
func (a Assertions) Tags(tags ...string) Assertions {
	var out Assertions
	for i := range tags {
		tags[i] = strings.ToLower(tags[i])
	}
	for _, assertion := range a {
		if containsAll(tags, assertion.Tags) {
			out = append(out, assertion)
		}
	}
	return out
}

// Assertion represents a set of expected values from a Converter. It's the
// general implementation of asserter. It will compare each Expects to the
// associated Converter.T(). It returns an error for each failed conversion
// or an empty error slice if no failures occurred.
type Assertion struct {

	// Interface is the underlying value to be used in conversions.
	Interface interface{}

	// Name of the converter function / interface for T of interface.
	Name string

	// Code to construct Interface
	Code string

	// Type of Interface field using fmt %#T.
	Type string

	// File is the file the assertion is defined in.
	File string

	// Line is the line number the assertion is defined in.
	Line int

	// Description for this assertion.
	Desc string

	// Tags is a slice of strings to mark assertions for lookup later.
	Tags []string

	// Expects contains a list of values this Assertion expects to see as a result
	// of calling the conversion function for the associated type. For example
	// appending a single int64 to this slice would mean it is expected that:
	//   conv.Int64(assertion.Interface) == Expect[0].(int64)
	// When appending multiple values of the same type, the last one appended is
	// used.
	Expects []interface{}
}

// assert will create an assertion type for the first argument given. Args is
// consumed into the Expects slice, with slices of interfaces being flattened.
func assert(value interface{}, args ...interface{}) Assertion {
	a := Assertion{
		Interface: value,
		Type:      typename(value),
		Tags:      []string{},
	}
	split := strings.SplitN(a.Type, ".", 2)
	a.Name = strings.Title(split[len(split)-1])
	a.Code = fmt.Sprintf("%#v", value)
	_, a.File, a.Line, _ = runtime.Caller(1)
	a.File = path.Base(a.File)

	var slurp func(value interface{})
	slurp = func(value interface{}) {
		switch T := value.(type) {
		case []interface{}:
			for _, t := range T {
				slurp(t)
			}
		case []Asserter:
			for _, t := range T {
				slurp(t)
			}
		default:
			a.Expects = append(a.Expects, value)
		}
	}
	for _, arg := range args {
		slurp(arg)
	}
	return a
}

// Assert implements the asserter interface.
func (a Assertion) Assert(c Converter) (out []error) {
	for _, expect := range a.Expects {
		if errs := a.expect(c, expect); len(errs) > 0 {
			out = append(out, errs...)
		}
	}
	return
}

// Make sure the converter as well as library functions return an expected value.
func (a Assertion) expect(c Converter, want interface{}) (out []error) {
	err := func(want, got interface{}) {
		errStr := "\n%v:%d assert(%[3]T(%#[3]v), ...%T(%#[4]v))\n"
		errStr += "  want %#[4]v\n  got: %v"
		out = append(out, fmt.Errorf(
			errStr, a.File, a.Line, a.Interface, want, got))
	}
	chk := func(want, convGot interface{}, convEq bool, funcGot interface{}, funcEq bool) {
		if !convEq {
			err(want, convGot)
		}
		if !funcEq {
			err(want, funcGot)
		}
	}

	switch T := Indirect(want).(type) {
	case Asserter:
		if asserterErrs := T.Assert(c); len(asserterErrs) > 0 {
			for _, asserterErr := range asserterErrs {
				out = append(out, fmt.Errorf(
					"\n%v:%d assert(%[3]T(%#[3]v), ...%T(%#[4]v))\n  Assertion error:\n    %s",
					a.File, a.Line, a.Interface, T, asserterErr))
			}
		}
	case time.Time:
		if got := c.Time(); !got.Equal(T) {
			err(T, got)
		}
	case time.Duration:
		convGot, funcGot := c.Duration(), Duration(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case string:
		convGot, funcGot := c.String(), String(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case bool:
		convGot, funcGot := c.Bool(), Bool(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case complex64:
		convGot, funcGot := c.Complex64(), Complex64(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case complex128:
		convGot, funcGot := c.Complex128(), Complex128(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case int:
		convGot, funcGot := c.Int(), Int(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case int8:
		convGot, funcGot := c.Int8(), Int8(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case int16:
		convGot, funcGot := c.Int16(), Int16(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case int32:
		convGot, funcGot := c.Int32(), Int32(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case int64:
		convGot, funcGot := c.Int64(), Int64(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case float32:
		convGot, funcGot := c.Float32(), Float32(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case float64:
		convGot, funcGot := c.Float64(), Float64(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case uint:
		convGot, funcGot := c.Uint(), Uint(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case uint8:
		convGot, funcGot := c.Uint8(), Uint8(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case uint16:
		convGot, funcGot := c.Uint16(), Uint16(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case uint32:
		convGot, funcGot := c.Uint32(), Uint32(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	case uint64:
		convGot, funcGot := c.Uint64(), Uint64(a.Interface)
		chk(T, convGot, convGot == T, funcGot, funcGot == T)
	default:
		out = append(out, fmt.Errorf(`Conversion of type %T is not supported`, want))
	}
	return
}

func summarize(s string) {
	s = stripws(s)
	summary = fmt.Sprintf("%s\n%s", summary, s)
}

// Section defines a doc section, such as "Numeric Conversions" from doc.go.
type Section struct {
	Name   string
	Groups []Group
}

func section(name string, groups ...Group) {
	d := Section{Name: name}
	for _, group := range groups {
		for i := range group.Assertions {
			group.Assertions[i].Tags = append(group.Assertions[i].Tags, name)
		}
		d.Groups = append(d.Groups, group)
		assertions = append(assertions, group.Assertions...)
	}
	sections = append(sections, d)
}

// Section defines a group within a section which has a small sentence describing
// it as well as a list of assertions to assert the documentation is valid.
type Group struct {
	Description string
	Assertions  Assertions
}

func group(desc string, a ...Assertion) Group {
	desc = stripws(desc)
	d := Group{Description: desc}
	for _, assertion := range a {
		d.Assertions = append(d.Assertions, assertion)
	}
	return d
}

// Assert implements the asserter interface by calling the Asserter on each
// member of this slice.
func (a Assertions) Assert(c Converter) []error {
	var out []error
	for _, assertion := range a {
		errs := assertion.Assert(c)
		if len(errs) > 0 {
			out = append(out, errs...)
		}
	}
	return out
}

// FloatAssertion asserts that (converter.Float64() - Expect) < epsilon64
type Float64Assertion struct {
	Expect float64
}

const (
	epsilon64 = float64(.00000000000000001)
)

func (a Float64Assertion) Assert(converter Converter) []error {
	abs := math.Abs(converter.Float64() - a.Expect)
	if abs < epsilon64 {
		return nil
	}
	return []error{fmt.Errorf("Float64Assertion%v.assert(%#v): abs value %v exceeded epsilon %v",
		a, converter, abs, epsilon64),
	}
}

// Float32Assertion asserts that (converter.Float32() - Expect) < epsilon64
type Float32Assertion struct {
	Expect float32
}

func (a Float32Assertion) Assert(converter Converter) []error {
	abs := math.Abs(float64(converter.Float32() - a.Expect))
	if abs < epsilon64 {
		return nil
	}
	return []error{fmt.Errorf("Float32Assertion%v.assert(%#v): abs value %v exceeded epsilon %v",
		a, converter, abs, epsilon64),
	}
}

// DocAssertion is always true, allows annotation without equality. Useful for
// string representation of time.Now() with generation from templates.
type DocAssertion string

func (a DocAssertion) Assert(converter Converter) []error { return nil }

// TimeAssertion helps validate time.Time() conversions, specifically because
// under some conversions time.Now() may be used.
type TimeAssertion struct {
	Moment  time.Time
	Offset  time.Duration
	Epsilon time.Duration
}

// Assert implements the asserter interface for timeAssertion. It will check
// that the difference between the Moment time plus a.Offset is within a.Epsilon
// when subtracted by the converter.Time().
func (a TimeAssertion) Assert(converter Converter) []error {
	epsilon := a.Epsilon
	offset := a.Moment.Add(a.Offset)
	diff := converter.Time().Sub(offset)

	if epsilon == 0 {
		epsilon = time.Millisecond
	}
	if diff < 0 {
		diff = -diff
	}
	if diff > epsilon {
		return []error{
			fmt.Errorf("TimeAssertion%v.assert(%#v): time diff %v exceeded epsilon %v",
				a, converter, diff, epsilon),
		}
	}
	return nil
}

// NowAssertion is like TimeAssertion but makes `Moment` time.Now
type NowAssertion struct {
	Offset  time.Duration
	Epsilon time.Duration
}

func (a NowAssertion) Assert(c Converter) []error {
	assertion := TimeAssertion{
		Moment:  time.Now(),
		Offset:  a.Offset,
		Epsilon: a.Epsilon,
	}
	return assertion.Assert(c)
}

// bounds returns 8-64 bit signed & unsigned boundary edges
func bounds() (ints []int64, uints []uint64) {
	for i := uint64(4); i < 64; {
		i *= 2
		mins := 1<<(i-1) - 1
		minu := 1<<i - 1
		ints = append(ints, []int64{
			int64(mins - 1), int64(mins), int64(mins + 1),
			int64(minu - 1), int64(minu), int64(minu + 1),
		}...)
		uints = append(uints, []uint64{
			uint64(minu - 1), uint64(minu), uint64(minu + 1),
		}...)
	}
	return
}

func typename(value interface{}) (name string) {
	parts := strings.SplitN(fmt.Sprintf("%T", value), ".", 2)

	if len(parts) == 2 {
		name = parts[0] + "." + parts[1]
	} else {
		name = parts[0]
	}
	return
}

func durations(d time.Duration) []interface{} {
	return []interface{}{
		NowAssertion{Offset: d},
		time.Duration(d),
	}
}

func numerics(from int64) (out []interface{}) {
	out = append(out, ints(int64(from))...)
	out = append(out, uints(uint64(from))...)
	out = append(out, floats(float64(from))...)
	return
}

func ints(from int64) []interface{} {
	return []interface{}{int(from), int8(from), int16(from), int32(from), from}
}

func uints(from uint64) []interface{} {
	return []interface{}{uint(from), uint8(from), uint16(from), uint32(from), from}
}

func floats(from float64) []interface{} {
	return []interface{}{Float32Assertion{float32(from)}, Float64Assertion{float64(from)}}
}

func contains(search string, items []string) bool {
	for _, item := range items {
		if strings.ToLower(search) == strings.ToLower(item) {
			return true
		}
	}
	return false
}

func containsAll(searches, items []string) bool {
	for _, search := range searches {
		if !contains(search, items) {
			return false
		}
	}
	return true
}

func stripws(s string) string {
	var b bytes.Buffer
	white := func(r byte) bool {
		return r == 0x9 || r == 0xA || r == 0x20
	}
	for i := 1; i < len(s); i++ {
		lb := s[i-1]

		if white(lb) && white(s[i]) {
			b.WriteByte(' ')
			for i < len(s)-1 {
				i++
				if !white(s[i]) {
					i--
					break
				}
			}
		} else {
			b.WriteByte(lb)
		}
	}
	b.WriteString(s[len(s)-1:])
	return b.String()
}
