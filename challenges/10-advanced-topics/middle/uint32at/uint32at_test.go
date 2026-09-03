package uint32at

import (
	"testing"
	"unsafe"
)

func aligned32(n int) []byte {
	u := make([]uint32, (n+3)/4)
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), n)
}

func TestUint32At(t *testing.T) {
	b := aligned32(8)
	p := (*uint32)(unsafe.Pointer(unsafe.SliceData(b)))
	*p = 0x01020304
	got, ok := Uint32At(b, 0)
	if !ok || got != 0x01020304 {
		t.Errorf("Uint32At = %#x, %v, want 0x01020304, true", got, ok)
	}
}

func TestUint32AtSecondWord(t *testing.T) {
	b := aligned32(8)
	q := (*uint32)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(b)), 4))
	*q = 0xdeadbeef
	got, ok := Uint32At(b, 4)
	if !ok || got != 0xdeadbeef {
		t.Errorf("Uint32At = %#x, %v, want 0xdeadbeef, true", got, ok)
	}
}

func TestUint32AtOutOfRange(t *testing.T) {
	b := aligned32(8)
	for _, off := range []int{-1, 5, 8, 9, 100} {
		if _, ok := Uint32At(b, off); ok {
			t.Errorf("Uint32At(off=%d) reported ok, want false", off)
		}
	}
	if _, ok := Uint32At(nil, 0); ok {
		t.Error("Uint32At(nil) reported ok, want false")
	}
}

func TestUint32AtMisaligned(t *testing.T) {
	b := aligned32(16)
	if _, ok := Uint32At(b[1:], 0); ok {
		t.Error("a misaligned read reported ok, want false")
	}
	if _, ok := Uint32At(b, 2); ok {
		t.Error("an odd offset reported ok, want false")
	}
}
