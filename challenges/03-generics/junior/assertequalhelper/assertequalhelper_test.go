package assertequalhelper

import "testing"

type fakeT struct {
	testing.TB
	failures int
	helper   bool
}

func (f *fakeT) Helper()                           { f.helper = true }
func (f *fakeT) Errorf(format string, args ...any) { f.failures++ }

func TestEqualPasses(t *testing.T) {
	f := &fakeT{}
	if !Equal(f, 1, 1) {
		t.Error("Equal(t, 1, 1) = false, want true")
	}
	if f.failures != 0 {
		t.Errorf("recorded %d failures, want 0", f.failures)
	}
	if !f.helper {
		t.Error("Equal did not call t.Helper()")
	}
}

func TestEqualFails(t *testing.T) {
	f := &fakeT{}
	if Equal(f, 1, 2) {
		t.Error("Equal(t, 1, 2) = true, want false")
	}
	if f.failures != 1 {
		t.Errorf("recorded %d failures, want 1", f.failures)
	}
}

func TestEqualStrings(t *testing.T) {
	f := &fakeT{}
	if !Equal(f, "a", "a") {
		t.Error(`Equal(t, "a", "a") = false, want true`)
	}
}
