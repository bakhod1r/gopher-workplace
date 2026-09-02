package mapdeepclonegen

import "testing"

func TestDeepClone(t *testing.T) {
	m := map[string][]int{"a": {1, 2}, "b": {3}}
	got := DeepClone(m)
	if len(got) != 2 || len(got["a"]) != 2 || got["b"][0] != 3 {
		t.Fatalf("DeepClone = %v, want the same content", got)
	}
}

func TestDeepCloneIsIndependent(t *testing.T) {
	m := map[string][]int{"a": {1, 2}}
	got := DeepClone(m)
	got["a"][0] = 99
	if m["a"][0] != 1 {
		t.Errorf("writing into the copy changed the original: %v", m)
	}
	got["b"] = []int{4}
	if len(m) != 1 {
		t.Errorf("adding to the copy changed the original: %v", m)
	}
}

func TestDeepCloneNil(t *testing.T) {
	got := DeepClone(map[string][]int(nil))
	if got == nil {
		t.Fatal("DeepClone(nil) = nil, want an empty non-nil map")
	}
	if len(got) != 0 {
		t.Errorf("DeepClone(nil) = %v, want {}", got)
	}
}
