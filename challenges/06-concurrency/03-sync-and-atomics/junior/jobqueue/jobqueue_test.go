package jobqueue

import (
	"strconv"
	"sync"
	"testing"
)

func TestJobQueueDrain(t *testing.T) {
	cases := []struct {
		name      string
		submitted []string
		takes     int
		want      []string
	}{
		{"single_job", []string{"a"}, 1, []string{"a"}},
		{"fifo_order", []string{"a", "b"}, 2, []string{"a", "b"}},
		{"closed_empty", nil, 1, nil},
		{"take_past_end", []string{"a"}, 2, []string{"a"}},
		{"three_jobs", []string{"a", "b", "c"}, 3, []string{"a", "b", "c"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			q := NewJobQueue()
			for _, j := range tc.submitted {
				q.Submit(j)
			}
			q.Close()

			var got []string
			for i := 0; i < tc.takes; i++ {
				job, ok := q.Take()
				if !ok {
					break
				}
				got = append(got, job)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("took %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("job %d = %q, want %q", i, got[i], tc.want[i])
				}
			}
		})
	}
}

func TestJobQueueBlockingWorkers(t *testing.T) {
	q := NewJobQueue()
	const workers = 4
	const jobs = 40

	taken := make(chan string, jobs)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				job, ok := q.Take()
				if !ok {
					return
				}
				taken <- job
			}
		}()
	}

	for i := 0; i < jobs; i++ {
		q.Submit(strconv.Itoa(i))
	}
	q.Close()
	wg.Wait()
	close(taken)

	if got := len(taken); got != jobs {
		t.Errorf("jobs taken = %d, want %d", got, jobs)
	}
}
