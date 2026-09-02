package parseintorgen

import "testing"

func TestParseOr(t *testing.T) {
	if got := ParseOr[int]("42", 0); got != 42 {
		t.Errorf(`ParseOr("42", 0) = %v, want 42`, got)
	}
	if got := ParseOr[int]("-7", 0); got != -7 {
		t.Errorf(`ParseOr("-7", 0) = %v, want -7`, got)
	}
	if got := ParseOr[int64]("9", 0); got != 9 {
		t.Errorf(`ParseOr("9", 0) = %v, want 9`, got)
	}
}

func TestParseOrFallsBack(t *testing.T) {
	if got := ParseOr[int]("abc", 7); got != 7 {
		t.Errorf(`ParseOr("abc", 7) = %v, want 7`, got)
	}
	if got := ParseOr[int]("", 3); got != 3 {
		t.Errorf(`ParseOr("", 3) = %v, want 3`, got)
	}
	if got := ParseOr[int]("1.5", 3); got != 3 {
		t.Errorf(`ParseOr("1.5", 3) = %v, want 3`, got)
	}
}

func TestParseOrInfersFromDefault(t *testing.T) {
	got := ParseOr("5", Retries(1))
	if got != Retries(5) {
		t.Errorf("ParseOr = %v, want Retries(5)", got)
	}
}
