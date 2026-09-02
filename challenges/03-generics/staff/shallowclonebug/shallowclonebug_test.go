package shallowclonebug

import (
	"reflect"
	"testing"
)

func TestCloneDocCopiesFields(t *testing.T) {
	d := Doc[string]{Title: "spec", Tags: []string{"go", "draft"}}
	c := CloneDoc(d)
	if c.Title != "spec" {
		t.Errorf("Title = %q, want spec", c.Title)
	}
	if !reflect.DeepEqual(c.Tags, []string{"go", "draft"}) {
		t.Errorf("Tags = %v, want [go draft]", c.Tags)
	}
}

func TestCloneDocIsIndependent(t *testing.T) {
	d := Doc[string]{Title: "spec", Tags: []string{"go", "draft"}}
	c := CloneDoc(d)
	c.Tags[0] = "edited"
	if d.Tags[0] != "go" {
		t.Errorf("d.Tags[0] = %q, want go: the clone shares the original's storage", d.Tags[0])
	}
}

func TestCloneAllIsIndependent(t *testing.T) {
	ds := []Doc[int]{{Title: "a", Tags: []int{1, 2}}, {Title: "b", Tags: []int{3}}}
	cs := CloneAll(ds)
	cs[0].Tags[0] = 99
	cs[1].Tags[0] = 99
	if ds[0].Tags[0] != 1 || ds[1].Tags[0] != 3 {
		t.Errorf("originals = %v, %v, want [1 2] and [3]", ds[0].Tags, ds[1].Tags)
	}
}

func TestCloneDocEmptyTags(t *testing.T) {
	c := CloneDoc(Doc[string]{Title: "x"})
	if len(c.Tags) != 0 {
		t.Errorf("Tags = %v, want empty", c.Tags)
	}
}
