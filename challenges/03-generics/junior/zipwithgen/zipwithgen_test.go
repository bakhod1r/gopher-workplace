package zipwithgen

import (
	"reflect"
	"strconv"
	"testing"
)

func TestZipWith(t *testing.T) {
	add := func(a, b int) int { return a + b }
	if got := ZipWith([]int{1, 2}, []int{10, 20}, add); !reflect.DeepEqual(got, []int{11, 22}) {
		t.Errorf("ZipWith([1 2], [10 20], add) = %v, want [11 22]", got)
	}
	if got := ZipWith([]int{1, 2, 3}, []int{10}, add); !reflect.DeepEqual(got, []int{11}) {
		t.Errorf("ZipWith([1 2 3], [10], add) = %v, want [11]", got)
	}
	if got := ZipWith([]int{1}, []int{}, add); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("ZipWith([1], [], add) = %v, want []", got)
	}
}

func TestZipWithMixedTypes(t *testing.T) {
	label := func(name string, n int) string { return name + strconv.Itoa(n) }
	got := ZipWith([]string{"a", "b"}, []int{1, 2}, label)
	if !reflect.DeepEqual(got, []string{"a1", "b2"}) {
		t.Errorf("ZipWith = %v, want [a1 b2]", got)
	}
}
