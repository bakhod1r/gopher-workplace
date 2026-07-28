package appendnilreturn

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add(nil, 5); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("Add(nil,5)=%v; want [5]", got)
	}
	if got := Add([]int{1, 2}, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Add=%v; want [1 2 3]", got)
	}
}
