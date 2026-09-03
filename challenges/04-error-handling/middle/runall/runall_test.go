package runall

import (
	"errors"
	"testing"
)

func TestRunAll(t *testing.T) {
	errA := errors.New("a failed")
	errB := errors.New("b failed")

	t.Run("no_steps", func(t *testing.T) {
		if got := RunAll(); got != nil {
			t.Errorf("RunAll() = %v, want nil", got)
		}
	})

	t.Run("all_succeed", func(t *testing.T) {
		calls := 0
		err := RunAll(
			func() error { calls++; return nil },
			func() error { calls++; return nil },
		)
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if calls != 2 {
			t.Errorf("calls = %d, want 2", calls)
		}
	})

	t.Run("runs_every_step", func(t *testing.T) {
		calls := 0
		err := RunAll(
			func() error { calls++; return errA },
			func() error { calls++; return nil },
			func() error { calls++; return errB },
		)
		if calls != 3 {
			t.Fatalf("calls = %d, want 3", calls)
		}
		if !errors.Is(err, errA) || !errors.Is(err, errB) {
			t.Errorf("err = %v, want it to match both failures", err)
		}
	})
}
