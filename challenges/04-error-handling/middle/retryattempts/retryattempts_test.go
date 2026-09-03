package retryattempts

import (
	"errors"
	"testing"
)

func TestRetry(t *testing.T) {
	errBoom := errors.New("boom")

	t.Run("zero_attempts", func(t *testing.T) {
		calls := 0
		err := Retry(0, func() error { calls++; return nil })
		if !errors.Is(err, ErrNoAttempts) {
			t.Fatalf("err = %v, want ErrNoAttempts", err)
		}
		if calls != 0 {
			t.Errorf("calls = %d, want 0", calls)
		}
	})

	t.Run("first_try", func(t *testing.T) {
		calls := 0
		err := Retry(3, func() error { calls++; return nil })
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if calls != 1 {
			t.Errorf("calls = %d, want 1", calls)
		}
	})

	t.Run("succeeds_on_third", func(t *testing.T) {
		calls := 0
		err := Retry(3, func() error {
			calls++
			if calls < 3 {
				return errBoom
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

	t.Run("exhausted", func(t *testing.T) {
		calls := 0
		err := Retry(3, func() error { calls++; return errBoom })
		if calls != 3 {
			t.Fatalf("calls = %d, want 3", calls)
		}
		if !errors.Is(err, errBoom) {
			t.Fatalf("err = %v, want it to wrap errBoom", err)
		}
		if err.Error() != "after 3 attempts: boom" {
			t.Errorf("message = %q, want %q", err.Error(), "after 3 attempts: boom")
		}
	})
}
