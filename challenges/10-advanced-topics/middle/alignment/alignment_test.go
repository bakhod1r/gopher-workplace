package alignment

import (
	"testing"
	"unsafe"
)

func TestAlignedOnAWideSlice(t *testing.T) {
	u := make([]uint64, 4)
	b := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), 32)
	if !Aligned(b, 8) {
		t.Error("a []uint64's storage must be 8-byte aligned")
	}
	if !Aligned(b, 1) {
		t.Error("everything is 1-byte aligned")
	}
}

func TestAlignedDetectsAnOffset(t *testing.T) {
	u := make([]uint64, 4)
	b := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), 32)
	if Aligned(b[1:], 8) {
		t.Error("a slice starting one byte in cannot be 8-byte aligned")
	}
	if !Aligned(b[8:], 8) {
		t.Error("eight bytes further along is aligned again")
	}
}

func TestAlignedRejectsBadInput(t *testing.T) {
	b := make([]byte, 8)
	if Aligned(nil, 8) {
		t.Error("Aligned(nil) = true, want false")
	}
	if Aligned(b, 0) {
		t.Error("Aligned(b, 0) = true, want false")
	}
	if Aligned(b, 3) {
		t.Error("Aligned(b, 3) = true, want false: n must be a power of two")
	}
	if Aligned(b, 6) {
		t.Error("Aligned(b, 6) = true, want false")
	}
}
