package rolloutplan

import (
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func TestPlanPercent(t *testing.T) {
	cases := []struct {
		name    string
		initial map[string]int
		sets    map[string]int
		lookup  string
		want    int
	}{
		{"from_initial", map[string]int{"checkout": 25}, nil, "checkout", 25},
		{"unknown_feature", map[string]int{"checkout": 25}, nil, "search", 0},
		{"nil_initial", nil, nil, "checkout", 0},
		{"set_overrides", map[string]int{"checkout": 25}, map[string]int{"checkout": 60}, "checkout", 60},
		{"clamped_high", nil, map[string]int{"checkout": 300}, "checkout", 100},
		{"clamped_low", map[string]int{"checkout": 40}, map[string]int{"checkout": -5}, "checkout", 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewPlan(tc.initial)
			for f, pct := range tc.sets {
				p.SetPercent(f, pct)
			}
			if got := p.Percent(tc.lookup); got != tc.want {
				t.Errorf("Percent(%q) = %d, want %d", tc.lookup, got, tc.want)
			}
		})
	}
}

func TestPlanDoesNotAliasInitial(t *testing.T) {
	initial := map[string]int{"checkout": 25}
	p := NewPlan(initial)
	initial["checkout"] = 99
	if got := p.Percent("checkout"); got != 25 {
		t.Errorf("Percent after mutating caller's map = %d, want 25", got)
	}
}

func TestPlanFeaturesSorted(t *testing.T) {
	p := NewPlan(map[string]int{"search": 10})
	p.SetPercent("checkout", 20)
	p.SetPercent("billing", 30)
	want := []string{"billing", "checkout", "search"}
	if got := p.Features(); !reflect.DeepEqual(got, want) {
		t.Errorf("Features() = %v, want %v", got, want)
	}
}

func TestPlanConcurrentReadersAndWriters(t *testing.T) {
	p := NewPlan(map[string]int{"checkout": 1})
	const readers, writers, perGoroutine = 8, 4, 300

	var wg sync.WaitGroup
	wg.Add(readers + writers)
	for r := 0; r < readers; r++ {
		go func() {
			defer wg.Done()
			for i := 0; i < perGoroutine; i++ {
				if got := p.Percent("checkout"); got < 1 || got > 100 {
					t.Errorf("Percent(checkout) = %d, outside 1..100", got)
					return
				}
				_ = p.Features()
			}
		}()
	}
	for w := 0; w < writers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < perGoroutine; i++ {
				p.SetPercent("flag"+strconv.Itoa(w), i%101)
			}
		}(w)
	}
	wg.Wait()

	// Every writer's final flag must survive: no update may be lost.
	for w := 0; w < writers; w++ {
		name := "flag" + strconv.Itoa(w)
		if got := p.Percent(name); got != (perGoroutine-1)%101 {
			t.Errorf("Percent(%q) = %d, want %d", name, got, (perGoroutine-1)%101)
		}
	}
	if got, want := len(p.Features()), writers+1; got != want {
		t.Errorf("len(Features()) = %d, want %d", got, want)
	}
}
