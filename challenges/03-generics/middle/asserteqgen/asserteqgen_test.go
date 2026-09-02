package asserteqgen

import "testing"

type recorder struct {
	testing.TB
	failures int
	helper   bool
	last     string
}

func (r *recorder) Helper()                           { r.helper = true }
func (r *recorder) Errorf(format string, args ...any) { r.failures++ }

func TestEqualSlicePasses(t *testing.T) {
	r := &recorder{}
	if !EqualSlice(r, []int{1, 2}, []int{1, 2}) {
		t.Error("EqualSlice = false, want true")
	}
	if r.failures != 0 {
		t.Errorf("failures = %d, want 0", r.failures)
	}
	if !r.helper {
		t.Error("EqualSlice did not call t.Helper()")
	}
}

func TestEqualSliceLengthMismatch(t *testing.T) {
	r := &recorder{}
	if EqualSlice(r, []int{1}, []int{1, 2}) {
		t.Error("EqualSlice = true, want false")
	}
	if r.failures != 1 {
		t.Errorf("failures = %d, want exactly 1", r.failures)
	}
}

func TestEqualSliceElementMismatch(t *testing.T) {
	r := &recorder{}
	if EqualSlice(r, []string{"a", "b"}, []string{"a", "c"}) {
		t.Error("EqualSlice = true, want false")
	}
	if r.failures != 1 {
		t.Errorf("failures = %d, want exactly 1 (report the first difference only)", r.failures)
	}
}

func TestEqualSliceEmpty(t *testing.T) {
	r := &recorder{}
	if !EqualSlice(r, []int{}, []int(nil)) {
		t.Error("EqualSlice(empty, nil) = false, want true")
	}
}
