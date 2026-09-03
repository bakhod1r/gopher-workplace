package firstsuccess

import (
	"errors"
	"testing"
)

func TestFirst(t *testing.T) {
	errA := errors.New("a failed")
	errB := errors.New("b failed")

	t.Run("no_sources", func(t *testing.T) {
		v, err := First()
		if !errors.Is(err, ErrNoSources) || v != 0 {
			t.Errorf("First() = %d, %v, want 0, ErrNoSources", v, err)
		}
	})

	t.Run("first_wins", func(t *testing.T) {
		calls := 0
		v, err := First(
			func() (int, error) { calls++; return 7, nil },
			func() (int, error) { calls++; return 9, nil },
		)
		if err != nil || v != 7 {
			t.Fatalf("First = %d, %v, want 7, nil", v, err)
		}
		if calls != 1 {
			t.Errorf("calls = %d, want 1", calls)
		}
	})

	t.Run("falls_back", func(t *testing.T) {
		v, err := First(
			func() (int, error) { return 0, errA },
			func() (int, error) { return 7, nil },
		)
		if err != nil || v != 7 {
			t.Errorf("First = %d, %v, want 7, nil", v, err)
		}
	})

	t.Run("all_fail", func(t *testing.T) {
		v, err := First(
			func() (int, error) { return 0, errA },
			func() (int, error) { return 0, errB },
		)
		if v != 0 {
			t.Errorf("v = %d, want 0", v)
		}
		if !errors.Is(err, errA) || !errors.Is(err, errB) {
			t.Errorf("err = %v, want it to match both failures", err)
		}
	})
}
