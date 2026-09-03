package poolreset

import "testing"

func TestRenderOnce(t *testing.T) {
	if got := Render([]int{1, 2, 3}); got != "1,2,3" {
		t.Errorf("Render = %q, want \"1,2,3\"", got)
	}
}

func TestRenderRepeatedly(t *testing.T) {
	for i := 0; i < 200; i++ {
		if got := Render([]int{7}); got != "7" {
			t.Fatalf("call %d: Render = %q, want \"7\": the pooled buffer was not reset", i, got)
		}
	}
}

func TestRenderStaysBounded(t *testing.T) {
	for i := 0; i < 500; i++ {
		Render([]int{1, 2, 3, 4})
	}
	b := pool.Get().([]byte)
	if cap(b) > 4096 {
		t.Errorf("pooled buffer grew to cap %d: every call is appending to the last one's output", cap(b))
	}
}
