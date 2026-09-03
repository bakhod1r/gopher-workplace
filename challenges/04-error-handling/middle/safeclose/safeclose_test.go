package safeclose

import (
	"errors"
	"testing"
)

func TestDo(t *testing.T) {
	errWork := errors.New("work failed")
	errCleanup := errors.New("cleanup failed")

	t.Run("both_succeed", func(t *testing.T) {
		calls := 0
		err := Do(
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

	t.Run("cleanup_runs_after_failure", func(t *testing.T) {
		cleaned := false
		err := Do(
			func() error { return errWork },
			func() error { cleaned = true; return nil },
		)
		if !cleaned {
			t.Error("cleanup did not run")
		}
		if !errors.Is(err, errWork) {
			t.Errorf("err = %v, want it to match errWork", err)
		}
	})

	t.Run("cleanup_failure_reported", func(t *testing.T) {
		err := Do(
			func() error { return nil },
			func() error { return errCleanup },
		)
		if !errors.Is(err, errCleanup) {
			t.Errorf("err = %v, want it to match errCleanup", err)
		}
	})

	t.Run("both_fail", func(t *testing.T) {
		err := Do(
			func() error { return errWork },
			func() error { return errCleanup },
		)
		if !errors.Is(err, errWork) || !errors.Is(err, errCleanup) {
			t.Errorf("err = %v, want it to match both failures", err)
		}
	})
}
