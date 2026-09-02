package funcadapt

import (
	"strings"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	h := HandlerFunc(strings.ToUpper)
	if got := h.Handle("hi"); got != "HI" {
		t.Errorf("Handle = %q, want \"HI\"", got)
	}
	if got := Run(h, "abc"); got != "ABC" {
		t.Errorf("Run = %q, want \"ABC\"", got)
	}
}

func TestChain(t *testing.T) {
	upper := HandlerFunc(strings.ToUpper)
	exclaim := HandlerFunc(func(s string) string { return s + "!" })

	cases := []struct {
		name string
		hs   []Handler
		in   string
		want string
	}{
		{"two", []Handler{upper, exclaim}, "hi", "HI!"},
		{"order_matters", []Handler{exclaim, upper}, "hi", "HI!"},
		{"one", []Handler{exclaim}, "hi", "hi!"},
		{"none", nil, "hi", "hi"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Chain(tc.hs...).Handle(tc.in); got != tc.want {
				t.Errorf("Chain = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestChainIsHandler(t *testing.T) {
	var h Handler = Chain(HandlerFunc(strings.ToLower))
	if got := Run(h, "ABC"); got != "abc" {
		t.Errorf("Run = %q, want \"abc\"", got)
	}
}
