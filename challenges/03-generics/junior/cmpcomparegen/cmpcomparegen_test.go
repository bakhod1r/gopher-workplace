package cmpcomparegen

import "testing"

func TestByLengthThenName(t *testing.T) {
	if got := ByLengthThenName("a", "bb"); got >= 0 {
		t.Errorf(`ByLengthThenName("a", "bb") = %d, want a negative number`, got)
	}
	if got := ByLengthThenName("bb", "a"); got <= 0 {
		t.Errorf(`ByLengthThenName("bb", "a") = %d, want a positive number`, got)
	}
	if got := ByLengthThenName("bb", "aa"); got <= 0 {
		t.Errorf(`ByLengthThenName("bb", "aa") = %d, want a positive number (tie broken alphabetically)`, got)
	}
	if got := ByLengthThenName("aa", "bb"); got >= 0 {
		t.Errorf(`ByLengthThenName("aa", "bb") = %d, want a negative number`, got)
	}
	if got := ByLengthThenName("a", "a"); got != 0 {
		t.Errorf(`ByLengthThenName("a", "a") = %d, want 0`, got)
	}
}
