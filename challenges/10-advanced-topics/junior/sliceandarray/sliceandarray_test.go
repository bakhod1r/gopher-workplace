package sliceandarray

import (
	"testing"
	"unsafe"
)

func TestSizes(t *testing.T) {
	var (
		a [8]int
		s []int
	)
	arr, sl := Sizes()
	if arr != unsafe.Sizeof(a) {
		t.Errorf("array = %d, want %d", arr, unsafe.Sizeof(a))
	}
	if sl != unsafe.Sizeof(s) {
		t.Errorf("slice = %d, want %d", sl, unsafe.Sizeof(s))
	}
}

func TestArrayIsBigger(t *testing.T) {
	arr, sl := Sizes()
	if arr <= sl {
		t.Errorf("array = %d, slice = %d, want the array to be larger", arr, sl)
	}
}

func TestSliceSizeIsIndependentOfLength(t *testing.T) {
	_, sl := Sizes()
	long := make([]int, 100000)
	if got := unsafe.Sizeof(long); got != sl {
		t.Errorf("a 100000-element slice header is %d, want %d: the header does not grow", got, sl)
	}
}

func TestArrayIsEightWords(t *testing.T) {
	arr, _ := Sizes()
	var one int
	if arr != 8*unsafe.Sizeof(one) {
		t.Errorf("array = %d, want %d", arr, 8*unsafe.Sizeof(one))
	}
}
