package sizeof

import (
	"testing"
	"unsafe"
)

func TestSizes(t *testing.T) {
	h, id, name := Sizes()
	var want Header
	if h != unsafe.Sizeof(want) {
		t.Errorf("header = %d, want %d", h, unsafe.Sizeof(want))
	}
	if id != unsafe.Sizeof(want.Id) {
		t.Errorf("id = %d, want %d", id, unsafe.Sizeof(want.Id))
	}
	if name != unsafe.Sizeof(want.Name) {
		t.Errorf("name = %d, want %d", name, unsafe.Sizeof(want.Name))
	}
}

func TestSizesAreTypeSizes(t *testing.T) {
	_, _, name := Sizes()
	long := Header{Name: "a very long name that changes nothing about the header"}
	if got := unsafe.Sizeof(long.Name); got != name {
		t.Errorf("string header = %d, want %d: Sizeof measures the type, not the bytes", got, name)
	}
}

func TestHeaderIsBiggerThanItsScalarField(t *testing.T) {
	h, id, _ := Sizes()
	if h <= id {
		t.Errorf("header = %d, id = %d: the struct must be larger than one field", h, id)
	}
}
