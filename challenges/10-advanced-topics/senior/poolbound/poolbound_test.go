package poolbound

import "testing"

func TestRender(t *testing.T) {
	if got := Render(16); got != 16 {
		t.Errorf("Render = %d, want 16", got)
	}
	if got := Render(0); got != 0 {
		t.Errorf("Render = %d, want 0", got)
	}
	if got := Render(-5); got != 0 {
		t.Errorf("Render = %d, want 0", got)
	}
}

func TestRenderLarge(t *testing.T) {
	if got := Render(1 << 20); got != 1<<20 {
		t.Errorf("Render = %d, want %d", got, 1<<20)
	}
}

func TestOversizedBuffersAreNotPooled(t *testing.T) {
	// drain whatever is in the pool
	for i := 0; i < 8; i++ {
		PooledCap()
	}
	Render(1 << 20)
	for i := 0; i < 8; i++ {
		if c := PooledCap(); c > maxScratch {
			t.Fatalf("the pool holds a %d-byte buffer, want at most %d: drop oversized buffers", c, maxScratch)
		}
	}
}

func TestSmallBuffersAreStillPooled(t *testing.T) {
	Render(32)
	if c := PooledCap(); c == 0 {
		t.Error("the pool is empty after a small render: normal buffers must go back")
	}
}

func TestRenderStaysCorrectAfterALargeRequest(t *testing.T) {
	Render(1 << 20)
	for i := 0; i < 50; i++ {
		if got := Render(8); got != 8 {
			t.Fatalf("call %d: Render = %d, want 8", i, got)
		}
	}
}
