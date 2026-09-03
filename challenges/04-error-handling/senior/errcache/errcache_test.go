package errcache

import (
	"errors"
	"testing"
)

func TestCache(t *testing.T) {
	t.Run("caches_success", func(t *testing.T) {
		calls := 0
		c := &Cache{Load: func(string) (int, error) { calls++; return 7, nil }}
		for i := 0; i < 3; i++ {
			v, err := c.Get("a")
			if err != nil || v != 7 {
				t.Fatalf("Get = %d, %v, want 7, nil", v, err)
			}
		}
		if calls != 1 {
			t.Errorf("Load called %d times, want 1", calls)
		}
	})

	t.Run("caches_failure", func(t *testing.T) {
		errBoom := errors.New("boom")
		calls := 0
		c := &Cache{Load: func(string) (int, error) { calls++; return 0, errBoom }}
		for i := 0; i < 3; i++ {
			if _, err := c.Get("a"); !errors.Is(err, errBoom) {
				t.Fatalf("err = %v, want errBoom", err)
			}
		}
		if calls != 1 {
			t.Errorf("Load called %d times, want 1", calls)
		}
	})

	t.Run("per_key", func(t *testing.T) {
		calls := 0
		c := &Cache{Load: func(k string) (int, error) { calls++; return len(k), nil }}
		c.Get("a")
		c.Get("bb")
		c.Get("a")
		if calls != 2 {
			t.Errorf("Load called %d times, want 2", calls)
		}
	})
}
