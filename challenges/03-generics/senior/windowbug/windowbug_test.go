package windowbug

import (
	"reflect"
	"testing"
)

func TestWindowsFullOnly(t *testing.T) {
	if got := Windows([]int{1, 2, 3}, 2); !reflect.DeepEqual(got, [][]int{{1, 2}, {2, 3}}) {
		t.Errorf("Windows = %v, want [[1 2] [2 3]] (no partial tail)", got)
	}
	if got := Windows([]int{1, 2, 3, 4}, 3); !reflect.DeepEqual(got, [][]int{{1, 2, 3}, {2, 3, 4}}) {
		t.Errorf("Windows = %v, want [[1 2 3] [2 3 4]]", got)
	}
}

func TestWindowsCount(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	for _, n := range []int{1, 2, 3, 5} {
		got := Windows(s, n)
		if want := len(s) - n + 1; len(got) != want {
			t.Errorf("Windows(n=%d) produced %d windows, want %d", n, len(got), want)
		}
		for i, w := range got {
			if len(w) != n {
				t.Errorf("window %d has length %d, want %d", i, len(w), n)
			}
		}
	}
}

func TestWindowsEdges(t *testing.T) {
	if got := Windows([]int{1, 2}, 3); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Windows(n > len) = %v, want []", got)
	}
	if got := Windows([]int{1}, 0); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Windows(0) = %v, want []", got)
	}
}
