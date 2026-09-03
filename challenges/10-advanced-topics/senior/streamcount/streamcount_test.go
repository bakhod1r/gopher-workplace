package streamcount

import (
	"errors"
	"io"
	"runtime"
	"strings"
	"testing"
)

func TestCountLines(t *testing.T) {
	n, err := CountLines(strings.NewReader("a\nb\nc"))
	if err != nil || n != 2 {
		t.Errorf("CountLines = %d, %v, want 2, nil", n, err)
	}
	if n, err := CountLines(strings.NewReader("")); err != nil || n != 0 {
		t.Errorf("CountLines = %d, %v, want 0, nil", n, err)
	}
	if n, err := CountLines(strings.NewReader("\n\n\n")); err != nil || n != 3 {
		t.Errorf("CountLines = %d, %v, want 3, nil", n, err)
	}
}

type errReader struct{ n int }

func (e *errReader) Read(p []byte) (int, error) {
	if e.n == 0 {
		return 0, errors.New("boom")
	}
	e.n--
	p[0] = '\n'
	return 1, nil
}

func TestCountLinesPropagatesErrors(t *testing.T) {
	if _, err := CountLines(&errReader{n: 2}); err == nil {
		t.Error("want the reader's error, got nil")
	}
}

// zeros yields n bytes, one '\n' every 1024, without allocating.
type zeros struct{ left int64 }

func (z *zeros) Read(p []byte) (int, error) {
	if z.left <= 0 {
		return 0, io.EOF
	}
	c := int64(len(p))
	if c > z.left {
		c = z.left
	}
	for i := range p[:c] {
		if (z.left-int64(i))%1024 == 0 {
			p[i] = '\n'
		} else {
			p[i] = 'x'
		}
	}
	z.left -= c
	return int(c), nil
}

func TestCountLinesStaysUnderTheMemoryCeiling(t *testing.T) {
	const size = 64 << 20
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	if _, err := CountLines(&zeros{left: size}); err != nil {
		t.Fatal(err)
	}
	runtime.ReadMemStats(&after)
	if used := after.TotalAlloc - before.TotalAlloc; used > 1<<20 {
		t.Errorf("allocated %d bytes for a %d byte stream, want under 1 MiB: do not buffer the input", used, size)
	}
}
