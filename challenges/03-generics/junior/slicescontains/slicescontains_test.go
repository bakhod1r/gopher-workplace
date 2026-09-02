package slicescontains

import "testing"

func TestHasTag(t *testing.T) {
	if !HasTag([]string{"a", "b"}, "a") {
		t.Error(`HasTag([]string{"a", "b"}, "a") = false, want true`)
	}
	if HasTag([]string{"a"}, "b") {
		t.Error(`HasTag([]string{"a"}, "b") = true, want false`)
	}
	if !HasTag([]int{1, 2}, 2) {
		t.Error("HasTag([]int{1, 2}, 2) = false, want true")
	}
	if HasTag([]int{}, 1) {
		t.Error("HasTag([]int{}, 1) = true, want false")
	}
}
