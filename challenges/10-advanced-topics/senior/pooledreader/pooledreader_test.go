package pooledreader

import (
	"errors"
	"strings"
	"testing"
)

func TestFirstLine(t *testing.T) {
	got, err := FirstLine(strings.NewReader("alpha\nbeta\n"))
	if err != nil || got != "alpha" {
		t.Errorf("FirstLine = %q, %v, want \"alpha\", nil", got, err)
	}
}

func TestFirstLineNoTrailingNewline(t *testing.T) {
	got, err := FirstLine(strings.NewReader("only"))
	if err != nil || got != "only" {
		t.Errorf("FirstLine = %q, %v, want \"only\", nil", got, err)
	}
}

func TestFirstLineEmpty(t *testing.T) {
	got, err := FirstLine(strings.NewReader(""))
	if err != nil || got != "" {
		t.Errorf("FirstLine = %q, %v, want empty, nil", got, err)
	}
}

func TestFirstLineAcrossManyRequests(t *testing.T) {
	for i := 0; i < 100; i++ {
		want := strings.Repeat("x", i%7+1)
		got, err := FirstLine(strings.NewReader(want + "\ntail\n"))
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Fatalf("request %d: FirstLine = %q, want %q: the pooled reader kept the previous source", i, got, want)
		}
	}
}

type boom struct{}

func (boom) Read(p []byte) (int, error) { return 0, errors.New("boom") }

func TestFirstLinePropagatesErrors(t *testing.T) {
	if _, err := FirstLine(boom{}); err == nil {
		t.Error("want the reader's error, got nil")
	}
}
