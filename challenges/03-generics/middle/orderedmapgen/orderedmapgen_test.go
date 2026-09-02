package orderedmapgen

import (
	"reflect"
	"testing"
)

func TestOrderedKeepsInsertionOrder(t *testing.T) {
	o := NewOrdered[string, int]()
	o.Set("b", 1)
	o.Set("a", 2)
	o.Set("c", 3)
	if got := o.Keys(); !reflect.DeepEqual(got, []string{"b", "a", "c"}) {
		t.Errorf("Keys() = %v, want [b a c]", got)
	}
}

func TestOrderedUpdateKeepsPosition(t *testing.T) {
	o := NewOrdered[string, int]()
	o.Set("a", 1)
	o.Set("b", 2)
	o.Set("a", 3)
	if got := o.Keys(); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Keys() = %v, want [a b]", got)
	}
	if v, ok := o.Get("a"); v != 3 || !ok {
		t.Errorf("Get(a) = %v, %v, want 3, true", v, ok)
	}
}

func TestOrderedKeysIsACopy(t *testing.T) {
	o := NewOrdered[string, int]()
	o.Set("a", 1)
	ks := o.Keys()
	ks[0] = "zzz"
	if again := o.Keys(); again[0] != "a" {
		t.Errorf("Keys() returned internal storage: %v", again)
	}
	if _, ok := o.Get("missing"); ok {
		t.Error("Get(missing) reported ok")
	}
}
