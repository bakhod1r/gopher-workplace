package countdown

import (
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	if got := Countdown(0); len(got) != 0 {
		t.Errorf("Countdown(0)=%v want empty", got)
	}
	if got := Countdown(3); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("=%v want [3 2 1]", got)
	}
}
