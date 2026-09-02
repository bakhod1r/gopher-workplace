package whennotgenericgen

import (
	"strconv"
	"testing"
)

type buffer struct {
	lines []string
}

func (b *buffer) Write(line string) { b.lines = append(b.lines, line) }

func TestWriteAll(t *testing.T) {
	b := &buffer{}
	if got := WriteAll(b, []string{"a", "b"}); got != 2 {
		t.Errorf("WriteAll = %d, want 2", got)
	}
	if len(b.lines) != 2 || b.lines[0] != "a" || b.lines[1] != "b" {
		t.Errorf("buffer = %v, want [a b]", b.lines)
	}
	if got := WriteAll(b, nil); got != 0 {
		t.Errorf("WriteAll(nil) = %d, want 0", got)
	}
}

func TestWriteEach(t *testing.T) {
	b := &buffer{}
	if got := WriteEach(b, []int{1, 2}, strconv.Itoa); got != 2 {
		t.Errorf("WriteEach = %d, want 2", got)
	}
	if len(b.lines) != 2 || b.lines[0] != "1" || b.lines[1] != "2" {
		t.Errorf("buffer = %v, want [1 2]", b.lines)
	}
}

func TestWriteEachStructs(t *testing.T) {
	type row struct{ name string }
	b := &buffer{}
	WriteEach(b, []row{{"x"}}, func(r row) string { return r.name })
	if len(b.lines) != 1 || b.lines[0] != "x" {
		t.Errorf("buffer = %v, want [x]", b.lines)
	}
}
