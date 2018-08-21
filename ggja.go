// Package ggja—the Go Generic JSON Api—is a small helper library to make
// working with Go's generic JSON types more comfortable.
package ggja

import (
	"fmt"
	"time"
)

type GenObj = map[string]interface{}
type GenArr = []interface{}

type Obj struct {
	Bare map[string]interface{}
	// When OnError is nil all methods will panic on error, otherwise
	// OnError(err) will be called. Methods returing values will return zero
	// value on error if OnError has no non-local exists.
	OnError func(error)
}

func (o *Obj) Set(key string, v interface{}) *Obj {
	if o.Bare == nil {
		o.fail(fmt.Errorf("cannot set object member, key '%s' not exists", key))
	}
	if _, ok := o.Bare[key]; ok {
		o.Bare[key] = v
	} else {
		o.fail(fmt.Errorf("cannot set object member, key '%s' not exists", key))
	}
	return o
}

func (o *Obj) Sets(s fmt.Stringer, v interface{}) *Obj {
	return o.Set(s.String(), v)
}

func (o *Obj) Put(key string, v interface{}) *Obj {
	if o.Bare == nil {
		o.Bare = make(map[string]interface{})
	}
	o.Bare[key] = v
	return o
}

func (o *Obj) Puts(s fmt.Stringer, v interface{}) *Obj {
	return o.Put(s.String(), v)
}

func (o *Obj) Obj(key string) (sub *Obj) {
	if o == nil || o.Bare == nil {
		return nil
	}
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(map[string]interface{}); ok {
			return &Obj{Bare: res, OnError: o.OnError}
		} else {
			o.fail(fmt.Errorf("object member '%s' is no JSON-object: '%v'", key, tmp))
			return nil
		}
	} else {
		return nil
	}
}

func (o *Obj) Objs(s fmt.Stringer) (sub *Obj) {
	return o.Obj(s.String())
}

func (o *Obj) CObj(key string) (sub *Obj) {
	if sub = o.Obj(key); sub == nil {
		obj := make(map[string]interface{})
		o.Put(key, obj)
		sub = &Obj{Bare: obj, OnError: o.OnError}
	}
	return sub
}

func (o *Obj) CObjs(s fmt.Stringer) (sub *Obj) {
	return o.CObj(s.String())
}

func (o *Obj) MObj(key string) (sub *Obj) {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(map[string]interface{}); ok {
			return &Obj{Bare: res, OnError: o.OnError}
		} else {
			o.fail(fmt.Errorf("object member '%s' is no JSON-object: '%v'", key, tmp))
			return nil
		}
	} else {
		o.fail(fmt.Errorf("no JSON-object object member '%s'", key))
		return nil
	}
}

func (o *Obj) MObjs(s fmt.Stringer) (sub *Obj) {
	return o.MObj(s.String())
}

func (o *Obj) Arr(key string) *Arr {
	if o == nil {
		return nil
	}
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.([]interface{}); ok {
			return &Arr{Bare: res, OnError: o.OnError}
		} else {
			o.fail(fmt.Errorf("object member '%s' is no JSON-array: '%v'", key, tmp))
			return nil
		}
	} else {
		return nil
	}
}

func (o *Obj) Arrs(s fmt.Stringer) *Arr {
	return o.Arr(s.String())
}

func (o *Obj) CArr(key string) (sub *Arr) {
	if sub = o.Arr(key); sub == nil {
		arr := make([]interface{}, 0)
		o.Put(key, arr)
		sub = &Arr{Bare: arr, OnError: o.OnError}
	}
	return sub
}

func (o *Obj) CArrs(s fmt.Stringer) (sub *Arr) {
	return o.CArr(s.String())
}

func (o *Obj) MArr(key string) *Arr {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.([]interface{}); ok {
			return &Arr{Bare: res, OnError: o.OnError}
		} else {
			o.fail(fmt.Errorf("object member '%s' is no JSON-array: '%v'", key, tmp))
			return nil
		}
	} else {
		o.fail(fmt.Errorf("no JSON-array object member '%s'", key))
		return nil
	}
}

