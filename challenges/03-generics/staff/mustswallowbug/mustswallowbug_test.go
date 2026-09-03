package mustswallowbug

import (
	"errors"
	"testing"
)

func TestMustReturnsValue(t *testing.T) {
	if got := Must(7, nil); got != 7 {
		t.Errorf("Must = %d, want 7", got)
	}
	if got := Must("x", nil); got != "x" {
		t.Errorf("Must = %q, want \"x\"", got)
	}
}

func TestMustPanicsOnError(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Must returned instead of panicking")
		}
	}()
	_ = Must(0, errors.New("boom"))
}

func TestMustPanicsForStructType(t *testing.T) {
	type cfg struct{ Addr string }
	defer func() {
		if recover() == nil {
			t.Errorf("Must returned instead of panicking")
		}
	}()
	_ = Must(cfg{}, errors.New("missing file"))
}
