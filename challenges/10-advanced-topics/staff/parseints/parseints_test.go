package parseints

import (
	"bytes"
	"errors"
	"testing"
)

var (
	sinkT int64
	sinkC int
)

func TestParseInts(t *testing.T) {
	if total, n, err := ParseInts([]byte("1,2,3"), ','); err != nil || total != 6 || n != 3 {
		t.Errorf("ParseInts = %d, %d, %v, want 6, 3, nil", total, n, err)
	}
	if total, n, err := ParseInts([]byte("-4,+6"), ','); err != nil || total != 2 || n != 2 {
		t.Errorf("ParseInts = %d, %d, %v, want 2, 2, nil", total, n, err)
	}
	if total, n, err := ParseInts(nil, ','); err != nil || total != 0 || n != 0 {
		t.Errorf("ParseInts = %d, %d, %v, want 0, 0, nil", total, n, err)
	}
	if _, _, err := ParseInts([]byte("1,,2"), ','); !errors.Is(err, ErrSyntax) {
		t.Errorf("err = %v, want ErrSyntax", err)
	}
	if _, _, err := ParseInts([]byte("1,x"), ','); !errors.Is(err, ErrSyntax) {
		t.Errorf("err = %v, want ErrSyntax", err)
	}
	if _, _, err := ParseInts([]byte("-"), ','); !errors.Is(err, ErrSyntax) {
		t.Errorf("err = %v, want ErrSyntax", err)
	}
}

func TestParseIntsSingleField(t *testing.T) {
	if total, n, err := ParseInts([]byte("42"), ','); err != nil || total != 42 || n != 1 {
		t.Errorf("ParseInts = %d, %d, %v, want 42, 1, nil", total, n, err)
	}
}

func TestParseIntsAllocatesNothing(t *testing.T) {
	line := bytes.Repeat([]byte("12345,"), 200)
	line = line[:len(line)-1]
	n := testing.AllocsPerRun(100, func() {
		sinkT, sinkC, _ = ParseInts(line, ',')
	})
	if n != 0 {
		t.Errorf("ParseInts made %v allocations, want 0: parse the bytes in place", n)
	}
}

func BenchmarkParseInts(b *testing.B) {
	line := bytes.Repeat([]byte("12345,"), 200)
	line = line[:len(line)-1]
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkT, sinkC, _ = ParseInts(line, ',')
	}
}
