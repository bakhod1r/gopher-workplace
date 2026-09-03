package padwaste

import (
	"testing"
	"unsafe"
)

type gappy struct {
	A byte
	B int64
	C byte
}

type packed struct {
	B int64
	A byte
	C byte
}

type none struct {
	A int64
	B int64
}

func TestWasteGappy(t *testing.T) {
	var g gappy
	want := unsafe.Sizeof(g) - (unsafe.Sizeof(g.A) + unsafe.Sizeof(g.B) + unsafe.Sizeof(g.C))
	if got := Waste(gappy{}); got != want {
		t.Errorf("Waste = %d, want %d", got, want)
	}
	if got := Waste(gappy{}); got == 0 {
		t.Error("Waste = 0, want a positive number for a badly ordered struct")
	}
}

func TestWastePackedIsSmaller(t *testing.T) {
	if Waste(packed{}) >= Waste(gappy{}) {
		t.Errorf("packed wastes %d and gappy wastes %d, want the packed layout to waste less",
			Waste(packed{}), Waste(gappy{}))
	}
}

func TestWasteNone(t *testing.T) {
	if got := Waste(none{}); got != 0 {
		t.Errorf("Waste = %d, want 0: same-width fields need no padding", got)
	}
}

func TestWasteNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, "s", []int{1}} {
		if got := Waste(in); got != 0 {
			t.Errorf("Waste(%#v) = %d, want 0", in, got)
		}
	}
}

func TestWasteEmptyStruct(t *testing.T) {
	if got := Waste(struct{}{}); got != 0 {
		t.Errorf("Waste = %d, want 0", got)
	}
}
