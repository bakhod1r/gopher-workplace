package samplesymbolize

import (
	"reflect"
	"testing"
)

var table = []Symbol{
	{100, "a"},
	{200, "b"},
	{300, "c"},
}

func TestResolve(t *testing.T) {
	cases := []struct {
		addr uint64
		want string
		ok   bool
	}{
		{100, "a", true},
		{150, "a", true},
		{199, "a", true},
		{200, "b", true},
		{299, "b", true},
		{300, "c", true},
		{1_000_000, "c", true},
		{99, "", false},
		{0, "", false},
	}
	for _, c := range cases {
		got, ok := Resolve(table, c.addr)
		if got != c.want || ok != c.ok {
			t.Errorf("Resolve(%d) = %q, %v, want %q, %v", c.addr, got, ok, c.want, c.ok)
		}
	}
}

func TestResolveEmptyTable(t *testing.T) {
	if got, ok := Resolve(nil, 100); got != "" || ok {
		t.Errorf("Resolve(nil, 100) = %q, %v, want \"\", false", got, ok)
	}
}

func TestSymbolize(t *testing.T) {
	got := Symbolize(table, []uint64{150, 250, 1, 350})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Symbolize = %v, want [a b c]", got)
	}
	got = Symbolize(table, []uint64{1, 2})
	if got == nil || len(got) != 0 {
		t.Errorf("Symbolize = %v, want empty non-nil slice", got)
	}
}

func TestResolveScalesToALargeTable(t *testing.T) {
	big := make([]Symbol, 0, 50_000)
	for i := 0; i < 50_000; i++ {
		big = append(big, Symbol{Start: uint64(i) * 16, Func: string(rune('a' + i%26))})
	}
	for _, addr := range []uint64{0, 15, 16, 799_984, 799_999, 10_000_000} {
		if _, ok := Resolve(big, addr); !ok {
			t.Errorf("Resolve(%d) reported not found", addr)
		}
	}
	if got, _ := Resolve(big, 16*3+5); got != big[3].Func {
		t.Errorf("Resolve = %q, want %q", got, big[3].Func)
	}
}
