package appendifgen

import (
	"reflect"
	"testing"
)

func TestAppendIf(t *testing.T) {
	if got := AppendIf([]int{1}, 2, true); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("AppendIf([]int{1}, 2, true) = %v, want [1 2]", got)
	}
	if got := AppendIf([]int{1}, 2, false); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("AppendIf([]int{1}, 2, false) = %v, want [1]", got)
	}
	if got := AppendIf([]string{}, "a", true); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("AppendIf([]string{}, \"a\", true) = %v, want [a]", got)
	}
	if got := AppendIf([]int(nil), 1, false); len(got) != 0 {
		t.Errorf("AppendIf(nil, 1, false) = %v, want an empty result", got)
	}
}
