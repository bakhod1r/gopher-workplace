package retrypolicy

import (
	"errors"
	"fmt"
	"testing"
)

func TestRetry(t *testing.T) {
	t.Run("succeeds_first", func(t *testing.T) {
		calls := 0
		if err := Retry(3, func() error { calls++; return nil }); err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if calls != 1 {
			t.Errorf("calls = %d, want 1", calls)
		}
	})

	t.Run("stops_on_non_retryable", func(t *testing.T) {
		calls := 0
		err := Retry(3, func() error { calls++; return ErrInvalid })
		if calls != 1 {
			t.Fatalf("calls = %d, want 1", calls)
		}
		if !errors.Is(err, ErrInvalid) {
			t.Errorf("err = %v, want ErrInvalid", err)
		}
	})

	t.Run("retries_transient", func(t *testing.T) {
		calls := 0
		err := Retry(3, func() error {
			calls++
			if calls < 3 {
				return fmt.Errorf("attempt %d: %w", calls, ErrTransient)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if calls != 3 {
			t.Errorf("calls = %d, want 3", calls)
		}
	})

	t.Run("exhausts_budget", func(t *testing.T) {
		calls := 0
		err := Retry(3, func() error { calls++; return ErrTransient })
		if calls != 3 {
			t.Fatalf("calls = %d, want 3", calls)
		}
		if !errors.Is(err, ErrTransient) {
			t.Errorf("err = %v, want it to match ErrTransient", err)
		}
	})
}
