package orderedmapdupbug

import (
	"reflect"
	"testing"
)

func TestUpdateKeepsOneKey(t *testing.T) {
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

func TestKeysMatchEntries(t *testing.T) {
	o := NewOrdered[string, int]()
	for i := 0; i < 5; i++ {
		o.Set("x", i)
	}
	if got := o.Keys(); len(got) != 1 {
		t.Errorf("Keys() = %v, want exactly one key", got)
	}
}

func TestInsertionOrder(t *testing.T) {
	o := NewOrdered[string, int]()
	o.Set("b", 1)
	o.Set("a", 2)
	if got := o.Keys(); !reflect.DeepEqual(got, []string{"b", "a"}) {
		t.Errorf("Keys() = %v, want [b a]", got)
	}
}
