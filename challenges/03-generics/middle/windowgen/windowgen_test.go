package windowgen

import (
	"reflect"
	"testing"
)

func TestWindows(t *testing.T) {
	if got := Windows([]int{1, 2, 3}, 2); !reflect.DeepEqual(got, [][]int{{1, 2}, {2, 3}}) {
		t.Errorf("Windows = %v, want [[1 2] [2 3]]", got)
	}
	if got := Windows([]int{1, 2}, 2); !reflect.DeepEqual(got, [][]int{{1, 2}}) {
		t.Errorf("Windows = %v, want [[1 2]]", got)
	}
	if got := Windows([]int{1, 2, 3, 4}, 3); !reflect.DeepEqual(got, [][]int{{1, 2, 3}, {2, 3, 4}}) {
		t.Errorf("Windows = %v, want [[1 2 3] [2 3 4]]", got)
	}
}

func TestWindowsEdges(t *testing.T) {
	if got := Windows([]int{1, 2}, 3); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Windows(n > len) = %v, want []", got)
	}
	if got := Windows([]int{1, 2}, 0); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Windows(0) = %v, want []", got)
	}
	if got := Windows([]int{}, 1); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Windows(empty) = %v, want []", got)
	}
}

func TestWindowsAreIndependent(t *testing.T) {
	in := []int{1, 2, 3}
	got := Windows(in, 2)
	got[0][1] = 99
	if in[1] != 2 || got[1][0] != 2 {
		t.Errorf("windows share storage: in=%v got=%v", in, got)
	}
}