func (o *Obj) MArrs(s fmt.Stringer) *Arr {
	return o.MArr(s.String())
}

func (o *Obj) Bool(key string, nvl bool) bool {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(bool); ok {
			return res
		} else {
			o.fail(fmt.Errorf("object member '%s' is not bool: '%v'", key, tmp))
			return nvl
		}
	} else {
		return nvl
	}
}

func (o *Obj) Bools(s fmt.Stringer, nvl bool) bool {
	return o.Bool(s.String(), nvl)
}

func (o *Obj) MBool(key string) bool {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(bool); ok {
			return res
		} else {
			o.fail(fmt.Errorf("object member '%s' is not bool: '%v'", key, tmp))
			return false
		}
	} else {
		o.fail(fmt.Errorf("no boolean object member '%s'", key))
		return false
	}
}

func (o *Obj) MBools(s fmt.Stringer) bool {
	return o.MBool(s.String())
}

func (o *Obj) F64(key string, nvl float64) float64 {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			return res
		} else {
			o.fail(fmt.Errorf("object member '%s' is not float64: '%v'", key, tmp))
			return nvl
		}
	} else {
		return nvl
	}
}

func (o *Obj) F64s(s fmt.Stringer, nvl float64) float64 {
	return o.F64(s.String(), nvl)
}

func (o *Obj) MF64(key string) float64 {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			return res
		} else {
			o.fail(fmt.Errorf("object member '%s' is not float64: '%v'", key, tmp))
			return 0
		}
	} else {
		o.fail(fmt.Errorf("no float64 object member '%s'", key))
		return 0
	}
}

func (o *Obj) MF64s(s fmt.Stringer) float64 {
	return o.MF64(s.String())
}

func (o *Obj) F32(key string, nvl float32) float32 {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			if f32Range(res) {
				return float32(res)
			} else {
				o.fail(fmt.Errorf("object member '%s' out of float32 range: %f",
					key, res))
				return 0
			}
		} else {
			o.fail(fmt.Errorf("object member '%s' is not float32: '%v'", key, tmp))
			return nvl
		}
	} else {
		return nvl
	}
}

func (o *Obj) F32s(s fmt.Stringer, nvl float32) float32 {
	return o.F32(s.String(), nvl)
}

func (o *Obj) MF32(key string) float32 {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			if f32Range(res) {
				return float32(res)
			} else {
				o.fail(fmt.Errorf("object member '%s' out of float32 range: %f",
					key, res))
			}
		} else {
			o.fail(fmt.Errorf("object member '%s' is not float32: '%v'", key, tmp))
		}
	} else {
		o.fail(fmt.Errorf("no float32 object member '%s'", key))
	}
	return 0
}

func (o *Obj) MF32s(s fmt.Stringer) float32 {
	return o.MF32(s.String())
}

func (o *Obj) Int(key string, nvl int) int {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			if intRange(res) {
				return int(res)
			} else {
				o.fail(fmt.Errorf("object member '%s' out of int range: %f",
					key, res))
				return 0
			}
		} else {
			o.fail(fmt.Errorf("object member '%s' is not int: '%v'", key, tmp))
			return nvl
		}
	} else {
		return nvl
	}
}

func (o *Obj) Ints(s fmt.Stringer, nvl int) int {
	return o.Int(s.String(), nvl)
}

func (o *Obj) MInt(key string) int {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			if intRange(res) {
				return int(res)
			} else {
				o.fail(fmt.Errorf("object member '%s' out of int range: %f",
					key, res))
			}
		} else {
			o.fail(fmt.Errorf("object member '%s' is not int: '%v'", key, tmp))
		}
	} else {
		o.fail(fmt.Errorf("no int object member '%s'", key))
	}
	return 0
}

func (o *Obj) MInts(s fmt.Stringer) int {
	return o.MInt(s.String())
}

