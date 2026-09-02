package poolencoder

import "testing"

func TestEncode(t *testing.T) {
	e := NewPooledEncoder()
	cases := []struct {
		name   string
		fields []string
		want   string
	}{
		{"two", []string{"a", "b"}, "a,b"},
		{"one", []string{"solo"}, "solo"},
		{"empty", nil, ""},
		{"empty_fields", []string{"", ""}, ","},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := e.Encode(tc.fields); got != tc.want {
				t.Errorf("Encode = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestResultDoesNotAliasPool(t *testing.T) {
	e := NewPooledEncoder()
	first := e.Encode([]string{"aaaa", "bbbb"})
	// The next call reuses the same pooled buffer.
	e.Encode([]string{"cccc", "dddd"})
	if first != "aaaa,bbbb" {
		t.Errorf("earlier result was corrupted by buffer reuse: %q", first)
	}
}

func TestBufferIsResetBetweenCalls(t *testing.T) {
	e := NewPooledEncoder()
	e.Encode([]string{"xxxxxxxx"})
	if got := e.Encode([]string{"y"}); got != "y" {
		t.Errorf("Encode = %q, want \"y\" (stale buffer contents leaked)", got)
	}
}

func TestEncodeAll(t *testing.T) {
	e := NewPooledEncoder()
	got := EncodeAll(e, [][]string{{"a", "b"}, {"c"}})
	if len(got) != 2 || got[0] != "a,b" || got[1] != "c" {
		t.Errorf("EncodeAll = %v", got)
	}
	if n := len(EncodeAll(e, nil)); n != 0 {
		t.Errorf("EncodeAll(nil) len = %d, want 0", n)
	}
}

func TestBoundedAllocations(t *testing.T) {
	e := NewPooledEncoder()
	fields := []string{"alpha", "beta", "gamma"}
	e.Encode(fields) // warm the pool

	avg := testing.AllocsPerRun(1000, func() {
		_ = e.Encode(fields)
	})
	if avg > 2 {
		t.Errorf("Encode allocated %.2f times per call; the buffer must be reused", avg)
	}
}

func BenchmarkEncode(b *testing.B) {
	e := NewPooledEncoder()
	fields := []string{"alpha", "beta", "gamma"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = e.Encode(fields)
	}
}
