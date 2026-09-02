package asserthelperpassbug

import (
	"fmt"
	"testing"
)

type fakeT struct{ msgs []string }

func (f *fakeT) Errorf(format string, args ...any) {
	f.msgs = append(f.msgs, fmt.Sprintf(format, args...))
}

func TestAssertEqualReportsFailure(t *testing.T) {
	f := &fakeT{}
	if AssertEqual(f, 1, 2) {
		t.Errorf("AssertEqual(1, 2) = true, want false")
	}
	if len(f.msgs) != 1 {
		t.Errorf("logged %d messages, want 1", len(f.msgs))
	}
}

func TestAssertEqualReportsSuccess(t *testing.T) {
	f := &fakeT{}
	if !AssertEqual(f, 1, 1) {
		t.Errorf("AssertEqual(1, 1) = false, want true")
	}
	if len(f.msgs) != 0 {
		t.Errorf("logged %v, want nothing", f.msgs)
	}
}

func TestAssertEqualStrings(t *testing.T) {
	f := &fakeT{}
	if AssertEqual(f, "a", "b") {
		t.Errorf("AssertEqual(a, b) = true, want false")
	}
}
