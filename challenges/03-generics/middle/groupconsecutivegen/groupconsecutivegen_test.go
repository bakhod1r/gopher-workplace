package groupconsecutivegen

import (
	"reflect"
	"testing"
)

func TestGroupRuns(t *testing.T) {
	if got := GroupRuns([]int{1, 1, 2, 1}); !reflect.DeepEqual(got, [][]int{{1, 1}, {2}, {1}}) {
		t.Errorf("GroupRuns = %v, want [[1 1] [2] [1]]", got)
	}
	if got := GroupRuns([]int{1, 2}); !reflect.DeepEqual(got, [][]int{{1}, {2}}) {
		t.Errorf("GroupRuns = %v, want [[1] [2]]", got)
	}
	if got := GroupRuns([]int{1, 1, 1}); !reflect.DeepEqual(got, [][]int{{1, 1, 1}}) {
		t.Errorf("GroupRuns = %v, want [[1 1 1]]", got)
	}
	if got := GroupRuns([]int{}); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("GroupRuns(empty) = %v, want []", got)
	}
	if got := GroupRuns([]string{"a", "a", "b"}); !reflect.DeepEqual(got, [][]string{{"a", "a"}, {"b"}}) {
		t.Errorf("GroupRuns = %v, want [[a a] [b]]", got)
	}
}
