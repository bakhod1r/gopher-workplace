package splitcount

import (
	"bytes"
	"testing"
)

var (
	sinkF int
	sinkS int
)

func TestCountFields(t *testing.T) {
	cases := []struct {
		in           string
		fields, size int
	}{
		{"ab,c", 2, 3},
		{"", 0, 0},
		{"abc", 1, 3},
		{",", 2, 0},
		{"a,,b", 3, 2},
		{",a", 2, 1},
	}
	for _, c := range cases {
		f, s := CountFields([]byte(c.in), ',')
		if f != c.fields || s != c.size {
			t.Errorf("CountFields(%q) = %d, %d, want %d, %d", c.in, f, s, c.fields, c.size)
		}
	}
}

func TestCountFieldsAllocatesNothing(t *testing.T) {
	line := bytes.Repeat([]byte("column,"), 256)
	line = line[:len(line)-1]
	n := testing.AllocsPerRun(100, func() { sinkF, sinkS = CountFields(line, ',') })
	if n != 0 {
		t.Errorf("CountFields made %v allocations, want 0: scan the bytes, do not split them", n)
	}
}

func BenchmarkCountFields(b *testing.B) {
	line := bytes.Repeat([]byte("column,"), 256)
	line = line[:len(line)-1]
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkF, sinkS = CountFields(line, ',')
	}
}
