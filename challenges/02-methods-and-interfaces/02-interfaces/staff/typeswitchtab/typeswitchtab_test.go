package typeswitchtab

import "testing"

func table() *Table {
	t := NewTable()
	t.Register(0, "int")
	t.Register("", "string")
	t.Register(false, "bool")
	return t
}

func TestDecodeSwitch(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{1, "int"},
		{"x", "string"},
		{true, "bool"},
		{3.5, "unknown"},
		{nil, "unknown"},
		{int64(1), "unknown"},
	}

	for _, tc := range cases {
		if got := DecodeSwitch(tc.in); got != tc.want {
			t.Errorf("DecodeSwitch(%#v) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestDecodeTableMatchesSwitch(t *testing.T) {
	tab := table()
	inputs := []any{1, "x", true, 3.5, nil, int64(1), []int{1}}

	for _, in := range inputs {
		a := DecodeSwitch(in)
		b := tab.DecodeTable(in)
		if a != b {
			t.Errorf("%#v: switch = %q, table = %q", in, a, b)
		}
	}
}

func TestRegistrationOrderIrrelevant(t *testing.T) {
	a := NewTable()
	a.Register(0, "int")
	a.Register("", "string")

	b := NewTable()
	b.Register("", "string")
	b.Register(0, "int")

	for _, in := range []any{1, "x"} {
		if a.DecodeTable(in) != b.DecodeTable(in) {
			t.Errorf("%#v: %q vs %q", in, a.DecodeTable(in), b.DecodeTable(in))
		}
	}
}

func TestRegisterNilIsIgnored(t *testing.T) {
	tab := NewTable()
	tab.Register(nil, "nope")
	if got := tab.DecodeTable(nil); got != "unknown" {
		t.Errorf("DecodeTable(nil) = %q, want \"unknown\"", got)
	}
}

func TestNamedTypesAreDistinct(t *testing.T) {
	type MyInt int

	tab := table()
	if got := tab.DecodeTable(MyInt(1)); got != "unknown" {
		t.Errorf("DecodeTable(MyInt) = %q, want \"unknown\" (a named type is not int)", got)
	}

	tab.Register(MyInt(0), "myint")
	if got := tab.DecodeTable(MyInt(1)); got != "myint" {
		t.Errorf("DecodeTable(MyInt) = %q, want \"myint\"", got)
	}
	if got := tab.DecodeTable(1); got != "int" {
		t.Errorf("DecodeTable(int) = %q, want \"int\"", got)
	}
}

func TestSwitchDoesNotAllocate(t *testing.T) {
	var v any = 42
	if avg := testing.AllocsPerRun(1000, func() { _ = DecodeSwitch(v) }); avg > 0 {
		t.Errorf("DecodeSwitch allocated %.2f times per call, want 0", avg)
	}
}

func BenchmarkDecodeSwitch(b *testing.B) {
	var v any = "x"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = DecodeSwitch(v)
	}
}

func BenchmarkDecodeTable(b *testing.B) {
	tab := table()
	var v any = "x"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = tab.DecodeTable(v)
	}
}
