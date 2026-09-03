package errcollector

import (
	"errors"
	"testing"
)

func TestCollector(t *testing.T) {
	errA := errors.New("a")
	errB := errors.New("b")
	errC := errors.New("c")

	t.Run("empty", func(t *testing.T) {
		c := &Collector{Limit: 2}
		if c.Count() != 0 {
			t.Errorf("Count = %d, want 0", c.Count())
		}
		if c.Err() != nil {
			t.Errorf("Err = %v, want nil", c.Err())
		}
	})

	t.Run("ignores_nil", func(t *testing.T) {
		c := &Collector{Limit: 2}
		c.Add(nil)
		if c.Count() != 0 || c.Err() != nil {
			t.Errorf("Count = %d, Err = %v, want 0, nil", c.Count(), c.Err())
		}
	})

	t.Run("counts_beyond_limit", func(t *testing.T) {
		c := &Collector{Limit: 2}
		c.Add(errA)
		c.Add(errB)
		c.Add(errC)
		if c.Count() != 3 {
			t.Errorf("Count = %d, want 3", c.Count())
		}
		err := c.Err()
		if !errors.Is(err, errA) || !errors.Is(err, errB) {
			t.Errorf("Err = %v, want it to match the first two", err)
		}
		if errors.Is(err, errC) {
			t.Error("Err matched an error past the limit")
		}
	})

	t.Run("zero_limit_stores_nothing", func(t *testing.T) {
		c := &Collector{Limit: 0}
		c.Add(errA)
		if c.Count() != 1 {
			t.Errorf("Count = %d, want 1", c.Count())
		}
		if c.Err() != nil {
			t.Errorf("Err = %v, want nil", c.Err())
		}
	})
}
