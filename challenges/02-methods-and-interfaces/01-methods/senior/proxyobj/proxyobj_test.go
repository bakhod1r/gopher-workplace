package proxyobj

import "testing"

func TestProxy(t *testing.T) {
	w := &Worker{}

	p1 := &Proxy{w: w, role: "user"}
	if got := p1.Do(); got != "denied" {
		t.Errorf("user got %q", got)
	}
	if w.calls != 0 {
		t.Error("worker called by user")
	}

	p2 := &Proxy{w: w, role: "admin"}
	if got := p2.Do(); got != "done" {
		t.Errorf("admin got %q", got)
	}
	if w.calls != 1 {
		t.Error("worker not called by admin")
	}
}
