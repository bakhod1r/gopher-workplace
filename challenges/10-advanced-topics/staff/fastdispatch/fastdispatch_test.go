package fastdispatch

import (
	"bytes"
	"testing"
)

type myInt int
type myString string

var sink []byte

func TestRenderFastPath(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{42, "42"},
		{int64(-7), "-7"},
		{"text", "text"},
		{true, "true"},
		{false, "false"},
		{[]byte("bytes"), "bytes"},
		{nil, "<nil>"},
	}
	for _, c := range cases {
		if got := Render(nil, c.in); !bytes.Equal(got, []byte(c.want)) {
			t.Errorf("Render(%#v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRenderFallback(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{myInt(5), "5"},
		{myString("named"), "named"},
		{uint8(200), "200"},
		{int32(-3), "-3"},
		{uint64(1 << 40), "1099511627776"},
	}
	for _, c := range cases {
		if got := Render(nil, c.in); !bytes.Equal(got, []byte(c.want)) {
			t.Errorf("Render(%#v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRenderUnsupported(t *testing.T) {
	if got := Render(nil, struct{ A int }{}); !bytes.Equal(got, []byte("?")) {
		t.Errorf("Render = %q, want \"?\"", got)
	}
	if got := Render(nil, 1.5); !bytes.Equal(got, []byte("?")) {
		t.Errorf("Render = %q, want \"?\"", got)
	}
}

func TestRenderAppends(t *testing.T) {
	got := Render([]byte("pre:"), 9)
	if !bytes.Equal(got, []byte("pre:9")) {
		t.Errorf("Render = %q, want \"pre:9\"", got)
	}
}

func TestRenderFastPathAllocatesNothing(t *testing.T) {
	dst := make([]byte, 0, 64)
	for _, v := range []any{42, "text", true, int64(9)} {
		v := v
		if n := testing.AllocsPerRun(200, func() { sink = Render(dst[:0], v) }); n != 0 {
			t.Errorf("Render(%#v) made %v allocations, want 0", v, n)
		}
	}
}

func BenchmarkRenderInt(b *testing.B) {
	dst := make([]byte, 0, 64)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = Render(dst[:0], 42)
	}
}