func (o *Obj) Uint32(key string, nvl uint32) uint32 {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			if uint32Range(res) {
				return uint32(res)
			} else {
				o.fail(fmt.Errorf("object member '%s' out of uint32 range: %f",
					key, res))
				return 0
			}
		} else {
			o.fail(fmt.Errorf("object member '%s' is not uint32: '%v'", key, tmp))
			return nvl
		}
	} else {
		return nvl
	}
}

func (o *Obj) Uint32s(s fmt.Stringer, nvl uint32) uint32 {
	return o.Uint32(s.String(), nvl)
}

func (o *Obj) MUint32(key string) uint32 {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			if uint32Range(res) {
				return uint32(res)
			} else {
				o.fail(fmt.Errorf("object member '%s' out of uint32 range: %f",
					key, res))
			}
		} else {
			o.fail(fmt.Errorf("object member '%s' is not uint32: '%v'", key, tmp))
		}
	} else {
		o.fail(fmt.Errorf("no uint32 object member '%s'", key))
	}
	return 0
}

func (o *Obj) MUint32s(s fmt.Stringer) uint32 {
	return o.MUint32(s.String())
}

func (o *Obj) Int64(key string, nvl int64) int64 {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			if int64Range(res) {
				return int64(res)
			} else {
				o.fail(fmt.Errorf("object member '%s' out of int64 range: %f",
					key, res))
				return 0
			}
		} else {
			o.fail(fmt.Errorf("object member '%s' is not int64: '%v'", key, tmp))
			return nvl
		}
	} else {
		return nvl
	}
}

func (o *Obj) Int64s(s fmt.Stringer, nvl int64) int64 {
	return o.Int64(s.String(), nvl)
}

func (o *Obj) MInt64(key string) int64 {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(float64); ok {
			if int64Range(res) {
				return int64(res)
			} else {
				o.fail(fmt.Errorf("object member '%s' out of int64 range: %f",
					key, res))
			}
		} else {
			o.fail(fmt.Errorf("object member '%s' is not int64: '%v'", key, tmp))
		}
	} else {
		o.fail(fmt.Errorf("no int64 object member '%s'", key))
	}
	return 0
}

func (o *Obj) MInt64s(s fmt.Stringer) int64 {
	return o.MInt64(s.String())
}

func (o *Obj) Str(key, nvl string) string {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(string); ok {
			return res
		} else {
			o.fail(fmt.Errorf("object member '%s' is not string: '%v'", key, tmp))
			return nvl
		}
	} else {
		return nvl
	}
}

func (o *Obj) Strs(s fmt.Stringer, nvl string) string {
	return o.Str(s.String(), nvl)
}

func (o *Obj) MStr(key string) string {
	if tmp, ok := o.Bare[key]; ok {
		if res, ok := tmp.(string); ok {
			return res
		} else {
			o.fail(fmt.Errorf("object member '%s' is not string: '%v'", key, tmp))
			return ""
		}
	} else {
		o.fail(fmt.Errorf("no string object member '%s'", key))
		return ""
	}
}

func (o *Obj) MStrs(s fmt.Stringer) string {
	return o.MStr(s.String())
}

func (o *Obj) Time(key string, nvl time.Time) time.Time {
	tStr := o.Str(key, "")
	if len(tStr) == 0 {
		return nvl
	}
	res, err := time.Parse(time.RFC3339, tStr)
	if err != nil {
		o.fail(err)
		return time.Time{}
	}
	return res
}

func (o *Obj) Times(s fmt.Stringer, nvl time.Time) time.Time {
	return o.Time(s.String(), nvl)
}

func (o *Obj) MTime(key string) time.Time {
	key = o.MStr(key)
	res, err := time.Parse(time.RFC3339, key)
	if err != nil {
		o.fail(err)
		return time.Time{}
	}
	return res
}

func (o *Obj) MTimes(s fmt.Stringer) time.Time {
	return o.MTime(s.String())
}

func (o *Obj) fail(err error) {
	if o.OnError == nil {
		panic(err)
	} else {
		o.OnError(err)
	}
}

