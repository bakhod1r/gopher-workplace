package triplegen

import "testing"

func TestMakeTriple(t *testing.T) {
	tr := MakeTriple(1, "a", true)
	if tr.First != 1 || tr.Second != "a" || tr.Third != true {
		t.Errorf("MakeTriple = %+v, want {1 a true}", tr)
	}
}

func TestRotated(t *testing.T) {
	r := MakeTriple(1, "a", true).Rotated()
	if r.First != "a" {
		t.Errorf("Rotated().First = %v, want a", r.First)
	}
	if r.Second != true {
		t.Errorf("Rotated().Second = %v, want true", r.Second)
	}
	if r.Third != 1 {
		t.Errorf("Rotated().Third = %v, want 1", r.Third)
	}
	back := MakeTriple(1, "a", true).Rotated().Rotated().Rotated()
	if back.First != 1 || back.Second != "a" || back.Third != true {
		t.Errorf("three rotations = %+v, want {1 a true}", back)
	}
}
