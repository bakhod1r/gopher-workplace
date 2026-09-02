package lazyload

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestNotBuiltUntilUsed(t *testing.T) {
	var builds int64
	l := NewLazy(BuilderFunc(func() string {
		atomic.AddInt64(&builds, 1)
		return "v"
	}))

	if l.Built() {
		t.Error("Built = true before any Get")
	}
	if builds != 0 {
		t.Errorf("builder ran %d times before Get", builds)
	}
}

func TestBuildsOnce(t *testing.T) {
	var builds int64
	l := NewLazy(BuilderFunc(func() string {
		atomic.AddInt64(&builds, 1)
		return "v"
	}))

	if got := l.Get(); got != "v" {
		t.Errorf("Get = %q, want \"v\"", got)
	}
	if got := l.Get(); got != "v" {
		t.Errorf("Get = %q, want \"v\"", got)
	}
	if builds != 1 {
		t.Errorf("builder ran %d times, want 1", builds)
	}
	if !l.Built() {
		t.Error("Built = false after Get")
	}
}

func TestConcurrentGetBuildsOnce(t *testing.T) {
	var builds int64
	l := NewLazy(BuilderFunc(func() string {
		atomic.AddInt64(&builds, 1)
		return "shared"
	}))

	const n = 100
	results := make([]string, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = l.Get()
		}(i)
	}
	wg.Wait()

	if builds != 1 {
		t.Errorf("builder ran %d times, want 1", builds)
	}
	for i, r := range results {
		if r != "shared" {
			t.Fatalf("results[%d] = %q, want \"shared\"", i, r)
		}
	}
}

func TestEmptyValue(t *testing.T) {
	l := NewLazy(BuilderFunc(func() string { return "" }))
	if got := l.Get(); got != "" {
		t.Errorf("Get = %q, want empty", got)
	}
	if !l.Built() {
		t.Error("Built = false after building an empty value")
	}
}

func BenchmarkGetAfterBuild(b *testing.B) {
	l := NewLazy(BuilderFunc(func() string { return "v" }))
	l.Get()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = l.Get()
		}
	})
}
