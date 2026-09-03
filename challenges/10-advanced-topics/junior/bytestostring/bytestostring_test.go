package bytestostring

import (
	"testing"
	"unsafe"
)

var sink string

func TestStr(t *testing.T) {
	if got := Str([]byte("hello")); got != "hello" {
		t.Errorf("Str = %q, want \"hello\"", got)
	}
	if got := Str(nil); got != "" {
		t.Errorf("Str(nil) = %q, want empty", got)
	}
	if got := Str([]byte{}); got != "" {
		t.Errorf("Str([]) = %q, want empty", got)
	}
}

func TestStrSharesTheBytes(t *testing.T) {
	b := []byte("abc")
	s := Str(b)
	if unsafe.StringData(s) != unsafe.SliceData(b) {
		t.Error("Str copied the bytes; it must share them")
	}
}

func TestStrDoesNotAllocate(t *testing.T) {
	b := make([]byte, 4096)
	for i := range b {
		b[i] = 'x'
	}
	if n := testing.AllocsPerRun(200, func() { sink = Str(b) }); n != 0 {
		t.Errorf("Str made %v allocations, want 0", n)
	}
}
