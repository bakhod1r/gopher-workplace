package joinvsloop

import (
	"reflect"
	"strings"
	"testing"
)

var sink string

func TestJoinPath(t *testing.T) {
	cases := []struct {
		parts []string
		want  string
	}{
		{[]string{"a", "b"}, "a/b"},
		{[]string{"a"}, "a"},
		{nil, ""},
		{[]string{"a", "", "b"}, "a//b"},
	}
	for _, c := range cases {
		if got := JoinPath(c.parts); got != c.want {
			t.Errorf("JoinPath(%q) = %q, want %q", c.parts, got, c.want)
		}
	}
}

func TestJoinPathAllocatesOnce(t *testing.T) {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "segment"
	}
	allocs := testing.AllocsPerRun(50, func() { sink = JoinPath(parts) })
	if allocs > 1 {
		t.Errorf("JoinPath made %v allocations, want at most 1 — a += loop allocates once per segment", allocs)
	}
}

func TestSplitPath(t *testing.T) {
	if got := SplitPath("a/b"); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("SplitPath = %v, want [a b]", got)
	}
	if got := SplitPath("a"); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("SplitPath = %v, want [a]", got)
	}
}

func TestSplitPathEmpty(t *testing.T) {
	got := SplitPath("")
	if got == nil || len(got) != 0 {
		t.Errorf("SplitPath(\"\") = %v, want empty non-nil slice", got)
	}
}

func TestRoundTrip(t *testing.T) {
	parts := []string{"usr", "local", "bin"}
	if got := SplitPath(JoinPath(parts)); !reflect.DeepEqual(got, parts) {
		t.Errorf("round trip = %v, want %v", got, parts)
	}
	if got := JoinPath(SplitPath("a/b/c")); got != "a/b/c" {
		t.Errorf("round trip = %q, want %q", got, "a/b/c")
	}
	if !strings.Contains(JoinPath(parts), "/") {
		t.Error("JoinPath produced no separator")
	}
}
