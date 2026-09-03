package outparam

import (
	"bytes"
	"testing"
)

func TestFill(t *testing.T) {
	buf := make([]byte, 3)
	if n := Fill(buf, 'x'); n != 3 {
		t.Errorf("Fill = %d, want 3", n)
	}
	if !bytes.Equal(buf, []byte("xxx")) {
		t.Errorf("buf = %q, want \"xxx\"", buf)
	}
	if n := Fill(nil, 'x'); n != 0 {
		t.Errorf("Fill(nil) = %d, want 0", n)
	}
}

func TestFillWritesOnlyTheView(t *testing.T) {
	buf := []byte("abcd")
	Fill(buf[1:3], 'z')
	if !bytes.Equal(buf, []byte("azzd")) {
		t.Errorf("buf = %q, want \"azzd\"", buf)
	}
}

func TestFillAllocatesNothing(t *testing.T) {
	buf := make([]byte, 256)
	if n := testing.AllocsPerRun(100, func() { Fill(buf, 1) }); n != 0 {
		t.Errorf("Fill made %v allocations, want 0", n)
	}
}
