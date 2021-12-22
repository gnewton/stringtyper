package feildur

import (
	"errors"
	"fmt"
	//"log"
	"math"
	"reflect"
	"strconv"
)

type StringTyper struct {
	MaxInt        *int64
	MinInt        *int64
	MaxUint       *uint64
	MinUint       *uint64
	MinFloat      *float64
	MaxFloat      *float64
	SmallestFloat *float64
	alwaysBool    bool
	alwaysFloat32 bool
	alwaysFloat64 bool
	alwaysInt08   bool
	alwaysInt16   bool
	alwaysInt32   bool
	alwaysInt64   bool
	alwaysUint08  bool
	alwaysUint16  bool
	alwaysUint32  bool
	alwaysUint64  bool
	maxLength     int
	errFloat64    error
}

func NewStringTyper() *StringTyper {
	ti := StringTyper{
		alwaysBool:    true,
		alwaysFloat32: true,
		alwaysFloat64: true,
		alwaysInt08:   true,
		alwaysInt16:   true,
		alwaysInt32:   true,
		alwaysInt64:   true,
		alwaysUint08:  true,
		alwaysUint16:  true,
		alwaysUint32:  true,
		alwaysUint64:  true,
	}
	return &ti
}

func (ti *StringTyper) CheckFieldTypeAndLength(v string) {
	l := len(v)
	if ti.maxLength < l {
		ti.maxLength = l
	}

	if _, err := strconv.ParseBool(v); err != nil {
		ti.alwaysBool = false
	}

	// If the string when converted to a float64 is smaller than the smallest non zero float32, then it should be a float64.
	// NB: math.SmallestNonzeroFloat32 is a float64
	//
	v64, err := strconv.ParseFloat(v, 64)
	if err == nil {
		v64 = math.Abs(v64)
		if v64 > 0.0 && v64 < math.SmallestNonzeroFloat32 {
			ti.alwaysFloat32 = false
		}
		ti.checkFloat(v64)
	}

	if _, err := strconv.ParseFloat(v, 32); err != nil {
		ti.alwaysFloat32 = false

	}

	if _, err := strconv.ParseFloat(v, 64); err != nil {
		ti.alwaysFloat64 = false
		ti.MinFloat = nil
		ti.MaxFloat = nil
		ti.SmallestFloat = nil
		ti.errFloat64 = err
	}

	var i int64
	var ui uint64

	if ui, err = strconv.ParseUint(v, 10, 8); err != nil {
		ti.alwaysUint08 = false
	} else {
		ti.checkUint(ui)
	}

	if ui, err = strconv.ParseUint(v, 10, 16); err != nil {
		ti.alwaysUint16 = false
	} else {
		ti.checkUint(ui)
	}

	if ui, err = strconv.ParseUint(v, 10, 32); err != nil {
		ti.alwaysUint32 = false
	} else {
		ti.checkUint(ui)
	}

	if ui, err = strconv.ParseUint(v, 10, 64); err != nil {
		ti.alwaysUint64 = false
	} else {
		ti.checkUint(ui)
	}

	if i, err = strconv.ParseInt(v, 10, 8); err != nil {
		ti.alwaysInt08 = false
	} else {
		ti.checkInt(i)
	}

	if i, err = strconv.ParseInt(v, 10, 16); err != nil {
		ti.alwaysInt16 = false
	} else {
		ti.checkInt(i)
	}

	if i, err = strconv.ParseInt(v, 10, 32); err != nil {
		ti.alwaysInt32 = false
	} else {
		ti.checkInt(i)
	}

	if i, err = strconv.ParseInt(v, 10, 64); err != nil {
		ti.alwaysInt64 = false
	} else {
		ti.checkInt(i)
	}

}

func (ti *StringTyper) checkUint(i uint64) {
	if ti.MinUint == nil {
		ti.MinUint = new(uint64)
		*ti.MinUint = i
	} else {
		if i < *ti.MinUint {
			*ti.MinUint = i
		}
	}
	if ti.MaxUint == nil {
		ti.MaxUint = new(uint64)
		*ti.MaxUint = i
	} else {
		if i > *ti.MaxUint {
			*ti.MaxUint = i
		}
	}
}

func (ti *StringTyper) alwaysInt() bool {
	return ti.alwaysInt08 || ti.alwaysInt16 || ti.alwaysInt32 || ti.alwaysInt64
}

func (ti *StringTyper) checkInt(i int64) {
	//log.Println("$ ", ti.MaxInt, i)
	if ti.MinInt == nil {
		ti.MinInt = &i
	} else {
		if i < *ti.MinInt {
			ti.MinInt = &i
		}
	}
	if ti.MaxInt == nil {
		ti.MaxInt = &i
	} else {
		if i > *ti.MaxInt {
			ti.MaxInt = &i
		}
	}
	//log.Println(*ti.MaxInt)
}

func (ti *StringTyper) checkFloat(v float64) {
	if ti.MinFloat == nil {
		ti.MinFloat = new(float64)
		*ti.MinFloat = v
	} else {
		if v < *ti.MinFloat {
			*ti.MinFloat = v
		}
	}

	if ti.MaxFloat == nil {
		ti.MaxFloat = new(float64)
		*ti.MaxFloat = v
	} else {
		if v > *ti.MaxFloat {
			*ti.MaxFloat = v
		}
	}

	v = math.Abs(v)
	if ti.SmallestFloat == nil {
		ti.SmallestFloat = new(float64)
		*ti.SmallestFloat = v
	} else {
		if v > 0 && v < *ti.SmallestFloat {
			*ti.SmallestFloat = v
		}
	}

}

func (ti *StringTyper) Kind() reflect.Kind {
	if ti.alwaysBool {
		return reflect.Bool
	}

	if ti.alwaysUint08 {
		return reflect.Uint8
	}

	if ti.alwaysUint16 {
		return reflect.Uint16
	}

	if ti.alwaysUint32 {
		return reflect.Uint32
	}

	if ti.alwaysUint64 {
		return reflect.Uint64
	}

	if ti.alwaysInt08 {
		return reflect.Int8
	}

	if ti.alwaysInt16 {
		return reflect.Int16
	}

	if ti.alwaysInt32 {
		return reflect.Int32
	}

	if ti.alwaysInt64 {
		return reflect.Int64
	}

	if ti.alwaysFloat32 {
		return reflect.Float32
	}

	if ti.alwaysFloat64 {
		return reflect.Float64
	}
	return reflect.String
}

type StringTypers []*StringTyper

func NewStringTypers(n int) (StringTypers, error) {
	if n <= 0 {
		return nil, errors.New("n<=1")
	}
	typeInfos := make(StringTypers, n)

	for i := 0; i < n; i++ {
		typeInfos[i] = NewStringTyper()
	}

	return typeInfos, nil
}

func (tim StringTypers) Kinds() []reflect.Kind {
	kinds := make([]reflect.Kind, len(tim))
	for i, _ := range kinds {
		kinds[i] = tim[i].Kind()
	}
	return kinds
}

func (tim StringTypers) CheckFieldTypeAndLength(vs []string) error {
	if len(vs) != len(tim) {
		return fmt.Errorf("String array size=%d does not match existing StringTypers size=%d", len(vs), len(tim))
	}

	for i := 0; i < len(tim); i++ {
		tim[i].CheckFieldTypeAndLength(vs[i])
	}
	return nil
}
