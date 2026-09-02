package queuepeek

import "testing"

func TestPeekJob(t *testing.T) {
	cases := []struct {
		name   string
		make   func() <-chan int
		wantID int
		wantOK bool
	}{
		{"job_queued", func() <-chan int {
			ch := make(chan int, 1)
			ch <- 5
			return ch
		}, 5, true},
		{"idle_runner", func() <-chan int {
			return make(chan int, 1)
		}, 0, false},
		{"closed_queue", func() <-chan int {
			ch := make(chan int)
			close(ch)
			return ch
		}, 0, true},
		{"unbuffered_idle", func() <-chan int {
			return make(chan int)
		}, 0, false},
		{"job_id_zero", func() <-chan int {
			ch := make(chan int, 1)
			ch <- 0
			return ch
		}, 0, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotID, gotOK := PeekJob(tc.make())
			if gotID != tc.wantID || gotOK != tc.wantOK {
				t.Errorf("PeekJob() = %d, %t, want %d, %t",
					gotID, gotOK, tc.wantID, tc.wantOK)
			}
		})
	}
}
