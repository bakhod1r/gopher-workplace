package thumbnailworkers

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
	"testing"
)

var errRender = errors.New("render failed")

func cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func uploads(keys ...string) []Upload {
	out := make([]Upload, len(keys))
	for i, k := range keys {
		out[i] = Upload{Key: k}
	}
	return out
}

// render fails for any key starting with "bad".
func render(_ context.Context, u Upload) (Thumbnail, error) {
	if strings.HasPrefix(u.Key, "bad") {
		return Thumbnail{}, errRender
	}
	return Thumbnail{Key: u.Key + ".thumb", Width: len(u.Key)}, nil
}

func TestRenderQueue(t *testing.T) {
	cases := []struct {
		name    string
		ctx     context.Context
		queue   []Upload
		workers int
		want    []string
		wantErr error
	}{
		{"three_two_workers", context.Background(), uploads("a", "bb", "ccc"), 2, []string{"a.thumb", "bb.thumb", "ccc.thumb"}, nil},
		{"order_preserved", context.Background(), uploads("z", "m", "a"), 4, []string{"z.thumb", "m.thumb", "a.thumb"}, nil},
		{"single_worker", context.Background(), uploads("p", "q"), 1, []string{"p.thumb", "q.thumb"}, nil},
		{"one_bad_upload", context.Background(), uploads("a", "bad-1", "c"), 3, nil, errRender},
		{"empty_queue", context.Background(), nil, 2, nil, nil},
		{"client_hung_up", cancelled(), uploads("a", "b"), 2, nil, context.Canceled},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := RenderQueue(tc.ctx, tc.queue, tc.workers, render)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("RenderQueue() error = %v, want %v", err, tc.wantErr)
			}
			if err != nil {
				if got != nil {
					t.Fatalf("RenderQueue() = %v, want nil results on error", got)
				}
				return
			}
			if len(got) != len(tc.want) {
				t.Fatalf("RenderQueue() = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i].Key != tc.want[i] {
					t.Fatalf("RenderQueue() = %v, want %v", got, tc.want)
				}
			}
		})
	}
}

func TestRenderQueueUsesEveryWorkerOnce(t *testing.T) {
	var calls int64
	counting := func(ctx context.Context, u Upload) (Thumbnail, error) {
		atomic.AddInt64(&calls, 1)
		return render(ctx, u)
	}
	queue := uploads("a", "b", "c", "d", "e", "f")
	got, err := RenderQueue(context.Background(), queue, 3, counting)
	if err != nil {
		t.Fatalf("RenderQueue() error = %v", err)
	}
	if len(got) != len(queue) {
		t.Fatalf("got %d thumbnails, want %d", len(got), len(queue))
	}
	if n := atomic.LoadInt64(&calls); n != int64(len(queue)) {
		t.Errorf("render called %d times, want %d — every upload exactly once", n, len(queue))
	}
}
