package twophase

import (
	"errors"
	"testing"
)

func TestDo(t *testing.T) {
	errApply := errors.New("apply failed")
	errConfirm := errors.New("confirm failed")
	errRollback := errors.New("rollback failed")
	ok := func() error { return nil }

	t.Run("happy_path", func(t *testing.T) {
		rolled := false
		err := Do(ok, ok, func() error { rolled = true; return nil })
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if rolled {
			t.Error("rollback ran on the happy path")
		}
	})

	t.Run("apply_fails", func(t *testing.T) {
		confirmed, rolled := false, false
		err := Do(
			func() error { return errApply },
			func() error { confirmed = true; return nil },
			func() error { rolled = true; return nil },
		)
		if !errors.Is(err, errApply) {
			t.Fatalf("err = %v, want errApply", err)
		}
		if confirmed || rolled {
			t.Error("confirm or rollback ran after a failed apply")
		}
	})

	t.Run("confirm_fails_rolls_back", func(t *testing.T) {
		rolled := false
		err := Do(ok, func() error { return errConfirm }, func() error { rolled = true; return nil })
		if !rolled {
			t.Error("rollback did not run")
		}
		if !errors.Is(err, errConfirm) {
			t.Errorf("err = %v, want errConfirm", err)
		}
	})

	t.Run("rollback_also_fails", func(t *testing.T) {
		err := Do(ok, func() error { return errConfirm }, func() error { return errRollback })
		if !errors.Is(err, errConfirm) || !errors.Is(err, errRollback) {
			t.Errorf("err = %v, want it to match both failures", err)
		}
	})
}
