package chunkreader

import "testing"

func TestChunkSourceRead(t *testing.T) {
	s := &ChunkSource{Data: "abcd"}
	buf := make([]byte, 3)

	n, ok := s.Read(buf)
	if n != 3 || !ok || string(buf[:n]) != "abc" {
		t.Errorf("Read 1 = %d, %v, %q", n, ok, buf[:n])
	}
	n, ok = s.Read(buf)
	if n != 1 || !ok || string(buf[:n]) != "d" {
		t.Errorf("Read 2 = %d, %v, %q", n, ok, buf[:n])
	}
	if n, ok := s.Read(buf); n != 0 || ok {
		t.Errorf("Read 3 = %d, %v, want 0, false", n, ok)
	}
}

func TestCountLines(t *testing.T) {
	cases := []struct {
		name    string
		data    string
		bufSize int
		want    int
	}{
		{"two_lines", "a\nb\n", 4, 2},
		{"tiny_buffer", "a\nb\n", 1, 2},
		{"split_boundary", "aaa\nbbb\n", 2, 2},
		{"no_newline", "abc", 4, 0},
		{"empty", "", 4, 0},
		{"only_newlines", "\n\n\n", 2, 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountLines(&ChunkSource{Data: tc.data}, make([]byte, tc.bufSize))
			if got != tc.want {
				t.Errorf("CountLines = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestNoAllocationsPerChunk(t *testing.T) {
	data := ""
	for i := 0; i < 10000; i++ {
		data += "line\n"
	}
	buf := make([]byte, 512)

	avg := testing.AllocsPerRun(5, func() {
		CountLines(&ChunkSource{Data: data}, buf)
	})
	if avg > 4 {
		t.Errorf("CountLines allocated %.0f times; the buffer must be reused", avg)
	}
}

func BenchmarkCountLines(b *testing.B) {
	data := ""
	for i := 0; i < 1000; i++ {
		data += "line\n"
	}
	buf := make([]byte, 4096)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		CountLines(&ChunkSource{Data: data}, buf)
	}
}
