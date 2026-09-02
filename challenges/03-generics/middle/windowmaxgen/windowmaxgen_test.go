package windowmaxgen

import (
	"reflect"
	"testing"
)

func TestWindowMax(t *testing.T) {
	if got := WindowMax([]int{1, 3, 2}, 2); !reflect.DeepEqual(got, []int{3, 3}) {
		t.Errorf("WindowMax = %v, want [3 3]", got)
	}
	if got := WindowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3); !reflect.DeepEqual(got, []int{3, 3, 5, 5, 6, 7}) {
		t.Errorf("WindowMax = %v, want [3 3 5 5 6 7]", got)
	}
	if got := WindowMax([]int{1, 2, 3}, 3); !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("WindowMax = %v, want [3]", got)
	}
	if got := WindowMax([]int{4, 3, 2}, 1); !reflect.DeepEqual(got, []int{4, 3, 2}) {
		t.Errorf("WindowMax = %v, want [4 3 2]", got)
	}
}

func TestWindowMaxEdges(t *testing.T) {
	if got := WindowMax([]int{1}, 2); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("WindowMax(n > len) = %v, want []", got)
	}
	if got := WindowMax([]int{1, 2}, 0); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("WindowMax(0) = %v, want []", got)
	}
	if got := WindowMax([]string{"a", "c", "b"}, 2); !reflect.DeepEqual(got, []string{"c", "c"}) {
		t.Errorf("WindowMax = %v, want [c c]", got)
	}
}
