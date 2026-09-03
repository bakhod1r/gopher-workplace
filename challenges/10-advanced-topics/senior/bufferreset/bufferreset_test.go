package bufferreset

import (
	"sync"
	"testing"
)

func TestRenderOnce(t *testing.T) {
	if got := Render([]int{1, 2, 3}); got != "1-2-3" {
		t.Errorf("Render = %q, want \"1-2-3\"", got)
	}
}

func TestRenderRepeatedly(t *testing.T) {
	for i := 0; i < 200; i++ {
		if got := Render([]int{7}); got != "7" {
			t.Fatalf("call %d: Render = %q, want \"7\": the shared buffer was not reset", i, got)
		}
	}
}

func TestRenderEmpty(t *testing.T) {
	if got := Render(nil); got != "" {
		t.Errorf("Render = %q, want empty", got)
	}
}

func TestRenderStaysBounded(t *testing.T) {
	for i := 0; i < 2000; i++ {
		Render([]int{1, 2, 3, 4})
	}
	if got := len(Render([]int{9})); got != 1 {
		t.Errorf("Render returned %d bytes, want 1: the buffer keeps growing", got)
	}
}

func TestRenderConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	bad := make([]string, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			want := strconvItoa(w)
			for i := 0; i < 200; i++ {
				if got := Render([]int{w}); got != want {
					bad[w] = got
					return
				}
			}
		}(w)
	}
	wg.Wait()
	for w, got := range bad {
		if got != "" {
			t.Fatalf("worker %d saw %q", w, got)
		}
	}
}

func strconvItoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
