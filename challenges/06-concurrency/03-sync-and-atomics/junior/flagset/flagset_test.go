package flagset

import (
	"strconv"
	"sync"
	"testing"
)

type flagWrite struct {
	name string
	on   bool
}

func TestFlagSet(t *testing.T) {
	cases := []struct {
		name   string
		writes []flagWrite
		query  string
		want   bool
		wantN  int
	}{
		{"enabled", []flagWrite{{"new_ui", true}}, "new_ui", true, 1},
		{"disabled", []flagWrite{{"new_ui", false}}, "new_ui", false, 1},
		{"unknown_flag", []flagWrite{{"new_ui", true}}, "other", false, 1},
		{"empty_set", nil, "new_ui", false, 0},
		{"rollback", []flagWrite{{"a", true}, {"a", false}}, "a", false, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := NewFlagSet()
			for _, w := range tc.writes {
				f.Set(w.name, w.on)
			}
			if got := f.Enabled(tc.query); got != tc.want {
				t.Errorf("Enabled(%q) = %v, want %v", tc.query, got, tc.want)
			}
			if got := f.Len(); got != tc.wantN {
				t.Errorf("Len() = %d, want %d", got, tc.wantN)
			}
		})
	}
}

func TestFlagSetConcurrent(t *testing.T) {
	f := NewFlagSet()
	for i := 0; i < 10; i++ {
		f.Set(strconv.Itoa(i), true)
	}

	var wg sync.WaitGroup
	wg.Add(24)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				f.Enabled(strconv.Itoa(j % 10))
				f.Len()
			}
		}()
	}
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				f.Set(strconv.Itoa(j%10), j%2 == 0)
			}
		}()
	}
	wg.Wait()

	if got := f.Len(); got != 10 {
		t.Errorf("Len() = %d, want 10", got)
	}
}
