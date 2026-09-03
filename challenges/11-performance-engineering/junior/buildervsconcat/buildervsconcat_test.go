package buildervsconcat

import (
	"strings"
	"testing"
)

var sink string

func TestJoinSep(t *testing.T) {
	cases := []struct {
		parts []string
		sep   string
		want  string
	}{
		{[]string{"a", "b"}, ", ", "a, b"},
		{[]string{"a"}, ", ", "a"},
		{nil, ", ", ""},
		{[]string{"a", "b", "c"}, "", "abc"},
		{[]string{"", ""}, "-", "-"},
	}
	for _, c := range cases {
		if got := JoinSep(c.parts, c.sep); got != c.want {
			t.Errorf("JoinSep(%q, %q) = %q, want %q", c.parts, c.sep, got, c.want)
		}
	}
}

func TestJoinSepAllocatesOnce(t *testing.T) {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "abcdefgh"
	}
	allocs := testing.AllocsPerRun(50, func() { sink = JoinSep(parts, ", ") })
	if allocs > 1 {
		t.Errorf("JoinSep made %v allocations, want at most 1 — the final size is computable", allocs)
	}
}

func TestJoinSepMatchesStdlib(t *testing.T) {
	parts := []string{"go", "is", "fast"}
	if got, want := JoinSep(parts, "/"), strings.Join(parts, "/"); got != want {
		t.Errorf("JoinSep = %q, want %q", got, want)
	}
}
