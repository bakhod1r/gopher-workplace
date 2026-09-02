package resultgen

import "testing"

func TestOk(t *testing.T) {
	v, ok := Ok(5).Unwrap()
	if v != 5 || !ok {
		t.Errorf("Ok(5).Unwrap() = %v, %v, want 5, true", v, ok)
	}
	if got := Ok(5).Reason(); got != "" {
		t.Errorf(`Ok(5).Reason() = %q, want ""`, got)
	}
}

func TestFail(t *testing.T) {
	v, ok := Fail[int]("bad").Unwrap()
	if v != 0 || ok {
		t.Errorf(`Fail[int]("bad").Unwrap() = %v, %v, want 0, false`, v, ok)
	}
	if got := Fail[int]("bad").Reason(); got != "bad" {
		t.Errorf(`Reason() = %q, want "bad"`, got)
	}
}

func TestResultStrings(t *testing.T) {
	v, ok := Ok("x").Unwrap()
	if v != "x" || !ok {
		t.Errorf(`Ok("x").Unwrap() = %q, %v, want x, true`, v, ok)
	}
}
