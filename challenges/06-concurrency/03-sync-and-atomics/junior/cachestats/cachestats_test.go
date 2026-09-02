package cachestats

import (
	"sync"
	"testing"
)

func TestStats(t *testing.T) {
	cases := []struct {
		name   string
		hits   int
		misses int
		want   float64
	}{
		{"idle_node", 0, 0, 0},
		{"all_hits", 3, 0, 1},
		{"all_misses", 0, 3, 0},
		{"half", 1, 1, 0.5},
		{"three_quarters", 3, 1, 0.75},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s Stats
			for i := 0; i < tc.hits; i++ {
				s.Hit()
			}
			for i := 0; i < tc.misses; i++ {
				s.Miss()
			}
			if got := s.Hits(); got != int64(tc.hits) {
				t.Errorf("Hits() = %d, want %d", got, tc.hits)
			}
			if got := s.Misses(); got != int64(tc.misses) {
				t.Errorf("Misses() = %d, want %d", got, tc.misses)
			}
			if got := s.Ratio(); got != tc.want {
				t.Errorf("Ratio() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestStatsConcurrent(t *testing.T) {
	var s Stats
	const workers = 8
	const per = 250
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < per; j++ {
				if i%2 == 0 {
					s.Hit()
				} else {
					s.Miss()
				}
				s.Ratio()
			}
		}(i)
	}
	wg.Wait()

	if got, want := s.Hits(), int64(workers/2*per); got != want {
		t.Errorf("Hits() = %d, want %d", got, want)
	}
	if got, want := s.Misses(), int64(workers/2*per); got != want {
		t.Errorf("Misses() = %d, want %d", got, want)
	}
	if got := s.Ratio(); got != 0.5 {
		t.Errorf("Ratio() = %v, want 0.5", got)
	}
}
