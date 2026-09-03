package schedlatency

import (
	"math"
	"reflect"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestDelay(t *testing.T) {
	if d, ok := Delay(Event{1, 100, 150}); !ok || d != 50 {
		t.Errorf("Delay = %d, %v, want 50, true", d, ok)
	}
	if d, ok := Delay(Event{1, 100, 100}); !ok || d != 0 {
		t.Errorf("Delay = %d, %v, want 0, true — starting immediately is a zero delay", d, ok)
	}
	if _, ok := Delay(Event{1, 100, 99}); ok {
		t.Error("a running-before-runnable event reported ok")
	}
}

func TestDelays(t *testing.T) {
	got := Delays([]Event{{1, 0, 10}, {2, 0, 5}})
	if !reflect.DeepEqual(got, []int64{10, 5}) {
		t.Errorf("Delays = %v, want [10 5]", got)
	}
}

func TestDelaysSkipsMalformed(t *testing.T) {
	got := Delays([]Event{{1, 0, 10}, {2, 50, 40}, {3, 0, 5}})
	if !reflect.DeepEqual(got, []int64{10, 5}) {
		t.Errorf("Delays = %v, want [10 5]", got)
	}
	empty := Delays([]Event{{1, 50, 40}})
	if empty == nil || len(empty) != 0 {
		t.Errorf("Delays = %v, want empty non-nil slice", empty)
	}
}

func TestStats(t *testing.T) {
	mean, worst, ok := Stats([]Event{{1, 0, 10}, {2, 0, 20}})
	if !ok || !near(mean, 15) || worst != 20 {
		t.Errorf("Stats = %v, %d, %v, want 15, 20, true", mean, worst, ok)
	}
}

func TestStatsNoValidEvents(t *testing.T) {
	if _, _, ok := Stats(nil); ok {
		t.Error("Stats(nil) reported ok")
	}
	if _, _, ok := Stats([]Event{{1, 50, 40}}); ok {
		t.Error("Stats with only malformed events reported ok")
	}
}

func TestStatsMeanHidesTheTail(t *testing.T) {
	events := make([]Event, 0, 100)
	for i := 0; i < 99; i++ {
		events = append(events, Event{i, 0, 1})
	}
	events = append(events, Event{99, 0, 10_000})
	mean, worst, ok := Stats(events)
	if !ok {
		t.Fatal("Stats reported not ok")
	}
	if worst != 10_000 {
		t.Errorf("worst = %d, want 10000", worst)
	}
	if mean > 200 {
		t.Errorf("mean = %v, expected the mean to stay small while the worst case is huge", mean)
	}
}
