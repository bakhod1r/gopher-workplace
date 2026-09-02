package unsafeslice

import "testing"

func TestBytesToString(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
		want string
	}{
		{"basic", []byte("abc"), "abc"},
		{"empty", []byte{}, ""},
		{"nil", nil, ""},
		{"binary", []byte{0, 1, 2}, string([]byte{0, 1, 2})},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := BytesToString(tc.in); got != tc.want {
				t.Errorf("BytesToString = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestBytesToStringAliases(t *testing.T) {
	b := []byte("abc")
	s := BytesToString(b)

	b[0] = 'X'
	if s != "Xbc" {
		t.Errorf("s = %q; BytesToString must alias the source, not copy it", s)
	}
}

func TestSafeStringCopies(t *testing.T) {
	b := []byte("abc")
	s := SafeString(b)

	b[0] = 'X'
	if s != "abc" {
		t.Errorf("s = %q; SafeString must copy", s)
	}
}

func TestStringToBytes(t *testing.T) {
	got := StringToBytes("abc")
	if string(got) != "abc" {
		t.Errorf("StringToBytes = %q, want \"abc\"", got)
	}
	if len(got) != 3 {
		t.Errorf("len = %d, want 3", len(got))
	}
	if n := len(StringToBytes("")); n != 0 {
		t.Errorf("StringToBytes(\"\") len = %d, want 0", n)
	}
}

func TestRoundTrip(t *testing.T) {
	orig := []byte("round trip")
	s := BytesToString(orig)
	back := StringToBytes(s)

	if string(back) != "round trip" {
		t.Errorf("round trip = %q", back)
	}
	if len(back) != len(orig) {
		t.Errorf("len = %d, want %d", len(back), len(orig))
	}
}

func TestBytesToStringDoesNotAllocate(t *testing.T) {
	b := []byte("a reasonably long field value")

	if avg := testing.AllocsPerRun(1000, func() { sink = BytesToString(b) }); avg > 0 {
		t.Errorf("BytesToString allocated %.2f times per call, want 0", avg)
	}
}

// sink keeps the converted strings alive so the compiler cannot delete the
// conversion under the allocation measurement.
var sink string

func TestSafeStringDoesAllocate(t *testing.T) {
	b := []byte("a reasonably long field value")

	if avg := testing.AllocsPerRun(1000, func() { sink = SafeString(b) }); avg < 1 {
		t.Errorf("SafeString allocated %.2f times per call; the copy must cost something", avg)
	}
	if sink == "" {
		t.Error("sink was never written")
	}
}

func BenchmarkBytesToString(b *testing.B) {
	data := []byte("a reasonably long field value")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BytesToString(data)
	}
}

func BenchmarkSafeString(b *testing.B) {
	data := []byte("a reasonably long field value")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SafeString(data)
	}
}
