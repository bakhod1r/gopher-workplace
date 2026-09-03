package clipheaderbug

import (
	"reflect"
	"testing"
)

func TestShrinkContents(t *testing.T) {
	s := []int{1, 2, 3, 4}
	if got := Shrink(s, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Shrink = %v, want [1 2]", got)
	}
}

func TestShrinkClamps(t *testing.T) {
	s := []int{1, 2}
	if got := Shrink(s, 9); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Shrink = %v, want [1 2]", got)
	}
	if got := Shrink(s, -3); len(got) != 0 {
		t.Errorf("Shrink = %v, want []", got)
	}
}

func TestShrinkCapacityEqualsLength(t *testing.T) {
	s := []int{1, 2, 3, 4}
	got := Shrink(s, 2)
	if cap(got) != len(got) {
		t.Errorf("cap = %d, len = %d: capacity must be clipped", cap(got), len(got))
	}
}

func TestShrinkResultCannotOverwriteInput(t *testing.T) {
	s := []int{1, 2, 3, 4}
	page := Shrink(s, 2)
	_ = append(page, 99)
	if s[2] != 3 {
		t.Errorf("s[2] = %d, want 3: appending to the page overwrote the input", s[2])
	}
}
