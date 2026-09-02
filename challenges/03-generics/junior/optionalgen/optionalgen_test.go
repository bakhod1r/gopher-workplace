package optionalgen

import "testing"

func TestSomeAndNone(t *testing.T) {
	if got := Some(5).Or(9); got != 5 {
		t.Errorf("Some(5).Or(9) = %v, want 5", got)
	}
	if got := None[int]().Or(9); got != 9 {
		t.Errorf("None[int]().Or(9) = %v, want 9", got)
	}
	if got := Some(0).Or(9); got != 0 {
		t.Errorf("Some(0).Or(9) = %v, want 0", got)
	}
	if got := None[string]().Or("def"); got != "def" {
		t.Errorf(`None[string]().Or("def") = %q, want "def"`, got)
	}
	var zero Optional[int]
	if got := zero.Or(7); got != 7 {
		t.Errorf("zero Optional.Or(7) = %v, want 7", got)
	}
}
