package scratchbuf

import "testing"

var sink []byte

func TestAppendJoin(t *testing.T) {
	cases := []struct {
		parts []string
		sep   string
		want  string
	}{
		{[]string{"a", "b"}, "/", "a/b"},
		{[]string{"a"}, "/", "a"},
		{nil, "/", ""},
		{[]string{"a", "", "b"}, "-", "a--b"},
	}
	for _, c := range cases {
		if got := string(AppendJoin(nil, c.parts, c.sep)); got != c.want {
			t.Errorf("AppendJoin(%q, %q) = %q, want %q", c.parts, c.sep, got, c.want)
		}
	}
}

func TestAppendJoinAppendsToExistingContent(t *testing.T) {
	got := string(AppendJoin([]byte("prefix:"), []string{"a", "b"}, "/"))
	if got != "prefix:a/b" {
		t.Errorf("AppendJoin = %q, want %q", got, "prefix:a/b")
	}
}

func TestAppendJoinReusesScratch(t *testing.T) {
	scratch := make([]byte, 0, 256)
	parts := []string{"alpha", "beta", "gamma", "delta"}
	allocs := testing.AllocsPerRun(100, func() { sink = AppendJoin(scratch[:0], parts, "/") })
	if allocs != 0 {
		t.Errorf("AppendJoin into a big enough scratch made %v allocations, want 0", allocs)
	}
}

func TestSized(t *testing.T) {
	cases := []struct {
		parts []string
		sep   string
		want  int
	}{
		{[]string{"a", "b"}, "/", 3},
		{[]string{"a"}, "/", 1},
		{nil, "/", 0},
		{[]string{"ab", "cd", "ef"}, "--", 10},
	}
	for _, c := range cases {
		if got := Sized(c.parts, c.sep); got != c.want {
			t.Errorf("Sized(%q, %q) = %d, want %d", c.parts, c.sep, got, c.want)
		}
	}
}

func TestSizedMatchesAppendJoin(t *testing.T) {
	parts := []string{"one", "two", "three"}
	sep := ", "
	got := AppendJoin(nil, parts, sep)
	if len(got) != Sized(parts, sep) {
		t.Errorf("Sized = %d, AppendJoin wrote %d bytes", Sized(parts, sep), len(got))
	}
	// A scratch buffer sized by Sized is enough for exactly one join.
	scratch := make([]byte, 0, Sized(parts, sep))
	if allocs := testing.AllocsPerRun(50, func() { sink = AppendJoin(scratch[:0], parts, sep) }); allocs != 0 {
		t.Errorf("a Sized-sized scratch still reallocated: %v allocations", allocs)
	}
}
