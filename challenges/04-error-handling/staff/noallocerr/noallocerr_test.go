package noallocerr

import (
	"errors"
	"testing"
)

func TestCheck(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		if err := Check(""); !errors.Is(err, ErrEmpty) {
			t.Errorf("err = %v, want ErrEmpty", err)
		}
	})

	t.Run("non_empty", func(t *testing.T) {
		if err := Check("a"); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("identity", func(t *testing.T) {
		if Check("") != error(ErrEmpty) {
			t.Error("Check returned a different value, want the sentinel itself")
		}
	})

	t.Run("no_allocations", func(t *testing.T) {
		if n := testing.AllocsPerRun(100, func() { _ = Check("") }); n != 0 {
			t.Errorf("failure path allocated %v times per run, want 0", n)
		}
		if n := testing.AllocsPerRun(100, func() { _ = Check("a") }); n != 0 {
			t.Errorf("success path allocated %v times per run, want 0", n)
		}
	})
}
