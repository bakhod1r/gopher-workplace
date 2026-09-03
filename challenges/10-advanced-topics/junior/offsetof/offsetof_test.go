package offsetof

import (
	"testing"
	"unsafe"
)

func TestOffsets(t *testing.T) {
	var r Rec
	a, b, c := Offsets()
	if a != unsafe.Offsetof(r.A) || b != unsafe.Offsetof(r.B) || c != unsafe.Offsetof(r.C) {
		t.Errorf("Offsets = %d, %d, %d, want %d, %d, %d",
			a, b, c, unsafe.Offsetof(r.A), unsafe.Offsetof(r.B), unsafe.Offsetof(r.C))
	}
}

func TestFirstFieldStartsAtZero(t *testing.T) {
	if a, _, _ := Offsets(); a != 0 {
		t.Errorf("first offset = %d, want 0", a)
	}
}

func TestAlignmentCreatesAGap(t *testing.T) {
	var r Rec
	a, b, _ := Offsets()
	if b-a <= unsafe.Sizeof(r.A) {
		t.Errorf("B starts at %d, right after a %d-byte field: alignment should have pushed it further",
			b, unsafe.Sizeof(r.A))
	}
}

func TestOffsetsAreAscending(t *testing.T) {
	a, b, c := Offsets()
	if !(a < b && b < c) {
		t.Errorf("offsets %d, %d, %d are not ascending", a, b, c)
	}
}
