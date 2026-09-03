package slicelen

import (
	"testing"
	"unsafe"
)

func alignedBytes(n int) []byte {
	u := make([]uint32, (n+3)/4)
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), n)
}

func TestWordsLength(t *testing.T) {
	b := alignedBytes(8)
	v, ok := Words(b)
	if !ok {
		t.Fatal("Words reported false for an aligned 8-byte buffer")
	}
	if len(v) != 2 {
		t.Fatalf("len = %d, want 2: the count is in elements, not bytes", len(v))
	}
	if cap(v) != 2 {
		t.Errorf("cap = %d, want 2", cap(v))
	}
}

func TestWordsCoversExactlyTheBuffer(t *testing.T) {
	b := alignedBytes(16)
	v, ok := Words(b)
	if !ok {
		t.Fatal("Words reported false")
	}
	if len(v)*4 != len(b) {
		t.Errorf("the view covers %d bytes, want %d", len(v)*4, len(b))
	}
}

func TestWordsSharesStorage(t *testing.T) {
	b := alignedBytes(8)
	v, ok := Words(b)
	if !ok {
		t.Fatal("Words reported false")
	}
	v[0] = 0xffffffff
	for i := 0; i < 4; i++ {
		if b[i] != 0xff {
			t.Fatalf("b[%d] = %#x, want 0xff: the view does not share the bytes", i, b[i])
		}
	}
}

func TestWordsRejectsBadShapes(t *testing.T) {
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
		if _, ok := Words(c.in); ok {
			t.Errorf("%s: reported ok, want false", c.name)
		}
	}
}
