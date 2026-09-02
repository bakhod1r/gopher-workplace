package cloner

import "testing"

func TestClone(t *testing.T) {
	orig := &Config{Name: "db", Tags: []string{"a", "b"}}

	dup, ok := orig.Clone().(*Config)
	if !ok {
		t.Fatal("Clone did not return *Config")
	}
	if dup == orig {
		t.Fatal("Clone returned the same pointer")
	}
	if dup.Name != "db" {
		t.Errorf("Name = %q, want \"db\"", dup.Name)
	}

	dup.Tags[0] = "z"
	if orig.Tags[0] != "a" {
		t.Errorf("mutating the clone changed the original: %q", orig.Tags[0])
	}
}

func TestCopyOf(t *testing.T) {
	c := &Config{Name: "x", Tags: nil}
	got, ok := CopyOf(c).(*Config)
	if !ok {
		t.Fatal("CopyOf did not return *Config")
	}
	if got.Name != "x" {
		t.Errorf("Name = %q, want \"x\"", got.Name)
	}
	if got == c {
		t.Error("CopyOf returned the same pointer")
	}
}

func TestCloneEmptyTags(t *testing.T) {
	c := &Config{Name: "empty"}
	d := c.Clone().(*Config)
	if len(d.Tags) != 0 {
		t.Errorf("len(Tags) = %d, want 0", len(d.Tags))
	}
}
