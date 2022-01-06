package stringtyper

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

type Column string
type ColumnTest struct {
	column        []Column
	kind          reflect.Kind
	maxInt        *int64
	minInt        *int64
	maxUint       *uint64
	minUint       *uint64
	minFloat      *float64
	maxFloat      *float64
	smallestFloat *float64
}

var testCasesCorrectType = []ColumnTest{
	// BOOL
	ColumnTest{
		column: []Column{"1", "1", "0"},
		kind:   reflect.Bool,
	},
	ColumnTest{
		column: []Column{"true", "true", "false"},
		kind:   reflect.Bool,
	},
	ColumnTest{
		column: []Column{"true", "true", "false", "1", "0"},
		kind:   reflect.Bool,
	},
	ColumnTest{
		column: []Column{"true", "false", "TRUE", "FALSE", "True", "False", "1", "0", "f", "t", "T", "F"},
		kind:   reflect.Bool,
	},

	// INTs

	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "-128"},
		kind:   reflect.Int8,
		minInt: Int64(-128),
		maxInt: Int64(2),
	},

	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "127"},
		kind:   reflect.Int8,
		minInt: Int64(-1),
		maxInt: Int64(127),
	},

	ColumnTest{
		column:  []Column{"1", "1", "0", "2", "255"},
		kind:    reflect.Uint8,
		minUint: Uint64(0),
		maxUint: Uint64(255),
	},

	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "32767"},
		kind:   reflect.Int16,
		minInt: Int64(-1),
		maxInt: Int64(32767),
	},

	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "-32768"},
		kind:   reflect.Int16,
		minInt: Int64(-32768),
		maxInt: Int64(2),
	},

	ColumnTest{
		column:  []Column{"1", "1", "0", "2", "65535"},
		kind:    reflect.Uint16,
		minUint: Uint64(0),
		maxUint: Uint64(300),
	},

	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "300", "2147483647"},
		kind:   reflect.Int32,
		minInt: Int64(-1),
		maxInt: Int64(2147483647),
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "300", "-2147483648"},
		kind:   reflect.Int32,
		minInt: Int64(-2147483648),
		maxInt: Int64(300),
	},
	ColumnTest{
		column:  []Column{"1", "1", "0", "2", "300", "65536"},
		kind:    reflect.Uint32,
		minUint: Uint64(0),
		maxUint: Uint64(65536),
	},

	ColumnTest{
		column: []Column{"-1", "1", "0", "2", "300", "9223372036854775807"},
		kind:   reflect.Int64,
		minInt: Int64(-1),
		maxInt: Int64(9223372036854775807),
	},

	ColumnTest{
		column: []Column{"-1", "1", "0", "2", "300", "-9223372036854775808"},
		kind:   reflect.Int64,
		minInt: Int64(-9223372036854775808),
		maxInt: Int64(300),
	},

	ColumnTest{
		column:  []Column{"1", "0", "2", "300", "18446744073709551615"},
		kind:    reflect.Uint64,
		minUint: Uint64(0),
		maxUint: Uint64(18446744073709551615),
	},

	// FLOATS
	// float32
	ColumnTest{
		column: []Column{"-1", "1", "1", "0.", "2.0", ".98"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0.", "2.0", "-.98"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2.", ".98"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", ".98"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"-.1", "1", "1", "0", ".98"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", ".00000000000000000000000000001"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "-.00000000000000000000000000001"},
		kind:   reflect.Float32,
	},

	// math.MaxFloat32
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "3.40282346638528859811704183484516925440e+38"},
		kind:   reflect.Float32,
	},

	// -math.MaxFloat32
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "-3.40282346638528859811704183484516925440e+38"},
		kind:   reflect.Float32,
	},

	// math.SmallestNonzeroFloat32
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "1.401298464324817070923729583289916131280e-45"},
		kind:   reflect.Float32,
	},
	// -math.SmallestNonzeroFloat32
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "-1.401298464324817070923729583289916131280e-45"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "1e32"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0.0", "2", "0.98"},
		kind:   reflect.Float32,
	},

	// float64
	ColumnTest{
		column: []Column{"-1", "1", "1", "0.0", "99999999999999999999999999999999999999999999999999999999999999999999999"},
		kind:   reflect.Float64,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0.0", "-99999999999999999999999999999999999999999999999999999999999999999999999"},
		kind:   reflect.Float64,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "1e44"},
		kind:   reflect.Float64,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "1e308"},
		kind:   reflect.Float64,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "-1e308"},
		kind:   reflect.Float64,
	},

	// math.MaxFloat64
	ColumnTest{
		column: []Column{"1", "1", "0", "1.79769313486231570814527423731704356798070e+308"},
		kind:   reflect.Float64,
	},
	// -math.MaxFloat64
	ColumnTest{
		column: []Column{"1", "1", "0", "-1.79769313486231570814527423731704356798070e+308"},
		kind:   reflect.Float64,
	},

	// math.SmallestNonzeroFloat64
	ColumnTest{
		column: []Column{"1", "1", "0", "4.9406564584124654417656879286822137236505980e-324"},
		kind:   reflect.Float64,
	},

	// -math.SmallestNonzeroFloat64
	ColumnTest{
		column: []Column{"1", "1", "0", "-4.9406564584124654417656879286822137236505980e-324"},
		kind:   reflect.Float64,
	},

	// STRINGS
	ColumnTest{
		column: []Column{"a", "b", "10"},
		kind:   reflect.String,
	},

	ColumnTest{
		column: []Column{"a", "b", "c"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"a", "b", "c", "1"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"a", "b", "c", "1.5"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"a", "b", "c", "1e5"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"a", "b", "c", "true"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"1a", "2b", "4c"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"1a", "2b", "4c"},
		kind:   reflect.String,
	},

	// Using number ranges from math
	// INTs
	// 8
	ColumnTest{
		column: []Column{Column(strconv.FormatInt(int64(math.MinInt8), 10)), "0", Column(strconv.FormatInt(int64(math.MaxInt8), 10))},
		kind:   reflect.Int8,
	},
	ColumnTest{
		column: []Column{"0", Column(strconv.FormatUint(uint64(math.MaxUint8), 10))},
		kind:   reflect.Uint8,
	},
	// 16
	ColumnTest{
		column: []Column{Column(strconv.FormatInt(int64(math.MinInt16), 10)), "0", Column(strconv.FormatInt(int64(math.MaxInt16), 10))},
		kind:   reflect.Int16,
	},
	ColumnTest{
		column: []Column{"0", Column(strconv.FormatUint(uint64(math.MaxUint16), 10))},
		kind:   reflect.Uint16,
	},
	// 32
	ColumnTest{
		column: []Column{Column(strconv.FormatInt(int64(math.MinInt32), 10)), "0", Column(strconv.FormatInt(int64(math.MaxInt32), 10))},
		kind:   reflect.Int32,
	},
	ColumnTest{
		column: []Column{"0", Column(strconv.FormatUint(uint64(math.MaxUint32), 10))},
		kind:   reflect.Uint32,
	},
	// 64
	ColumnTest{
		column: []Column{Column(strconv.FormatInt(int64(math.MinInt64), 10)), "0", Column(strconv.FormatInt(int64(math.MaxInt64), 10))},
		kind:   reflect.Int64,
	},
	ColumnTest{
		column: []Column{"0", Column(strconv.FormatUint(uint64(math.MaxUint64), 10))},
		kind:   reflect.Uint64,
	},

	// Float
	// 32
	// ColumnTest{
	// 	column: []Column{"0",
	// 		// b,e,E,f,g,G,x,X from: https://pkg.go.dev/strconv#FormatFloat
	// 		//Column(strconv.FormatFloat(math.SmallestNonzeroFloat32, 'b', 2, 32)),
	// 		//Column(strconv.FormatFloat(math.MaxFloat32, 'b', -1, 32)),
	// 		Column(strconv.FormatFloat(math.SmallestNonzeroFloat32, 'e', -1, 32)),
	// 		Column(strconv.FormatFloat(math.MaxFloat32, 'e', -1, 32)),
	// 		Column(strconv.FormatFloat(math.SmallestNonzeroFloat32, 'E', -1, 32)),
	// 		Column(strconv.FormatFloat(math.MaxFloat32, 'E', -1, 32)),
	// 	},
	// 	kind: reflect.Float32,
	// },
}

