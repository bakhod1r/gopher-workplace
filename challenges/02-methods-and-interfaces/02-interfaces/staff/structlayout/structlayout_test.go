package structlayout

import (
	"testing"
	"unsafe"
)

func TestPackedIsSmaller(t *testing.T) {
	padded := unsafe.Sizeof(Padded{})
	packed := unsafe.Sizeof(Packed{})

	if packed >= padded {
		t.Errorf("sizeof(Packed) = %d, sizeof(Padded) = %d; the packed layout must be smaller",
			packed, padded)
	}
}

func TestSizeMethodsMatchSizeof(t *testing.T) {
	if got, want := (Padded{}).Size(), unsafe.Sizeof(Padded{}); got != want {
		t.Errorf("Padded.Size = %d, want %d", got, want)
	}
	if got, want := (Packed{}).Size(), unsafe.Sizeof(Packed{}); got != want {
		t.Errorf("Packed.Size = %d, want %d", got, want)
	}
}

func TestPackedHasNoTrailingWaste(t *testing.T) {
	// int64 (8) + bool + bool, rounded up to the 8-byte alignment = 16.
	if got := unsafe.Sizeof(Packed{}); got != 16 {
		t.Errorf("sizeof(Packed) = %d, want 16 on a 64-bit build", got)
	}
}

func TestFieldsSurvive(t *testing.T) {
	p := Packed{B: 42, A: true, C: false}
	if p.B != 42 || !p.A || p.C {
		t.Errorf("got %+v", p)
	}

	d := Padded{A: true, B: 42, C: false}
	if d.B != 42 || !d.A || d.C {
		t.Errorf("got %+v", d)
	}
}

func TestOffsetsAreAscending(t *testing.T) {
	var p Packed
	if unsafe.Offsetof(p.B) != 0 {
		t.Errorf("the largest field should come first, offset = %d", unsafe.Offsetof(p.B))
	}
	if unsafe.Offsetof(p.A) < unsafe.Offsetof(p.B) {
		t.Error("field offsets should increase in declaration order")
	}
}

func TestTotalBytes(t *testing.T) {
	const n = 1_000_000

	padded := TotalBytes(Padded{}, n)
	packed := TotalBytes(Packed{}, n)

	if padded != unsafe.Sizeof(Padded{})*n {
		t.Errorf("TotalBytes(Padded) = %d", padded)
	}
	if packed >= padded {
		t.Errorf("packed total = %d, padded total = %d; packed must be smaller", packed, padded)
	}
	if got := TotalBytes(Packed{}, 0); got != 0 {
		t.Errorf("TotalBytes(_, 0) = %d, want 0", got)
	}
	if got := TotalBytes(Packed{}, -5); got != 0 {
		t.Errorf("TotalBytes(_, -5) = %d, want 0", got)
	}
}

func TestIsSizer(t *testing.T) {
	sizers := []Sizer{Padded{}, Packed{}}
	for _, s := range sizers {
		if s.Size() == 0 {
			t.Errorf("%T: Size = 0", s)
		}
	}
}