type Arr struct {
	Bare []interface{}
	// When OnError is nil all methods will panic on error, otherwise
	// OnError(err) will be called. Methods returing values will return zero
	// value on error if OnError has no non-local exists.
	OnError func(error)
}

func (a *Arr) Set(idx int, v interface{}) *Arr {
	if idx > len(a.Bare) {
		a.fail(fmt.Errorf("cannot set array element, index %d out of range [0; %d)",
			idx,
			len(a.Bare)))
	} else {
		a.Bare[idx] = v
	}
	return a
}

func (a *Arr) Put(idx int, v interface{}) *Arr {
	if idx >= len(a.Bare) {
		if idx >= cap(a.Bare) {
			nb := make([]interface{}, idx+1)
			copy(nb, a.Bare)
			a.Bare = nb
		} else {
			a.Bare = a.Bare[:idx+1]
		}
	}
	a.Bare[idx] = v
	return a
}

func (a *Arr) adjIdx(idx int) int {
	if idx < 0 {
		return len(a.Bare) + idx
	}
	return idx
}

func (a *Arr) Obj(idx int) *Obj {
	if a == nil || idx >= len(a.Bare) {
		return nil
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		return nil
	} else if obj, ok := tmp.(map[string]interface{}); ok {
		return &Obj{Bare: obj, OnError: a.OnError}
	} else {
		a.fail(fmt.Errorf("array element %d is no JSON-object: '%v'", idx, tmp))
		return nil
	}
}

func (a *Arr) CObj(idx int) (elm *Obj) {
	if elm = a.Obj(idx); elm == nil {
		idx = a.adjIdx(idx)
		obj := make(map[string]interface{})
		a.Put(idx, obj)
		elm = &Obj{Bare: obj, OnError: a.OnError}
	}
	return elm
}

func (a *Arr) MObj(idx int) *Obj {
	if idx >= len(a.Bare) {
		a.fail(fmt.Errorf("no JSON-object array element at %d", idx))
		return nil
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		a.fail(fmt.Errorf("no JSON-object array element at %d", idx))
		return nil
	} else if obj, ok := tmp.(map[string]interface{}); ok {
		return &Obj{Bare: obj, OnError: a.OnError}
	} else {
		a.fail(fmt.Errorf("array element %d is no JSON-object: '%v'", idx, tmp))
		return nil
	}
}

func (a *Arr) Arr(idx int) *Arr {
	if a == nil || idx >= len(a.Bare) {
		return nil
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		return nil
	} else if arr, ok := tmp.([]interface{}); ok {
		return &Arr{Bare: arr, OnError: a.OnError}
	} else {
		a.fail(fmt.Errorf("array element %d is no JSON-array: '%v'", idx, tmp))
		return nil
	}
}

func (a *Arr) CArr(idx int) (elm *Arr) {
	if elm = a.Arr(idx); elm == nil {
		idx = a.adjIdx(idx)
		arr := make([]interface{}, 0)
		a.Put(idx, arr)
		elm = &Arr{Bare: arr, OnError: a.OnError}
	}
	return elm
}

func (a *Arr) MArr(idx int) *Arr {
	if idx >= len(a.Bare) {
		a.fail(fmt.Errorf("no JSON-array array element at %d", idx))
		return nil
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		a.fail(fmt.Errorf("no JSON-array array element at %d", idx))
		return nil
	} else if arr, ok := tmp.([]interface{}); ok {
		return &Arr{Bare: arr, OnError: a.OnError}
	} else {
		a.fail(fmt.Errorf("array element %d is no JSON-array: '%v'", idx, tmp))
		return nil
	}
}

func (a *Arr) Bool(idx int, nvl bool) bool {
	if idx >= len(a.Bare) {
		return false
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		return false
	} else if res, ok := tmp.(bool); ok {
		return res
	} else {
		a.fail(fmt.Errorf("array element %d is not boolean: '%v'", idx, tmp))
		return false
	}
}

