package readcopyp

import "testing"

func TestRCU(t *testing.T) {
	r := New()
	r.Update("v2")

	cfg := r.ptr.Load()
	if cfg.Data != "v2" {
		t.Errorf("got %q, want v2", cfg.Data)
	}
}
