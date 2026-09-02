package bagzerobug

import (
	"reflect"
	"testing"
)

func TestCountRepeats(t *testing.T) {
	got := Count([]string{"a", "a", "b"})
	if !reflect.DeepEqual(got, map[string]int{"a": 2, "b": 1}) {
		t.Errorf("Count = %v, want map[a:2 b:1]", got)
	}
}

func TestCountAllSame(t *testing.T) {
	got := Count([]int{1, 1, 1})
	if got[1] != 3 {
		t.Errorf("Count[1] = %d, want 3", got[1])
	}
}

func TestCountEmpty(t *testing.T) {
	if got := Count([]int{}); len(got) != 0 {
		t.Errorf("Count = %v, want empty", got)
	}
}
