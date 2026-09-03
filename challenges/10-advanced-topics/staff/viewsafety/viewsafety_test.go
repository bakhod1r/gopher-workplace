package viewsafety

import (
	"bytes"
	"testing"
	"unsafe"
)

func TestWindow(t *testing.T) {
	b := []byte("abcdef")
	got, ok := Window(b, 2, 3)
	if !ok || !bytes.Equal(got, []byte("cde")) {
		t.Errorf("Window = %q, %v, want \"cde\", true", got, ok)
	}
}

func TestWindowCapacityIsExact(t *testing.T) {
	b := []byte("abcdef")
	got, ok := Window(b, 2, 3)
	if !ok {
		t.Fatal("Window reported false")
	}
	if cap(got) != 3 {
		t.Fatalf("cap = %d, want 3", cap(got))
	}
	got = append(got, 'Z')
	if b[5] == 'Z' {
		t.Error("the append wrote past the window into the caller's buffer")
	}
	if string(b) != "abcdef" {
		t.Errorf("b = %q, want \"abcdef\"", b)
	}
}

func TestWindowSharesUntilItGrows(t *testing.T) {
	b := []byte("abcdef")
	got, _ := Window(b, 0, 3)
	got[0] = 'X'
	if b[0] != 'X' {
		t.Error("the window does not share the buffer")
	}
}

func TestWindowBounds(t *testing.T) {
	b := []byte("abcdef")
	for _, c := range [][2]int{{-1, 2}, {0, -1}, {4, 3}, {7, 0}, {0, 7}} {
		if _, ok := Window(b, c[0], c[1]); ok {
			t.Errorf("Window(off=%d, n=%d) reported ok, want false", c[0], c[1])
		}
	}
}

func TestWindowZeroLength(t *testing.T) {
	b := []byte("abc")
	got, ok := Window(b, 1, 0)
	if !ok || len(got) != 0 {
		t.Errorf("Window = %q, %v, want empty, true", got, ok)
	}
}

func TestWindowAllocatesNothing(t *testing.T) {
	b := make([]byte, 4096)
	var sink []byte
	if n := testing.AllocsPerRun(200, func() { sink, _ = Window(b, 8, 1024) }); n != 0 {
		t.Errorf("Window made %v allocations, want 0", n)
	}
	_ = sink
}

func TestWindowStartsWhereAsked(t *testing.T) {
	b := []byte("abcdef")
	got, _ := Window(b, 2, 2)
	want := (*byte)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(b)), 2))
	if unsafe.SliceData(got) != want {
		t.Error("the window does not start at the requested offset")
	}
}
