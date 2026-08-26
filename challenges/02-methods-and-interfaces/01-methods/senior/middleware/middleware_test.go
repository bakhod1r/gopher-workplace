package middleware

import "testing"

func TestChain(t *testing.T) {
	mw1 := func(next Handler) Handler {
		return func(req string) string { return "1(" + next(req) + ")1" }
	}
	mw2 := func(next Handler) Handler {
		return func(req string) string { return "2(" + next(req) + ")2" }
	}

	chained := Chain(mw1, mw2)
	handler := chained(func(req string) string { return "H:" + req })

	got := handler("req")
	want := "1(2(H:req)2)1"
	if got != want {
		t.Errorf("Chain() = %q, want %q", got, want)
	}
}
