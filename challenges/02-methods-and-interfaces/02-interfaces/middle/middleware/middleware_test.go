package middleware

import "testing"

func echo() Handler {
	return HandlerFunc(func(s string) string { return s })
}

func TestWithPrefix(t *testing.T) {
	h := WithPrefix("p:")(echo())
	if got := h.Handle("x"); got != "p:x" {
		t.Errorf("Handle = %q, want \"p:x\"", got)
	}
}

func TestWithCount(t *testing.T) {
	var n int
	h := WithCount(&n)(echo())
	h.Handle("a")
	h.Handle("b")
	if n != 2 {
		t.Errorf("count = %d, want 2", n)
	}
}

func TestApplyOrder(t *testing.T) {
	h := Apply(echo(), WithPrefix("a:"), WithPrefix("b:"))
	if got := h.Handle("x"); got != "a:b:x" {
		t.Errorf("Handle = %q, want \"a:b:x\"", got)
	}
}

func TestApplyEmpty(t *testing.T) {
	h := Apply(echo())
	if got := h.Handle("x"); got != "x" {
		t.Errorf("Handle = %q, want \"x\"", got)
	}
}

func TestApplyMixed(t *testing.T) {
	var n int
	h := Apply(echo(), WithCount(&n), WithPrefix("p:"))
	if got := h.Handle("x"); got != "p:x" {
		t.Errorf("Handle = %q, want \"p:x\"", got)
	}
	if n != 1 {
		t.Errorf("count = %d, want 1", n)
	}
}
