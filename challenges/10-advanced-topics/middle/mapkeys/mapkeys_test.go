package mapkeys

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	got := Keys(map[string]int{"b": 1, "a": 2, "c": 3})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Keys = %v, want [a b c]", got)
	}
}

func TestKeysIgnoresTheValueType(t *testing.T) {
	got := Keys(map[string][]byte{"z": nil, "y": nil})
	if !reflect.DeepEqual(got, []string{"y", "z"}) {
		t.Errorf("Keys = %v, want [y z]", got)
	}
}

func TestKeysRejectsOtherShapes(t *testing.T) {
	for _, in := range []any{nil, 3, []string{"a"}, map[int]string{1: "a"}} {
		if got := Keys(in); got != nil {
			t.Errorf("Keys(%#v) = %v, want nil", in, got)
		}
	}
}

func TestKeysEmptyMap(t *testing.T) {
	if got := Keys(map[string]int{}); len(got) != 0 {
		t.Errorf("Keys = %v, want empty", got)
	}
}

func TestKeysIsDeterministic(t *testing.T) {
	m := map[string]int{"d": 1, "a": 2, "c": 3, "b": 4}
	first := Keys(m)
	for i := 0; i < 50; i++ {
		if got := Keys(m); !reflect.DeepEqual(got, first) {
			t.Fatalf("run %d = %v, want %v: map iteration order must be sorted away", i, got, first)
		}
	}
}
