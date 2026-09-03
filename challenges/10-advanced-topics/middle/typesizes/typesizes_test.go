package typesizes

import (
	"testing"
	"unsafe"
)

type small struct {
	A byte
	B byte
}

type wide struct {
	A byte
	B int64
}

func TestSizes(t *testing.T) {
	cases := []struct {
		in    any
		size  uintptr
		align uintptr
	}{
		{int64(0), unsafe.Sizeof(int64(0)), unsafe.Alignof(int64(0))},
		{byte(0), 1, 1},
		{"", unsafe.Sizeof(""), unsafe.Alignof("")},
		{small{}, unsafe.Sizeof(small{}), unsafe.Alignof(small{})},
		{wide{}, unsafe.Sizeof(wide{}), unsafe.Alignof(wide{})},
	}
	for _, c := range cases {
		size, align, ok := Sizes(c.in)
		if !ok {
			t.Errorf("Sizes(%T) reported false", c.in)
			continue
		}
		if size != c.size {
			t.Errorf("Sizes(%T) size = %d, want %d", c.in, size, c.size)
		}
		if align != c.align {
			t.Errorf("Sizes(%T) align = %d, want %d", c.in, align, c.align)
		}
	}
}

func TestSizesNil(t *testing.T) {
	if _, _, ok := Sizes(nil); ok {
		t.Error("Sizes(nil) reported ok, want false")
	}
}

func TestSizesSliceHeader(t *testing.T) {
	size, _, ok := Sizes([]int{1, 2, 3})
	if !ok {
		t.Fatal("Sizes reported false")
	}
	if size != unsafe.Sizeof([]int(nil)) {
		t.Errorf("size = %d, want the header size %d: the elements do not count",
			size, unsafe.Sizeof([]int(nil)))
	}
}

func TestSizesAlignmentDividesSize(t *testing.T) {
	for _, in := range []any{int64(0), small{}, wide{}, "", []int(nil)} {
		size, align, _ := Sizes(in)
		if align == 0 || size%align != 0 {
			t.Errorf("%T: size %d is not a multiple of align %d", in, size, align)
		}
	}
}
