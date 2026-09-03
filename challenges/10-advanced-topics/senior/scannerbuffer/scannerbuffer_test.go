package scannerbuffer

import (
	"errors"
	"strings"
	"testing"
)

func TestLongestLineShort(t *testing.T) {
	got, err := LongestLine(strings.NewReader("ab\ncdef\ng"))
	if err != nil || got != 4 {
		t.Errorf("LongestLine = %d, %v, want 4, nil", got, err)
	}
}

func TestLongestLineEmpty(t *testing.T) {
	got, err := LongestLine(strings.NewReader(""))
	if err != nil || got != 0 {
		t.Errorf("LongestLine = %d, %v, want 0, nil", got, err)
	}
}

func TestLongestLineOverTheDefaultLimit(t *testing.T) {
	long := strings.Repeat("x", 200*1024)
	got, err := LongestLine(strings.NewReader("short\n" + long + "\n"))
	if err != nil {
		t.Fatalf("LongestLine returned %v, want nil: the scanner's buffer limit was never raised", err)
	}
	if got != len(long) {
		t.Errorf("LongestLine = %d, want %d", got, len(long))
	}
}

func TestLongestLineAtTheConfiguredMax(t *testing.T) {
	long := strings.Repeat("y", 1<<20)
	got, err := LongestLine(strings.NewReader(long))
	if err != nil {
		t.Fatalf("LongestLine returned %v, want nil", err)
	}
	if got != len(long) {
		t.Errorf("LongestLine = %d, want %d", got, len(long))
	}
}

type boom struct{}

func (boom) Read(p []byte) (int, error) { return 0, errors.New("boom") }

func TestLongestLinePropagatesErrors(t *testing.T) {
	if _, err := LongestLine(boom{}); err == nil {
		t.Error("want the reader's error, got nil")
	}
}
