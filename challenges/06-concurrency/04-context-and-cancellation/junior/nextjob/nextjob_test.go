package nextjob

import (
	"context"
	"testing"
)

func queueWith(vals ...string) <-chan string {
	ch := make(chan string, len(vals)+1)
	for _, v := range vals {
		ch <- v
	}
	return ch
}

func closedQueue() <-chan string {
	ch := make(chan string)
	close(ch)
	return ch
}

func cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func expired() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	cancel()
	return ctx
}

func TestNextJob(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name    string
		ctx     context.Context
		jobs    <-chan string
		want    string
		wantErr error
	}{
		{"takes_head_of_queue", live, queueWith("job-1", "job-2"), "job-1", nil},
		{"single_job", live, queueWith("job-9"), "job-9", nil},
		{"empty_job_id", live, queueWith(""), "", nil},
		{"queue_closed", live, closedQueue(), "", ErrQueueClosed},
		{"deploy_cancelled_worker", cancelled(), make(chan string), "", context.Canceled},
		{"worker_lease_expired", expired(), make(chan string), "", context.DeadlineExceeded},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NextJob(tc.ctx, tc.jobs)
			if err != tc.wantErr {
				t.Fatalf("NextJob() err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("NextJob() = %q, want %q", got, tc.want)
			}
		})
	}
}
