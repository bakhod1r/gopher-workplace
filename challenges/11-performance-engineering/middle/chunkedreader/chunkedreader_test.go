package chunkedreader

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	got, err := ReadAll(strings.NewReader("hello"), nil, 2)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if string(got) != "hello" {
		t.Errorf("ReadAll = %q, want %q", got, "hello")
	}
}

func TestReadAllEmptyReader(t *testing.T) {
	got, err := ReadAll(strings.NewReader(""), nil, 4)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if len(got) != 0 {
		t.Errorf("ReadAll = %q, want empty", got)
	}
}

func TestReadAllReusesTheBuffer(t *testing.T) {
	buf := make([]byte, 0, 1024)
	got, err := ReadAll(strings.NewReader("hello world"), buf, 4)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if string(got) != "hello world" {
		t.Errorf("ReadAll = %q, want %q", got, "hello world")
	}
	if cap(got) != 1024 {
		t.Errorf("cap = %d, want the caller's 1024 — the buffer must be reused", cap(got))
	}
}

func TestReadAllResetsTheBuffer(t *testing.T) {
	buf := []byte("stale")
	got, _ := ReadAll(strings.NewReader("new"), buf, 4)
	if string(got) != "new" {
		t.Errorf("ReadAll = %q, want %q", got, "new")
	}
}

type errReader struct{ n int }

func (e *errReader) Read(p []byte) (int, error) {
	if e.n > 0 {
		e.n--
		p[0] = 'x'
		return 1, nil
	}
	return 0, errors.New("boom")
}

func TestReadAllReturnsWhatItReadOnError(t *testing.T) {
	got, err := ReadAll(&errReader{n: 3}, nil, 4)
	if err == nil {
		t.Fatal("err = nil, want the reader's error")
	}
	if errors.Is(err, io.EOF) {
		t.Fatal("EOF must not be reported as an error")
	}
	if string(got) != "xxx" {
		t.Errorf("ReadAll = %q, want %q — data read before the error is not lost", got, "xxx")
	}
}

func TestCountChunks(t *testing.T) {
	cases := []struct{ n, chunk, want int }{
		{5, 2, 4}, // 2 + 2 + 1, then the EOF read
		{4, 2, 3}, // 2 + 2, then EOF
		{1, 4, 2}, // 1, then EOF
		{0, 4, 0},
		{5, 0, 0},
		{-1, 2, 0},
	}
	for _, c := range cases {
		if got := CountChunks(c.n, c.chunk); got != c.want {
			t.Errorf("CountChunks(%d, %d) = %d, want %d", c.n, c.chunk, got, c.want)
		}
	}
}
