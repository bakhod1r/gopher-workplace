package retrytwice

import (
	"errors"
	"testing"
)

func TestRetryTwice(t *testing.T) {
	errFirst := errors.New("first")
	errSecond := errors.New("second")

	t.Run("succeeds_immediately", func(t *testing.T) {
		calls := 0
		err := RetryTwice(func() error { calls++; return nil })
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if calls != 1 {
			t.Errorf("calls = %d, want 1", calls)
		}
	})

	t.Run("succeeds_on_retry", func(t *testing.T) {
		calls := 0
		err := RetryTwice(func() error {
			calls++
			if calls == 1 {
				return errFirst
			}
			return nil
		})
		if err != nil {
			t.Fatalf("err = %v, want nil", err)
		}
		if calls != 2 {
			t.Errorf("calls = %d, want 2", calls)
		}
	})

	t.Run("both_fail", func(t *testing.T) {
		calls := 0
		err := RetryTwice(func() error {
			calls++
			if calls == 1 {
				return errFirst
			}
			return errSecond
		})
		if !errors.Is(err, errSecond) {
			t.Fatalf("err = %v, want %v", err, errSecond)
		}
		if calls != 2 {
			t.Errorf("calls = %d, want 2", calls)
		}
	})
}
