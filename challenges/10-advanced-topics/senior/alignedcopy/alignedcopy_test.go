package alignedcopy

import (
	"testing"
	"unsafe"
)

func wordBytes(vals []uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(vals))), len(vals)*4)
}

func TestCopyWords(t *testing.T) {
	src := wordBytes([]uint32{1, 2, 3})
	dst := make([]uint32, 3)
	n, ok := CopyWords(dst, src)
	if !ok || n != 3 {
		t.Fatalf("CopyWords = %d, %v, want 3, true", n, ok)
	}
	if dst[0] != 1 || dst[1] != 2 || dst[2] != 3 {
		t.Errorf("dst = %v, want [1 2 3]", dst)
	}
}

func TestCopyWordsShortDst(t *testing.T) {
	src := wordBytes([]uint32{1, 2, 3})
	dst := make([]uint32, 2)
	n, ok := CopyWords(dst, src)
	if !ok || n != 2 {
		t.Errorf("CopyWords = %d, %v, want 2, true", n, ok)
	}
}

func TestCopyWordsIsACopy(t *testing.T) {
	vals := []uint32{1, 2}
	src := wordBytes(vals)
	dst := make([]uint32, 2)
	CopyWords(dst, src)
	vals[0] = 99
	if dst[0] != 1 {
		t.Error("dst aliases src; it must be a copy")
	}
}

func TestCopyWordsBadShapes(t *testing.T) {
	src := wordBytes([]uint32{1, 2, 3, 4})
	dst := make([]uint32, 4)
	for _, c := range []struct {
		name string
		in   []byte
	}{
		{"nil", nil},
		{"empty", src[:0]},
		{"length not a multiple of 4", src[:6]},
		{"misaligned", src[1:13]},
	} {
		n, ok := CopyWords(dst, c.in)
		if ok || n != 0 {
			t.Errorf("%s: CopyWords = %d, %v, want 0, false", c.name, n, ok)
		}
	}
}

func TestCopyWordsAllocatesNothing(t *testing.T) {
	src := wordBytes(make([]uint32, 1024))
	dst := make([]uint32, 1024)
	var n int
	if a := testing.AllocsPerRun(100, func() { n, _ = CopyWords(dst, src) }); a != 0 {
		t.Errorf("CopyWords made %v allocations, want 0", a)
	}
	_ = n
}
