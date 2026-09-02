package rangegen

import (
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	if got := Range(3); !reflect.DeepEqual(got, []int{0, 1, 2}) {
		t.Errorf("Range(3) = %v, want [0 1 2]", got)
	}
	if got := Range(int64(2)); !reflect.DeepEqual(got, []int64{0, 1}) {
		t.Errorf("Range(int64(2)) = %v, want [0 1]", got)
	}
	if got := Range(0); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Range(0) = %v, want []", got)
	}
	if got := Range(-2); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Range(-2) = %v, want []", got)
	}
}
