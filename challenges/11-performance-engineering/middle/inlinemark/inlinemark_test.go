package inlinemark

import (
	"reflect"
	"testing"
)

func TestPhysical(t *testing.T) {
	got := Physical([]Frame{{"a", false}, {"b", true}, {"c", false}})
	if !reflect.DeepEqual(got, []string{"a", "c"}) {
		t.Errorf("Physical = %v, want [a c]", got)
	}
}

func TestPhysicalKeepsOrder(t *testing.T) {
	got := Physical([]Frame{{"main", false}, {"helper", true}, {"inner", true}, {"work", false}})
	if !reflect.DeepEqual(got, []string{"main", "work"}) {
		t.Errorf("Physical = %v, want [main work]", got)
	}
}

func TestPhysicalAllInlined(t *testing.T) {
	got := Physical([]Frame{{"a", true}, {"b", true}})
	if got == nil || len(got) != 0 {
		t.Errorf("Physical = %v, want empty non-nil slice", got)
	}
}

func TestPhysicalDoesNotModifyInput(t *testing.T) {
	in := []Frame{{"a", false}, {"b", true}}
	before := append([]Frame(nil), in...)
	Physical(in)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input changed: %v, want %v", in, before)
	}
}

func TestAttribute(t *testing.T) {
	if fn, ok := Attribute([]Frame{{"a", false}, {"b", true}}); fn != "a" || !ok {
		t.Errorf("Attribute = %q, %v, want a, true", fn, ok)
	}
	if fn, ok := Attribute([]Frame{{"a", false}, {"b", true}, {"c", false}}); fn != "c" || !ok {
		t.Errorf("Attribute = %q, %v, want c, true", fn, ok)
	}
	if fn, ok := Attribute([]Frame{{"a", false}}); fn != "a" || !ok {
		t.Errorf("Attribute = %q, %v, want a, true", fn, ok)
	}
}

func TestAttributeNothingPhysical(t *testing.T) {
	if fn, ok := Attribute([]Frame{{"a", true}, {"b", true}}); fn != "" || ok {
		t.Errorf("Attribute = %q, %v, want \"\", false", fn, ok)
	}
	if fn, ok := Attribute(nil); fn != "" || ok {
		t.Errorf("Attribute(nil) = %q, %v, want \"\", false", fn, ok)
	}
}
