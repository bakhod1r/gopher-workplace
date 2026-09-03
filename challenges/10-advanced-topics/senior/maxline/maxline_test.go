package maxline

import (
	"errors"
	"io"
	"runtime"
	"strings"
	"testing"
)

func TestMaxLine(t *testing.T) {
	if n, err := MaxLine(strings.NewReader("ab\ncdef\ng")); err != nil || n != 4 {
		t.Errorf("MaxLine = %d, %v, want 4, nil", n, err)
	}
	if n, err := MaxLine(strings.NewReader("")); err != nil || n != 0 {
		t.Errorf("MaxLine = %d, %v, want 0, nil", n, err)
	}
	if n, err := MaxLine(strings.NewReader("\n\n")); err != nil || n != 0 {
		t.Errorf("MaxLine = %d, %v, want 0, nil", n, err)
	}
	if n, err := MaxLine(strings.NewReader("tail-without-newline")); err != nil || n != 20 {
		t.Errorf("MaxLine = %d, %v, want 20, nil", n, err)
	}
}

type boom struct{}

func (boom) Read(p []byte) (int, error) { return 0, errors.New("boom") }

func TestMaxLinePropagatesErrors(t *testing.T) {
	if _, err := MaxLine(boom{}); err == nil {
		t.Error("want an error, got nil")
	}
}

// long yields one line of size bytes then EOF.
type long struct{ left int64 }

func (l *long) Read(p []byte) (int, error) {
	if l.left <= 0 {
		return 0, io.EOF
	}
	c := int64(len(p))
	if c > l.left {
		c = l.left
	}
	for i := range p[:c] {
		p[i] = 'x'
	}
	l.left -= c
	return int(c), nil
}

func TestMaxLineDoesNotBufferTheLine(t *testing.T) {
	const size = 32 << 20
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	n, err := MaxLine(&long{left: size})
	if err != nil {
		t.Fatal(err)
	}
	if n != size {
		t.Fatalf("MaxLine = %d, want %d", n, size)
	}
	runtime.ReadMemStats(&after)
	if used := after.TotalAlloc - before.TotalAlloc; used > 1<<20 {
		t.Errorf("allocated %d bytes for one %d byte line, want under 1 MiB: count, do not accumulate", used, size)
	}
}