func (a *Arr) MBool(idx int) bool {
	if idx >= len(a.Bare) {
		a.fail(fmt.Errorf("no boolean array element at %d", idx))
		return false
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		a.fail(fmt.Errorf("no boolean array element at %d", idx))
		return false
	} else if res, ok := tmp.(bool); ok {
		return res
	} else {
		a.fail(fmt.Errorf("array element %d is not boolean: '%v'", idx, tmp))
		return false
	}
}

func (a *Arr) F64(idx int, nvl float64) float64 {
	if idx >= len(a.Bare) {
		return 0
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		return 0
	} else if res, ok := tmp.(float64); ok {
		return res
	} else {
		a.fail(fmt.Errorf("array element %d is not float64: '%v'", idx, tmp))
		return 0
	}
}

func (a *Arr) MF64(idx int) float64 {
	if idx >= len(a.Bare) {
		a.fail(fmt.Errorf("no float64 array element at %d", idx))
		return 0
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		a.fail(fmt.Errorf("no float64 array element at %d", idx))
		return 0
	} else if res, ok := tmp.(float64); ok {
		return res
	} else {
		a.fail(fmt.Errorf("array element %d is not float64: '%v'", idx, tmp))
		return 0
	}
}

func (a *Arr) Int(idx int, nvl int) int {
	if idx >= len(a.Bare) {
		return 0
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		return 0
	} else if res, ok := tmp.(float64); ok {
		if intRange(res) {
			return int(res)
		} else {
			a.fail(fmt.Errorf("array element %d out of int range: %f",
				idx, res))
			return 0
		}
	} else {
		a.fail(fmt.Errorf("array element %d is not int: '%v'", idx, tmp))
		return 0
	}
}

func (a *Arr) MInt(idx int) int {
	if idx >= len(a.Bare) {
		a.fail(fmt.Errorf("no int array element at %d", idx))
		return 0
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		a.fail(fmt.Errorf("no int array element at %d", idx))
		return 0
	} else if res, ok := tmp.(float64); ok {
		if intRange(res) {
			return int(res)
		} else {
			a.fail(fmt.Errorf("array element %d out of int range: %f",
				idx, res))
			return 0
		}
	} else {
		a.fail(fmt.Errorf("array element %d is not int: '%v'", idx, tmp))
		return 0
	}
}

func (a *Arr) Str(idx int, nvl string) string {
	if idx >= len(a.Bare) {
		return ""
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		return ""
	} else if res, ok := tmp.(string); ok {
		return res
	} else {
		a.fail(fmt.Errorf("array element %d is not string: '%v'", idx, tmp))
		return ""
	}
}

func (a *Arr) MStr(idx int) string {
	if idx >= len(a.Bare) {
		a.fail(fmt.Errorf("no string array element at %d", idx))
		return ""
	}
	idx = a.adjIdx(idx)
	if tmp := a.Bare[idx]; tmp == nil {
		a.fail(fmt.Errorf("no string array element at %d", idx))
		return ""
	} else if res, ok := tmp.(string); ok {
		return res
	} else {
		a.fail(fmt.Errorf("array element %d is not string: '%v'", idx, tmp))
		return ""
	}
}

func (a *Arr) fail(err error) {
	if a.OnError == nil {
		panic(err)
	} else {
		a.OnError(err)
	}
}

/*
uint8  : 0 to 255
uint16 : 0 to 65535
uint32 : 0 to 4294967295
uint64 : 0 to 18446744073709551615
int8   : -128 to 127
int16  : -32768 to 32767
int64  : -9223372036854775808 to 9223372036854775807
*/

// TODO int is machine dependent!!!
func intRange(x float64) bool {
	return -2147483648 <= x && x <= 2147483647
}

func uint32Range(x float64) bool {
	return 0 <= x && x <= 4294967295
}

func int64Range(x float64) bool {
	return -9223372036854775808 <= x && x <= 9223372036854775807
}

// TODO
func f32Range(x float64) bool {
	return true
}
