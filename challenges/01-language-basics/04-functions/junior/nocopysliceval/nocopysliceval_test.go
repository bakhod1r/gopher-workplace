package nocopysliceval

import (
	"reflect"
	"testing"
)

func TestSumKeep(t *testing.T) {
	xs := []int{3, 1, 2}
	if got := SumKeep(xs); got != 6 {
		t.Errorf("=%d want 6", got)
	}
	if !reflect.DeepEqual(xs, []int{3, 1, 2}) {
		t.Errorf("slice mutated: %v", xs)
	}
}
