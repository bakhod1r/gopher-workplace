package splitviews

import (
	"bytes"
	"testing"
)

var sink [][]byte

func TestFields(t *testing.T) {
	got := Fields(nil, []byte("a,bb,c"), ',')
	want := [][]byte{[]byte("a"), []byte("bb"), []byte("c")}
	if len(got) != len(want) {
		t.Fatalf("got %d fields, want %d", len(got), len(want))
	}
	for i := range want {
		if !bytes.Equal(got[i], want[i]) {
			t.Errorf("field %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFieldsEdges(t *testing.T) {
	if got := Fields(nil, nil, ','); len(got) != 0 {
		t.Errorf("Fields(nil) = %q, want empty", got)
	}
	got := Fields(nil, []byte("a,,b"), ',')
	if len(got) != 3 || len(got[1]) != 0 {
		t.Errorf("Fields = %q, want three fields with an empty middle", got)
	}
	got = Fields(nil, []byte(",x"), ',')
	if len(got) != 2 || len(got[0]) != 0 || !bytes.Equal(got[1], []byte("x")) {
		t.Errorf("Fields = %q, want an empty field then x", got)
	}
	got = Fields(nil, []byte("x,"), ',')
	if len(got) != 2 || len(got[1]) != 0 {
		t.Errorf("Fields = %q, want x then an empty field", got)
	}
}

func TestFieldsAppendsToDst(t *testing.T) {
	dst := [][]byte{[]byte("keep")}
	got := Fields(dst, []byte("a"), ',')
	if len(got) != 2 || !bytes.Equal(got[0], []byte("keep")) {
		t.Errorf("Fields = %q, want [keep a]", got)
	}
}

func TestFieldsAreViews(t *testing.T) {
	line := []byte("ab,cd")
	got := Fields(nil, line, ',')
	got[0][0] = 'X'
	if line[0] != 'X' {
		t.Error("the fields copied the bytes; they must be views into line")
	}
}

func TestFieldsCapAtTheBoundary(t *testing.T) {
	line := []byte("ab,cd")
	got := Fields(nil, line, ',')
	got[0] = append(got[0], 'Z')
	if line[2] == 'Z' {
		t.Error("appending to a field wrote into the next one: cap each view at its own end")
	}
	if string(line) != "ab,cd" {
		t.Errorf("line = %q, want \"ab,cd\"", line)
	}
}

func TestFieldsSteadyStateAllocatesNothing(t *testing.T) {
	line := bytes.Repeat([]byte("field,"), 32)
	line = line[:len(line)-1]
	dst := make([][]byte, 0, 64)
	Fields(dst[:0], line, ',')
	if n := testing.AllocsPerRun(100, func() { sink = Fields(dst[:0], line, ',') }); n != 0 {
		t.Errorf("Fields made %v allocations, want 0 when dst has room", n)
	}
}
