package stdmapsclonealiasbug

import (
	"reflect"
	"testing"
)

func TestCloneTagsIsDeep(t *testing.T) {
	m := map[string][]string{"a": {"x", "y"}}
	c := CloneTags(m)
	c["a"][0] = "z"
	if m["a"][0] != "x" {
		t.Errorf("original mutated through the clone: %v", m)
	}
}

func TestCloneTagsCopiesContent(t *testing.T) {
	m := map[string][]string{"a": {"x"}, "b": {}}
	c := CloneTags(m)
	if !reflect.DeepEqual(c["a"], []string{"x"}) {
		t.Errorf("clone[a] = %v, want [x]", c["a"])
	}
	if len(c) != 2 {
		t.Errorf("clone has %d keys, want 2", len(c))
	}
	c["new"] = []string{"q"}
	if _, ok := m["new"]; ok {
		t.Errorf("adding a key to the clone changed the original")
	}
}

func TestCloneTagsNil(t *testing.T) {
	if got := CloneTags(nil); got != nil {
		t.Errorf("CloneTags(nil) = %v, want nil", got)
	}
}
