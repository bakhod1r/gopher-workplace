package growonce

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestCollect(t *testing.T) {
	got, err := Collect(strings.NewReader("abc"), 3)
	if err != nil || !bytes.Equal(got, []byte("abc")) {
		t.Errorf("Collect = %q, %v, want \"abc\", nil", got, err)
	}
	if got, err := Collect(strings.NewReader(""), 0); err != nil || len(got) != 0 {
		t.Errorf("Collect = %q, %v, want empty, nil", got, err)
	}
}

func TestCollectHandlesAWrongHint(t *testing.T) {
	in := strings.Repeat("x", 5000)
	got, err := Collect(strings.NewReader(in), 4)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != in {
		t.Errorf("len = %d, want %d: an underestimate must still read everything", len(got), len(in))
	}
	got, err = Collect(strings.NewReader("ab"), 9999)
	if err != nil || string(got) != "ab" {
		t.Errorf("Collect = %q, %v, want \"ab\", nil", got, err)
	}
}

type badReader struct{}

func (badReader) Read(p []byte) (int, error) { return 0, errors.New("boom") }

func TestCollectPropagatesErrors(t *testing.T) {
	if _, err := Collect(badReader{}, 4); err == nil {
		t.Error("want an error, got nil")
	}
}

func TestCollectWithAGoodHintAllocatesOnce(t *testing.T) {
	data := strings.Repeat("y", 8192)
	n := testing.AllocsPerRun(50, func() {
		_, _ = Collect(strings.NewReader(data), len(data))
	})
	if n > 2 {
		t.Errorf("Collect made %v allocations for an exact hint, want at most 2", n)
	}
}

func TestCollectDoesNotLoseBytesAtTheEOF(t *testing.T) {
	got, err := Collect(io.MultiReader(strings.NewReader("ab"), strings.NewReader("cd")), 2)
	if err != nil || string(got) != "abcd" {
		t.Errorf("Collect = %q, %v, want \"abcd\", nil", got, err)
	}
}
