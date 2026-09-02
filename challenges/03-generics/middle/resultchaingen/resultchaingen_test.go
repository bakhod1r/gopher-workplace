package resultchaingen

import (
	"strconv"
	"testing"
)

func TestThenOnSuccess(t *testing.T) {
	double := func(n int) Result[int] { return Ok(n * 2) }
	v, ok := Then(Ok(2), double).Unwrap()
	if v != 4 || !ok {
		t.Errorf("Then = %v, %v, want 4, true", v, ok)
	}
}

func TestThenChangesType(t *testing.T) {
	show := func(n int) Result[string] { return Ok(strconv.Itoa(n)) }
	v, ok := Then(Ok(7), show).Unwrap()
	if v != "7" || !ok {
		t.Errorf("Then = %q, %v, want 7, true", v, ok)
	}
}

func TestThenShortCircuits(t *testing.T) {
	calls := 0
	f := func(n int) Result[int] { calls++; return Ok(n) }
	v, ok := Then(Fail[int]("bad"), f).Unwrap()
	if ok || v != 0 {
		t.Errorf("Then = %v, %v, want 0, false", v, ok)
	}
	if calls != 0 {
		t.Errorf("f called %d times, want 0", calls)
	}
}
