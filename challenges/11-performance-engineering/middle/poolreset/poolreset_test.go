package poolreset

import "testing"

func TestPutKeepsSmallBuffers(t *testing.T) {
	var p Pool
	if !p.Put(make([]byte, 0, 1024)) {
		t.Error("Put(1KB) = false, want true")
	}
	if !p.Put(make([]byte, 0, MaxCap)) {
		t.Error("Put(exactly MaxCap) = false, want true")
	}
	if got := p.Kept(); got != 2 {
		t.Errorf("Kept = %d, want 2", got)
	}
}

func TestPutDropsOversizedBuffers(t *testing.T) {
	var p Pool
	if p.Put(make([]byte, 0, MaxCap+1)) {
		t.Error("Put(MaxCap+1) = true, want false — an oversized buffer must not be retained")
	}
	if p.Put(nil) {
		t.Error("Put(nil) = true, want false")
	}
	if got := p.Kept(); got != 0 {
		t.Errorf("Kept = %d, want 0", got)
	}
}

func TestGetReturnsEmptyBufferOfAtLeastN(t *testing.T) {
	var p Pool
	b := p.Get(4096)
	if len(b) != 0 || cap(b) < 4096 {
		t.Errorf("len, cap = %d, %d, want 0 and at least 4096", len(b), cap(b))
	}
	if got := p.Get(0); cap(got) < 1024 {
		t.Errorf("cap = %d, want the default 1024", cap(got))
	}
}

func TestGetReusesAPooledBuffer(t *testing.T) {
	var p Pool
	b := make([]byte, 0, 2048)
	b = append(b, "old data"...)
	p.Put(b)
	got := p.Get(1024)
	if len(got) != 0 {
		t.Errorf("reused buffer has len %d, want 0", len(got))
	}
}

func TestGetSkipsTooSmallPooledBuffers(t *testing.T) {
	var p Pool
	p.Put(make([]byte, 0, 512))
	got := p.Get(8192)
	if cap(got) < 8192 {
		t.Errorf("cap = %d, want at least 8192 — a too-small pooled buffer cannot satisfy the request", cap(got))
	}
}
