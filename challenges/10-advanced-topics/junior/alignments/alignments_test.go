package alignments

import (
	"testing"
	"unsafe"
)

func TestAlignments(t *testing.T) {
	var (
		vb  byte
		v32 int32
		v64 int64
		vs  string
	)
	b, i32, i64, s := Alignments()
	if b != unsafe.Alignof(vb) {
		t.Errorf("byte = %d, want %d", b, unsafe.Alignof(vb))
	}
	if i32 != unsafe.Alignof(v32) {
		t.Errorf("int32 = %d, want %d", i32, unsafe.Alignof(v32))
	}
	if i64 != unsafe.Alignof(v64) {
		t.Errorf("int64 = %d, want %d", i64, unsafe.Alignof(v64))
	}
	if s != unsafe.Alignof(vs) {
		t.Errorf("string = %d, want %d", s, unsafe.Alignof(vs))
	}
}

func TestAlignmentsAreAscending(t *testing.T) {
	b, i32, i64, _ := Alignments()
	if b != 1 {
		t.Errorf("byte alignment = %d, want 1", b)
	}
	if i32 < b || i64 < i32 {
		t.Errorf("alignments %d, %d, %d are not ascending", b, i32, i64)
	}
}

func TestAlignmentsArePowersOfTwo(t *testing.T) {
	for _, a := range []uintptr{mustAlign(0), mustAlign(1), mustAlign(2), mustAlign(3)} {
		if a == 0 || a&(a-1) != 0 {
			t.Errorf("alignment %d is not a power of two", a)
		}
	}
}

func mustAlign(i int) uintptr {
	b, i32, i64, s := Alignments()
	return []uintptr{b, i32, i64, s}[i]
}