var testCasesIncorrectType = []ColumnTest{
	// BOOL
	ColumnTest{
		column: []Column{"1", "1", "0", "mmm"},
		kind:   reflect.Bool,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "2"},
		kind:   reflect.Bool,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "0.2"},
		kind:   reflect.Bool,
	},
	ColumnTest{
		column: []Column{"true", "true", "false", "2"},
		kind:   reflect.Bool,
	},
	ColumnTest{
		column: []Column{"true", "true", "false", "1", "0", "m"},
		kind:   reflect.Bool,
	},
	ColumnTest{
		column: []Column{"true", "false", "TRUE", "FALSE", "True", "False", "1", "0", "f", "t", "T", "F", "x"},
		kind:   reflect.Bool,
	},

	// INTs

	ColumnTest{
		column: []Column{"-127", "127", "1", "0", "2", "m"},
		kind:   reflect.Int8,
	},

	ColumnTest{
		column: []Column{"1", "1", "0", "2", "m"},
		kind:   reflect.Uint8,
	},

	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "300", "-2", "32768"},
		kind:   reflect.Int16,
	},

	ColumnTest{
		column: []Column{"1", "1", "0", "2", "300", "-1"},
		kind:   reflect.Uint16,
	},

	ColumnTest{
		column: []Column{"1", "1", "0", "2", "300", "65536"},
		kind:   reflect.Uint16,
	},

	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "300", "2147483648"},
		kind:   reflect.Int32,
	},

	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2", "300", "-2147483649"},
		kind:   reflect.Int32,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "2", "300", "4294967296"},
		kind:   reflect.Uint32,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "2", "300", "m"},
		kind:   reflect.Uint32,
	},

	ColumnTest{
		column: []Column{"-1", "1", "0", "2", "300", "9223372036854775808"},
		kind:   reflect.Int64,
	},

	ColumnTest{
		column: []Column{"-1", "1", "0", "2", "300", "-9223372036854775809"},
		kind:   reflect.Int64,
	},
	ColumnTest{
		column: []Column{"-1", "1", "0", "2", "300", "m"},
		kind:   reflect.Int64,
	},

	ColumnTest{
		column: []Column{"-1", "1", "0", "2", "300", ".2"},
		kind:   reflect.Int64,
	},

	ColumnTest{
		column: []Column{"-1", "0", "2", "300", "5000000000"},
		kind:   reflect.Uint64,
	},
	ColumnTest{
		column: []Column{"1", "0", "2", "300", "500000000000000000000"},
		kind:   reflect.Uint64,
	},

	// FLOATS
	ColumnTest{
		column: []Column{"-1", "1", "1", "0.", "2.0", "1e309"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0", "2.", "m"},
		kind:   reflect.Float32,
	},

	ColumnTest{
		column: []Column{"-1", "1", "1", "0.0", "2"},
		kind:   reflect.Float64,
	},
	ColumnTest{
		column: []Column{"-1", "1", "1", "0.0", "99999999999999999999999999999999999999999999999999999999999999999999999"},
		kind:   reflect.Float32,
	},

	ColumnTest{
		column: []Column{"1", "1", "0", "1e39"},
		kind:   reflect.Float32,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "1e309"},
		kind:   reflect.Float64,
	},

	ColumnTest{
		column: []Column{"1", "1", "0", "-1e310"},
		kind:   reflect.Float64,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "1e-1024"},
		kind:   reflect.Float64,
	},
	ColumnTest{
		column: []Column{"1", "1", "0", "-1e-1024"},
		kind:   reflect.Float64,
	},

	// STRINGS
	ColumnTest{
		column: []Column{"t", "f", "t"},
		kind:   reflect.String,
	},

	ColumnTest{
		column: []Column{"1", "4", "5"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{".1", "1", "5", "7"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"9", "2", "11", "-1.5"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"42", "1e99"},
		kind:   reflect.String,
	},
	ColumnTest{
		column: []Column{"256"},
		kind:   reflect.String,
	},
}

