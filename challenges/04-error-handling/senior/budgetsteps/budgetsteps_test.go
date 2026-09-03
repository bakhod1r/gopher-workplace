package budgetsteps

import (
	"errors"
	"testing"
)

func TestSpend(t *testing.T) {
	errStep := errors.New("step failed")

	t.Run("no_steps", func(t *testing.T) {
		if err := Spend(0); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("within_budget", func(t *testing.T) {
		ran := 0
		err := Spend(2,
			func() error { ran++; return nil },
			func() error { ran++; return nil },
		)
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if ran != 2 {
			t.Errorf("ran = %d, want 2", ran)
		}
	})

	t.Run("exceeds_budget", func(t *testing.T) {
		ran := 0
		err := Spend(1,
			func() error { ran++; return nil },
			func() error { ran++; return nil },
		)
		if !errors.Is(err, ErrBudgetExceeded) {
			t.Fatalf("err = %v, want ErrBudgetExceeded", err)
		}
		if ran != 1 {
			t.Errorf("ran = %d, want 1", ran)
		}
	})

	t.Run("step_failure_wins", func(t *testing.T) {
		ran := 0
		err := Spend(2,
			func() error { ran++; return errStep },
			func() error { ran++; return nil },
		)
		if !errors.Is(err, errStep) {
			t.Fatalf("err = %v, want errStep", err)
		}
		if ran != 1 {
			t.Errorf("ran = %d, want 1", ran)
		}
	})

	t.Run("zero_budget_with_steps", func(t *testing.T) {
		ran := 0
		err := Spend(0, func() error { ran++; return nil })
		if !errors.Is(err, ErrBudgetExceeded) || ran != 0 {
			t.Errorf("err = %v, ran = %d, want ErrBudgetExceeded and 0", err, ran)
		}
	})
}
