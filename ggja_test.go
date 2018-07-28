package ggja

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type key int

const (
	kFoo key = iota
	kBar
	kBaz
)

func (k key) String() string {
	switch k {
	case kFoo:
		return "foo"
	case kBar:
		return "bar"
	case kBaz:
		return "baz"
	default:
		return ""
	}
}

func fail(err error) {
	fmt.Println("ERROR:", err)
}

func TestObj_CallNPut1(t *testing.T) {
	var set = func(o *Obj) {
		o.Put("key", "val")
	}
	var obj = Obj{OnError: fail}
	set(&obj)
	if obj.MStr("key") != "val" {
		t.Fatal("failed to set key to val")
	}
}

func TestObj_CallNPut2(t *testing.T) {
	var set = func(o Obj) {
		o.Set("key1", "val1")
	}
	var obj = Obj{OnError: fail}
	obj.Put("key1", "-")
	obj.Put("key2", "-")
	set(obj)
	if obj.MStr("key1") != "val1" {
		t.Fatal("failed to set key to val")
	}
	obj.Set("key2", "val2")
	if obj.MStr("key2") != "val2" {
		t.Fatal("failed to set key to val")
	}
}

func ExampleObj_Put() {
	jBar := Obj{OnError: fail}
	jBar.Put("foo", 4711)
	jBar.CObjs(kBaz).Put("quux", true)
	err := json.NewEncoder(os.Stdout).Encode(jBar.Bare)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// {"baz":{"quux":true},"foo":4711}
}

func ExampleObj_Sets() {
	bar := map[string]interface{}{"foo": 4711}
	jBar := Obj{Bare: bar, OnError: fail}
	jBar.Sets(kFoo, "baz")
	fmt.Println(jBar.Bare)
	jBar.Sets(kBaz, "should fail")
	fmt.Println(jBar.MBools(kBaz))
	// Output:
	// map[foo:baz]
	// ERROR: cannot set object member, key 'baz' not exists
	// ERROR: no boolean object member 'baz'
	// false
}
