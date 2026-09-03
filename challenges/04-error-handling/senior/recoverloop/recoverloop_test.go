package recoverloop

import (
	"errors"
	"testing"
)

func TestProcess(t *testing.T) {
	errItem := errors.New("item failed")

	t.Run("empty", func(t *testing.T) {
		if err := Process(nil, func(int) error { return nil }); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})

	t.Run("all_ok", func(t *testing.T) {
		seen := 0
		if err := Process([]int{1, 2}, func(int) error { seen++; return nil }); err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if seen != 2 {
			t.Errorf("handled %d items, want 2", seen)
		}
	})

	t.Run("panic_isolated", func(t *testing.T) {
		seen := 0
		err := Process([]int{1, 2, 3}, func(n int) error {
			seen++
			if n == 2 {
				panic("boom")
			}
			return nil
		})
		if seen != 3 {
			t.Fatalf("handled %d items, want 3", seen)
		}
		if !errors.Is(err, ErrPanic) {
			t.Errorf("err = %v, want it to match ErrPanic", err)
		}
	})

	t.Run("errors_and_panics", func(t *testing.T) {
		err := Process([]int{1, 2}, func(n int) error {
			if n == 1 {
				return errItem
			}
			panic("boom")
		})
		if !errors.Is(err, errItem) || !errors.Is(err, ErrPanic) {
			t.Errorf("err = %v, want it to match both failures", err)
		}
	})
}
