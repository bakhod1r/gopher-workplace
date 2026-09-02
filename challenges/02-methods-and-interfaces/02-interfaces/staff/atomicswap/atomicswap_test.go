package atomicswap

import (
	"sync"
	"testing"
)

func TestFailsOpenWithNoPolicy(t *testing.T) {
	var s Strategy
	if !s.Allow("anything") {
		t.Error("Allow = false with no policy set, want true (fail open)")
	}
	if s.Get() != nil {
		t.Error("Get should be nil with no policy set")
	}
}

func TestSetAndAllow(t *testing.T) {
	var s Strategy

	s.Set(DenyAll{})
	if s.Allow("k") {
		t.Error("DenyAll should deny")
	}

	s.Set(AllowAll{})
	if !s.Allow("k") {
		t.Error("AllowAll should allow")
	}

	s.Set(PrefixPolicy{Prefix: "ok-"})
	if !s.Allow("ok-1") {
		t.Error("prefix match should allow")
	}
	if s.Allow("no-1") {
		t.Error("prefix mismatch should deny")
	}
}

func TestGetReturnsCurrent(t *testing.T) {
	var s Strategy
	s.Set(DenyAll{})
	if _, ok := s.Get().(DenyAll); !ok {
		t.Errorf("Get = %T, want DenyAll", s.Get())
	}
	s.Set(AllowAll{})
	if _, ok := s.Get().(AllowAll); !ok {
		t.Errorf("Get = %T, want AllowAll", s.Get())
	}
}

func TestNilPolicyFailsOpen(t *testing.T) {
	var s Strategy
	s.Set(nil)
	if !s.Allow("k") {
		t.Error("a nil policy should fail open, not panic")
	}
}

func TestSwapUnderLoad(t *testing.T) {
	var s Strategy
	s.Set(AllowAll{})

	stop := make(chan struct{})

	var swapper sync.WaitGroup
	swapper.Add(1)
	go func() {
		defer swapper.Done()
		for i := 0; ; i++ {
			select {
			case <-stop:
				return
			default:
			}
			if i%2 == 0 {
				s.Set(AllowAll{})
			} else {
				s.Set(PrefixPolicy{Prefix: "ok-"})
			}
		}
	}()

	var readers sync.WaitGroup
	for r := 0; r < 4; r++ {
		readers.Add(1)
		go func() {
			defer readers.Done()
			for i := 0; i < 20000; i++ {
				// Either policy allows this key, so the result must be
				// stable regardless of which one a request observes.
				if !s.Allow("ok-key") {
					t.Error("observed a half-applied swap")
					return
				}
			}
		}()
	}

	readers.Wait()
	close(stop)
	swapper.Wait()
}

func TestPrefixPolicyEdges(t *testing.T) {
	p := PrefixPolicy{Prefix: "ok-"}
	if p.Allow("ok") {
		t.Error("a shorter key must not match")
	}
	if !p.Allow("ok-") {
		t.Error("an exact-length key should match")
	}
	if !(PrefixPolicy{}).Allow("") {
		t.Error("an empty prefix should match everything")
	}
}

func BenchmarkAllow(b *testing.B) {
	var s Strategy
	s.Set(PrefixPolicy{Prefix: "ok-"})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = s.Allow("ok-key")
		}
	})
}
