package goroutinestates

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	got := Count([]G{
		{"running", "main.work"},
		{"chan receive", "main.consume"},
		{"chan receive", "main.consume"},
	})
	want := map[string]int{"running": 1, "chan receive": 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Count = %v, want %v", got, want)
	}
}

func TestCountUnknownState(t *testing.T) {
	got := Count([]G{{"", "x"}, {"running", "y"}})
	want := map[string]int{"unknown": 1, "running": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Count = %v, want %v", got, want)
	}
}

func TestCountEmpty(t *testing.T) {
	got := Count(nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Count(nil) = %v, want empty non-nil map", got)
	}
}

func TestBlocked(t *testing.T) {
	gs := []G{
		{"running", "a"},
		{"runnable", "b"},
		{"chan receive", "c"},
		{"select", "d"},
		{"IO wait", "e"},
	}
	if got := Blocked(gs); got != 3 {
		t.Errorf("Blocked = %d, want 3", got)
	}
	if got := Blocked(nil); got != 0 {
		t.Errorf("Blocked(nil) = %d, want 0", got)
	}
}

func TestLeakSuspects(t *testing.T) {
	gs := make([]G, 0, 210)
	for i := 0; i < 150; i++ {
		gs = append(gs, G{"chan receive", "main.consume"})
	}
	for i := 0; i < 50; i++ {
		gs = append(gs, G{"select", "main.dispatch"})
	}
	for i := 0; i < 10; i++ {
		gs = append(gs, G{"running", "main.work"})
	}
	if got := LeakSuspects(gs, 100); !reflect.DeepEqual(got, []string{"main.consume"}) {
		t.Errorf("LeakSuspects(100) = %v, want [main.consume]", got)
	}
	got := LeakSuspects(gs, 20)
	if !reflect.DeepEqual(got, []string{"main.consume", "main.dispatch"}) {
		t.Errorf("LeakSuspects(20) = %v, want [main.consume main.dispatch]", got)
	}
}

func TestLeakSuspectsIgnoresRunningGoroutines(t *testing.T) {
	gs := make([]G, 0, 100)
	for i := 0; i < 100; i++ {
		gs = append(gs, G{"running", "main.work"})
	}
	got := LeakSuspects(gs, 10)
	if got == nil || len(got) != 0 {
		t.Errorf("LeakSuspects = %v, want empty non-nil slice — running goroutines are not leaks", got)
	}
}

func TestLeakSuspectsThresholdGuard(t *testing.T) {
	gs := []G{{"chan receive", "a"}}
	for _, th := range []int{0, -5, 1} {
		if got := LeakSuspects(gs, th); !reflect.DeepEqual(got, []string{"a"}) {
			t.Errorf("LeakSuspects(%d) = %v, want [a]", th, got)
		}
	}
}
