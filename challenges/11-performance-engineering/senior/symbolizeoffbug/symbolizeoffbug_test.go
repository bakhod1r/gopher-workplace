package symbolizeoffbug

import (
	"reflect"
	"testing"
)

var table = []Symbol{
	{100, "a"},
	{200, "b"},
	{300, "c"},
}

func TestResolveInsideAFunction(t *testing.T) {
	cases := []struct {
		addr uint64
		want string
	}{
		{150, "a"},
		{199, "a"},
		{250, "b"},
		{299, "b"},
		{1_000_000, "c"},
	}
	for _, c := range cases {
		got, ok := Resolve(table, c.addr)
		if !ok || got != c.want {
			t.Errorf("Resolve(%d) = %q, %v, want %q, true", c.addr, got, ok, c.want)
		}
	}
}

func TestResolveOnAFunctionEntryPoint(t *testing.T) {
	// An address landing exactly on a symbol's start belongs to that symbol —
	// this is the function's first instruction, the most sampled address in
	// any short function.
	cases := []struct {
		addr uint64
		want string
	}{
		{100, "a"},
		{200, "b"},
		{300, "c"},
	}
	for _, c := range cases {
		got, ok := Resolve(table, c.addr)
		if !ok || got != c.want {
			t.Errorf("Resolve(%d) = %q, %v, want %q, true", c.addr, got, ok, c.want)
		}
	}
}

func TestResolveBelowTheTable(t *testing.T) {
	for _, addr := range []uint64{0, 99} {
		if got, ok := Resolve(table, addr); ok || got != "" {
			t.Errorf("Resolve(%d) = %q, %v, want \"\", false", addr, got, ok)
		}
	}
	if _, ok := Resolve(nil, 100); ok {
		t.Error("Resolve on an empty table reported found")
	}
}

func TestSymbolize(t *testing.T) {
	got := Symbolize(table, []uint64{100, 200, 1, 350})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Symbolize = %v, want [a b c]", got)
	}
}
