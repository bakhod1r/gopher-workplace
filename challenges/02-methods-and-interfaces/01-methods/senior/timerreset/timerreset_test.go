package timerreset

import (
	"testing"
	"time"
)

func TestSession(t *testing.T) {
	start := time.Unix(0, 0)
	s := &Session{lastPing: start, timeout: 5 * time.Second}

	if s.IsExpired(start.Add(4 * time.Second)) {
		t.Error("should not be expired at 4s")
	}
	if !s.IsExpired(start.Add(6 * time.Second)) {
		t.Error("should be expired at 6s")
	}

	s.Ping(start.Add(4 * time.Second))
	if s.IsExpired(start.Add(6 * time.Second)) {
		t.Error("should not be expired at 6s after ping at 4s")
	}
}
