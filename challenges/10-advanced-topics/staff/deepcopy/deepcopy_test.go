package deepcopy

import (
	"reflect"
	"testing"
)

type node struct {
	Name  string
	Tags  []string
	Meta  map[string]int
	Child *node
	Fixed [2]int
}

func TestDeepCopyEqual(t *testing.T) {
	in := node{
		Name:  "root",
		Tags:  []string{"a", "b"},
		Meta:  map[string]int{"k": 1},
		Child: &node{Name: "kid", Tags: []string{"c"}},
		Fixed: [2]int{7, 8},
	}
	out, ok := DeepCopy(in).(node)
	if !ok {
		t.Fatalf("DeepCopy returned %T, want node", DeepCopy(in))
	}
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("DeepCopy = %+v, want %+v", out, in)
	}
}

func TestDeepCopySharesNothing(t *testing.T) {
	in := node{
		Tags:  []string{"a"},
		Meta:  map[string]int{"k": 1},
		Child: &node{Name: "kid"},
	}
	out := DeepCopy(in).(node)

	out.Tags[0] = "changed"
	if in.Tags[0] != "a" {
		t.Error("the slice is shared")
	}
	out.Meta["k"] = 99
	if in.Meta["k"] != 1 {
		t.Error("the map is shared")
	}
	out.Child.Name = "changed"
	if in.Child.Name != "kid" {
		t.Error("the pointed-at struct is shared")
	}
	if out.Child == in.Child {
		t.Error("the pointer was copied, not what it points at")
	}
}

func TestDeepCopyNils(t *testing.T) {
	in := node{Name: "bare"}
	out := DeepCopy(in).(node)
	if out.Tags != nil || out.Meta != nil || out.Child != nil {
		t.Errorf("out = %+v, want the nil fields preserved", out)
	}
}

func TestDeepCopyScalarsAndNil(t *testing.T) {
	if got := DeepCopy(7); got != 7 {
		t.Errorf("DeepCopy(7) = %v, want 7", got)
	}
	if got := DeepCopy("s"); got != "s" {
		t.Errorf("DeepCopy = %v, want s", got)
	}
	if got := DeepCopy(nil); got != nil {
		t.Errorf("DeepCopy(nil) = %v, want nil", got)
	}
}

func TestDeepCopyNestedSlices(t *testing.T) {
	in := [][]int{{1, 2}, {3}}
	out := DeepCopy(in).([][]int)
	out[0][0] = 99
	if in[0][0] != 1 {
		t.Error("the inner slice is shared")
	}
}
