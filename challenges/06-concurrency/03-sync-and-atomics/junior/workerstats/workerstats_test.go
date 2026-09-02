package workerstats

import (
	"sync"
	"testing"
)

func TestPoolSnapshot(t *testing.T) {
	cases := []struct {
		name   string
		starts int
		fails  int
		want   Snapshot
	}{
		{"idle", 0, 0, Snapshot{0, 0}},
		{"one_start", 1, 0, Snapshot{1, 0}},
		{"one_failure", 1, 1, Snapshot{1, 1}},
		{"many_starts", 5, 0, Snapshot{5, 0}},
		{"mixed", 5, 2, Snapshot{5, 2}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var p Pool
			for i := 0; i < tc.starts; i++ {
				p.Start()
			}
			for i := 0; i < tc.fails; i++ {
				p.Fail()
			}
			if got := p.Snapshot(); got != tc.want {
				t.Errorf("Snapshot() = %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestPoolConcurrent(t *testing.T) {
	var p Pool
	const workers = 8
	const per = 250
	var wg sync.WaitGroup
	wg.Add(workers + 4)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < per; j++ {
				p.Start()
				p.Fail()
			}
		}()
	}
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				s := p.Snapshot()
				if s.Failed > s.Started {
					t.Errorf("inconsistent snapshot: %+v", s)
					return
				}
			}
		}()
	}
	wg.Wait()

	want := Snapshot{Started: workers * per, Failed: workers * per}
	if got := p.Snapshot(); got != want {
		t.Errorf("Snapshot() = %+v, want %+v", got, want)
	}
}
