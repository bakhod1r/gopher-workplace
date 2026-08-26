package backoff

import (
	"testing"
	"time"
)

func TestBackoff(t *testing.T) {
	b := New(5 * time.Second)

	wants := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		4 * time.Second,
		5 * time.Second,
		5 * time.Second,
	}

	for i, want := range wants {
		if got := b.Next(); got != want {
			t.Errorf("call %d: got %v, want %v", i, got, want)
		}
	}
}
