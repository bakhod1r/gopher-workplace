package int32view

import (
	"testing"
	"unsafe"
)

func alignedBytes(n int) []byte {
	u := make([]int32, (n+3)/4)
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), n)
}

func TestInt32sShape(t *testing.T) {
	b := alignedBytes(8)
	v, ok := Int32s(b)
	if !ok {
		t.Fatal("Int32s reported false for an aligned 8-byte buffer")
	}
	if len(v) != 2 || cap(v) != 2 {
		t.Errorf("len, cap = %d, %d, want 2, 2", len(v), cap(v))
	}
}

func TestInt32sSharesStorage(t *testing.T) {
	b := alignedBytes(8)
	v, ok := Int32s(b)
	if !ok {
		t.Fatal("Int32s reported false")
	}
	v[0] = 0x01020304
	if b[0] == 0 && b[1] == 0 && b[2] == 0 && b[3] == 0 {
		t.Error("the view does not share the bytes")
	}
	v[1] = -1
	for _, x := range b[4:8] {
		if x != 0xff {
			t.Errorf("b[4:8] = %v, want all 0xff", b[4:8])
			break
		}
	}
}

func TestInt32sRejectsBadShapes(t *testing.T) {
	b := alignedBytes(16)
	for _, c := range []struct {
		name string
		in   []byte
	}{
		{"nil", nil},
		{"empty", b[:0]},
		{"length not a multiple of 4", b[:6]},
		{"misaligned", b[1:13]},
	} {
		if _, ok := Int32s(c.in); ok {
			t.Errorf("%s: reported ok, want false", c.name)
		}
	}
}

func TestInt32sDoesNotAllocate(t *testing.T) {
	b := alignedBytes(4096)
	var sink []int32
	if n := testing.AllocsPerRun(100, func() { sink, _ = Int32s(b) }); n != 0 {
		t.Errorf("Int32s made %v allocations, want 0", n)
	}
	_ = sink
}
