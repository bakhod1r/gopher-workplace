package ratewindow

import (
	"math"
	"testing"
)

var events = []Event{
	{0, 1},
	{50, 2},
	{99, 4},
	{100, 8},
	{250, 16},
}

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestCountIn(t *testing.T) {
	cases := []struct {
		from, width int64
		want        int
	}{
		{0, 100, 3},
		{0, 101, 4},
		{100, 100, 1},
		{100, 200, 2},
		{300, 100, 0},
		{-100, 50, 0},
		{0, 0, 0},
		{0, -5, 0},
	}
	for _, c := range cases {
		if got := CountIn(events, c.from, c.width); got != c.want {
			t.Errorf("CountIn(%d, %d) = %d, want %d", c.from, c.width, got, c.want)
		}
	}
}

func TestCountInIsHalfOpen(t *testing.T) {
	// The event at 100 belongs to [100,200), never to [0,100).
	if got := CountIn(events, 0, 100); got != 3 {
		t.Errorf("CountIn = %d, want 3", got)
	}
	if got := CountIn(events, 100, 1); got != 1 {
		t.Errorf("CountIn = %d, want 1", got)
	}
}

func TestSumIn(t *testing.T) {
	if got := SumIn(events, 0, 100); got != 7 {
		t.Errorf("SumIn = %d, want 7", got)
	}
	if got := SumIn(events, 100, 200); got != 24 {
		t.Errorf("SumIn = %d, want 24", got)
	}
	if got := SumIn(events, 0, 0); got != 0 {
		t.Errorf("SumIn = %d, want 0", got)
	}
}

func TestRatePerSec(t *testing.T) {
	sec := int64(1_000_000_000)
	e := []Event{{0, 1}, {1, 1}, {2, 1}, {sec + 5, 1}}
	if got := RatePerSec(e, 0, sec); !near(got, 3) {
		t.Errorf("RatePerSec = %v, want 3", got)
	}
	if got := RatePerSec(e, 0, sec/2); !near(got, 6) {
		t.Errorf("RatePerSec = %v, want 6 — a half-second window doubles the rate", got)
	}
	if got := RatePerSec(e, 0, 0); got != 0 {
		t.Errorf("RatePerSec = %v, want 0", got)
	}
}

func TestScalesToManyEvents(t *testing.T) {
	big := make([]Event, 0, 100_000)
	for i := 0; i < 100_000; i++ {
		big = append(big, Event{NS: int64(i) * 1000, Value: 1})
	}
	if got := CountIn(big, 0, 1000); got != 1 {
		t.Errorf("CountIn = %d, want 1", got)
	}
	if got := CountIn(big, 50_000_000, 1_000_000); got != 1000 {
		t.Errorf("CountIn = %d, want 1000", got)
	}
	if got := SumIn(big, 50_000_000, 1_000_000); got != 1000 {
		t.Errorf("SumIn = %d, want 1000", got)
	}
}
