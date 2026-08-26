package chainresp

import "testing"

func TestChain(t *testing.T) {
	h10 := &H10{}
	h20 := &H20{}
	h10.SetNext(h20)

	if got := h10.Handle(10); got != "ten" {
		t.Errorf("10 = %q", got)
	}
	if got := h10.Handle(20); got != "twenty" {
		t.Errorf("20 = %q", got)
	}
	if got := h10.Handle(30); got != "unhandled" {
		t.Errorf("30 = %q", got)
	}
}
