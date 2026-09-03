package stringtobytes

import (
	"bytes"
	"testing"
	"unsafe"
)

var sink []byte

func TestBytes(t *testing.T) {
	if got := Bytes("hello"); !bytes.Equal(got, []byte("hello")) {
		t.Errorf("Bytes = %q, want \"hello\"", got)
	}
	if got := Bytes(""); len(got) != 0 {
		t.Errorf("Bytes(\"\") = %q, want empty", got)
	}
}

func TestBytesSharesTheString(t *testing.T) {
	s := "shared"
	b := Bytes(s)
	if unsafe.SliceData(b) != unsafe.StringData(s) {
		t.Error("Bytes copied the string; it must share it")
	}
}

func TestBytesLengthAndCapacity(t *testing.T) {
	b := Bytes("abcd")
	if len(b) != 4 {
		t.Errorf("len = %d, want 4", len(b))
	}
	if cap(b) != 4 {
		t.Errorf("cap = %d, want 4: an append must not write past the string", cap(b))
	}
}

func TestBytesDoesNotAllocate(t *testing.T) {
	s := string(make([]byte, 4096))
	if n := testing.AllocsPerRun(200, func() { sink = Bytes(s) }); n != 0 {
		t.Errorf("Bytes made %v allocations, want 0", n)
	}
}
