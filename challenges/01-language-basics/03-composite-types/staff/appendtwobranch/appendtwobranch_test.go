package appendtwobranch

import (
	"reflect"
	"testing"
)

func TestBranch(t *testing.T) {
	a := make([]int, 2, 10) // spare capacity
	a[0], a[1] = 1, 2
	b, c := Branch(a, 3, 4)
	if !reflect.DeepEqual(b, []int{1, 2, 3}) {
		t.Errorf("b=%v; want [1 2 3]", b)
	}
	if !reflect.DeepEqual(c, []int{1, 2, 4}) {
		t.Errorf("c=%v; want [1 2 4]", c)
	}
}
