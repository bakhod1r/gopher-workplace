package flushwriter

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestWriteAll(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteAll(&buf, []string{"a", "b"}); err != nil {
		t.Fatal(err)
	}
	if got := buf.String(); got != "a\nb\n" {
		t.Errorf("buf = %q, want \"a\\nb\\n\"", got)
	}
}

func TestWriteAllEmpty(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteAll(&buf, nil); err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 0 {
		t.Errorf("buf = %q, want empty", buf.String())
	}
}

func TestWriteAllLargeOutput(t *testing.T) {
	var buf bytes.Buffer
	lines := make([]string, 500)
	for i := range lines {
		lines[i] = strings.Repeat("x", 40)
	}
	if err := WriteAll(&buf, lines); err != nil {
		t.Fatal(err)
	}
	if got := buf.Len(); got != 500*41 {
		t.Errorf("wrote %d bytes, want %d: the last buffer was never flushed", got, 500*41)
	}
}

type failWriter struct{}

func (failWriter) Write(p []byte) (int, error) { return 0, errors.New("boom") }

func TestWriteAllPropagatesErrors(t *testing.T) {
	lines := make([]string, 500)
	for i := range lines {
		lines[i] = strings.Repeat("y", 100)
	}
	if err := WriteAll(failWriter{}, lines); err == nil {
		t.Error("want the writer's error, got nil")
	}
}
