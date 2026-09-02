package stdrepeatsharedbug

import (
	"reflect"
	"testing"
)

func TestBlankRowsAreIndependent(t *testing.T) {
	proto := []int{0, 0}
	b := Blank(proto, 3)
	b[0][0] = 7
	if b[1][0] != 0 || b[2][0] != 0 {
		t.Errorf("rows share storage: %v", b)
	}
	if proto[0] != 0 {
		t.Errorf("proto mutated: %v", proto)
	}
}

func TestBlankContent(t *testing.T) {
	b := Blank([]int{1, 2}, 2)
	if !reflect.DeepEqual(b, [][]int{{1, 2}, {1, 2}}) {
		t.Errorf("Blank = %v, want [[1 2] [1 2]]", b)
	}
}

func TestBlankNonPositive(t *testing.T) {
	if got := Blank([]int{1}, 0); len(got) != 0 {
		t.Errorf("Blank = %v, want []", got)
	}
	if got := Blank([]int{1}, -3); len(got) != 0 {
		t.Errorf("Blank = %v, want []", got)
	}
}
