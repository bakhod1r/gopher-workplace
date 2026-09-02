package iszerogen

import "testing"

func TestIsZero(t *testing.T) {
	if !IsZero(0) {
		t.Error("IsZero(0) = false, want true")
	}
	if !IsZero("") {
		t.Error(`IsZero("") = false, want true`)
	}
	if !IsZero(false) {
		t.Error("IsZero(false) = false, want true")
	}
	if IsZero(3) {
		t.Error("IsZero(3) = true, want false")
	}
	if IsZero("a") {
		t.Error(`IsZero("a") = true, want false`)
	}
	if IsZero(true) {
		t.Error("IsZero(true) = true, want false")
	}
}
