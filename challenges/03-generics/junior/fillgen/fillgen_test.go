package fillgen

import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	if got := Fill(3, 7); !reflect.DeepEqual(got, []int{7, 7, 7}) {
		t.Errorf("Fill(3, 7) = %v, want [7 7 7]", got)
	}
	if got := Fill(2, "x"); !reflect.DeepEqual(got, []string{"x", "x"}) {
		t.Errorf("Fill(2, %q) = %v, want [x x]", "x", got)
	}
	if got := Fill(0, 7); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Fill(0, 7) = %v, want []", got)
	}
	if got := Fill(-3, 7); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Fill(-3, 7) = %v, want []", got)
	}
}
