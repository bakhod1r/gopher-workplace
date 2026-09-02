package versionedgen

import "testing"

func TestVersioned(t *testing.T) {
	var v Versioned[int]
	if got, ok := v.Get(); got != 0 || ok {
		t.Errorf("Get() before Set = %v, %v, want 0, false", got, ok)
	}
	v.Set(1)
	v.Set(2)
	if got, ok := v.Get(); got != 2 || !ok {
		t.Errorf("Get() = %v, %v, want 2, true", got, ok)
	}
	if v.Versions() != 2 {
		t.Errorf("Versions() = %d, want 2", v.Versions())
	}
}

func TestVersionedUndo(t *testing.T) {
	var v Versioned[string]
	v.Set("a")
	v.Set("b")
	if !v.Undo() {
		t.Fatal("Undo() = false, want true")
	}
	if got, _ := v.Get(); got != "a" {
		t.Errorf("Get() after undo = %q, want a", got)
	}
	if !v.Undo() {
		t.Fatal("Undo() = false, want true")
	}
	if _, ok := v.Get(); ok {
		t.Error("Get() after undoing everything reported ok")
	}
	if v.Undo() {
		t.Error("Undo() with no history = true, want false")
	}
}
