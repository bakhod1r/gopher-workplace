package dropwhilebug

import (
	"reflect"
	"testing"
)

func TestDropWhileKeepsLaterMatches(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	if got := DropWhile([]int{2, 4, 5, 6}, isEven); !reflect.DeepEqual(got, []int{5, 6}) {
		t.Errorf("DropWhile = %v, want [5 6] (later matches are kept)", got)
	}
	if got := DropWhile([]int{2, 1, 2, 3}, isEven); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("DropWhile = %v, want [1 2 3]", got)
	}
}

func TestDropWhileEdges(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	if got := DropWhile([]int{1, 2}, isEven); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("DropWhile = %v, want [1 2]", got)
	}
	if got := DropWhile([]int{2, 4}, isEven); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("DropWhile = %v, want []", got)
	}
	if got := DropWhile([]int{}, isEven); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("DropWhile(empty) = %v, want []", got)
	}
}
