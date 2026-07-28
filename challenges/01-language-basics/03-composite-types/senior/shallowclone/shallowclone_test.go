package shallowclone

import "testing"

func TestClone(t *testing.T) {
	orig := Doc{Title: "a", Tags: []string{"x", "y"}}
	c := Clone(orig)
	c.Tags[0] = "MUT"
	if orig.Tags[0] != "x" {
		t.Errorf("clone shares Tags slice: orig=%v", orig.Tags)
	}
	if c.Title != "a" {
		t.Error("title not copied")
	}
}
