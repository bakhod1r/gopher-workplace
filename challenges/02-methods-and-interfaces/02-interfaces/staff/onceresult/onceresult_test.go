package onceresult

import (
	"errors"
	"strings"
	"sync"
	"testing"
)

func TestRunsOnce(t *testing.T) {
	calls := 0
	o := NewOnceValue(IniterFunc(func() (string, error) {
		calls++
		return "v", nil
	}))

	for i := 0; i < 3; i++ {
		v, err := o.Get()
		if v != "v" || err != nil {
			t.Fatalf("Get = %q, %v", v, err)
		}
	}
	if calls != 1 {
		t.Errorf("initialiser ran %d times, want 1", calls)
	}
	if o.Runs() != 1 {
		t.Errorf("Runs = %d, want 1", o.Runs())
	}
}

func TestCachesError(t *testing.T) {
	boom := errors.New("boom")
	calls := 0
	o := NewOnceValue(IniterFunc(func() (string, error) {
		calls++
		return "", boom
	}))

	for i := 0; i < 3; i++ {
		v, err := o.Get()
		if !errors.Is(err, boom) {
			t.Fatalf("err = %v, want boom", err)
		}
		if v != "" {
			t.Fatalf("value = %q, want empty", v)
		}
	}
	if calls != 1 {
		t.Errorf("initialiser ran %d times, want 1 (failures are cached)", calls)
	}
}

func TestPanicBecomesError(t *testing.T) {
	calls := 0
	o := NewOnceValue(IniterFunc(func() (string, error) {
		calls++
		panic("kaboom")
	}))

	v, err := o.Get()
	if err == nil {
		t.Fatal("err = nil, want an error from the panic")
	}
	if !strings.Contains(err.Error(), "kaboom") {
		t.Errorf("err = %v, want it to mention the panic value", err)
	}
	if v != "" {
		t.Errorf("value = %q, want empty", v)
	}

	// The second call must not panic either.
	if _, err2 := o.Get(); err2 == nil {
		t.Error("second Get lost the cached error")
	}
	if calls != 1 {
		t.Errorf("initialiser ran %d times, want 1", calls)
	}
}

func TestConcurrentGet(t *testing.T) {
	var calls int64
	var mu sync.Mutex
	o := NewOnceValue(IniterFunc(func() (string, error) {
		mu.Lock()
		calls++
		mu.Unlock()
		return "shared", nil
	}))

	const n = 200
	results := make([]string, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v, err := o.Get()
			if err != nil {
				t.Errorf("Get = %v", err)
				return
			}
			results[i] = v
		}(i)
	}
	wg.Wait()

	if calls != 1 {
		t.Errorf("initialiser ran %d times, want 1", calls)
	}
	for i, r := range results {
		if r != "shared" {
			t.Fatalf("results[%d] = %q, want \"shared\"", i, r)
		}
	}
}

func TestEmptyValueIsCached(t *testing.T) {
	calls := 0
	o := NewOnceValue(IniterFunc(func() (string, error) {
		calls++
		return "", nil
	}))

	o.Get()
	o.Get()
	if calls != 1 {
		t.Errorf("initialiser ran %d times, want 1", calls)
	}
}
