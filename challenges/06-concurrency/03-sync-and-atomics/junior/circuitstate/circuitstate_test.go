package circuitstate

import (
	"sync"
	"testing"
)

func TestBreaker(t *testing.T) {
	cases := []struct {
		name      string
		trips     int
		resets    int
		wantOpen  bool
		wantFirst bool
	}{
		{"healthy", 0, 0, false, true},
		{"tripped", 1, 0, true, true},
		{"tripped_twice", 2, 0, true, true},
		{"recovered", 1, 1, false, true},
		{"reset_when_closed", 0, 1, false, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var b Breaker
			for i := 0; i < tc.trips; i++ {
				if got := b.Trip(); i == 0 && got != tc.wantFirst {
					t.Errorf("first Trip() = %v, want %v", got, tc.wantFirst)
				}
			}
			for i := 0; i < tc.resets; i++ {
				b.Reset()
			}
			if got := b.Open(); got != tc.wantOpen {
				t.Errorf("Open() = %v, want %v", got, tc.wantOpen)
			}
		})
	}
}

func TestBreakerSingleTripper(t *testing.T) {
	var b Breaker
	const observers = 16
	tripped := make(chan bool, observers)
	var wg sync.WaitGroup
	wg.Add(observers)
	for i := 0; i < observers; i++ {
		go func() {
			defer wg.Done()
			tripped <- b.Trip()
			b.Open()
		}()
	}
	wg.Wait()
	close(tripped)

	n := 0
	for v := range tripped {
		if v {
			n++
		}
	}
	if n != 1 {
		t.Errorf("goroutines that tripped the breaker = %d, want 1", n)
	}
	if !b.Open() {
		t.Error("Open() = false, want true")
	}
	if !b.Reset() {
		t.Error("Reset() = false, want true")
	}
}