func TestCasesCorrectType(t *testing.T) {
	for _, test := range testCasesCorrectType {
		ti := NewStringTyper()

		for _, s := range test.column {
			ti.CheckFieldTypeAndLength(string(s))
		}

		if k := ti.Kind(); k != test.kind {
			t.Error(test.column, test.kind, k, ti.errFloat64)
		}
	}
}

func TestCasesCorrectIntRanges(t *testing.T) {
	for _, test := range testCasesCorrectType {
		ti := NewStringTyper()

		for _, s := range test.column {
			ti.CheckFieldTypeAndLength(string(s))
		}

		// check for ranges
		if ti.MaxInt != nil && ti.alwaysInt() && test.maxInt != nil {
			if *ti.MaxInt != *test.maxInt {
				t.Fatalf("ti.Maxint=%d != test.maxint=%d  \n  test=%+v\n  test.column=%+v\n  ti=%+v", *ti.MaxInt, *test.maxInt, test, test.column, ti)
			}
		}
	}
}

var FORMATS = []byte{ /*'b',*/ 'e', 'E', 'f', 'g', 'G', 'x', 'X'}

func TestCasesFloat32Formats(t *testing.T) {
	values := []float64{math.SmallestNonzeroFloat32, math.MaxFloat32, -math.SmallestNonzeroFloat32, -math.MaxFloat32}

	for _, format := range FORMATS {
		for _, value := range values {
			ti := NewStringTyper()
			// Convert to f64 so small values are not zeroed
			col := Column(strconv.FormatFloat(value, format, -1, 64))
			ti.CheckFieldTypeAndLength(string(col))
			if k := ti.Kind(); k != reflect.Float32 {
				t.Log(fmt.Errorf("value=%e format=%s  string=%s expected=%s detected=%s typeInfo=%+v", value, string(format), col, reflect.Float32, k, ti))
				t.Fatal()
			}
		}
	}
}

func TestCasesFloat64Formats(t *testing.T) {
	values := []float64{math.SmallestNonzeroFloat64, math.MaxFloat64, -math.SmallestNonzeroFloat64, -math.MaxFloat64}

	for _, format := range FORMATS {
		for _, value := range values {
			ti := NewStringTyper()
			col := Column(strconv.FormatFloat(value, format, -1, 64))
			ti.CheckFieldTypeAndLength(string(col))
			if k := ti.Kind(); k != reflect.Float64 {
				t.Log(fmt.Errorf("value=%e format=%s  string=%s expected=%s detected=%s typeInfo=%+v", value, string(format), col, reflect.Float64, k, ti))
				t.Fatal()
			}
		}
	}
}

func TestCasesIncorrectType(t *testing.T) {
	for _, test := range testCasesIncorrectType {
		ti := NewStringTyper()

		for _, s := range test.column {
			ti.CheckFieldTypeAndLength(string(s))
		}

		if k := ti.Kind(); k == test.kind {
			t.Error(test.column, test.kind, k)
		}
	}
}

func Int64(i int64) *int64 {
	return &i
}

func Uint64(i uint64) *uint64 {
	return &i
}

func Float64(f float64) *float64 {
	return &f
}
