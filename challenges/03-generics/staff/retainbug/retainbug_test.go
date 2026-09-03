package retainbug

import (
	"reflect"
	"testing"
)

func TestHeadContents(t *testing.T) {
	if got := Head([]int{1, 2, 3}, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Head = %v, want [1 2]", got)
	}
	if got := Head([]int{1, 2}, 9); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Head = %v, want [1 2]", got)
	}
	if got := Head([]int{1, 2}, -1); len(got) != 0 {
		t.Errorf("Head = %v, want []", got)
	}
}

func TestHeadIsIndependentOfInput(t *testing.T) {
	s := []int{1, 2, 3, 4}
	h := Head(s, 2)
	s[0] = 99
	if h[0] != 1 {
		t.Errorf("h[0] = %d, want 1: the head still views the input's storage", h[0])
	}
}

func TestHeadSurvivesAFullOverwrite(t *testing.T) {
	const n = 1 << 16
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	h := Head(s, 8)
	for i := range s {
		s[i] = -1
	}
	for i := 0; i < 8; i++ {
		if h[i] != i {
			t.Fatalf("h[%d] = %d, want %d: the head is a window onto the input, not a copy", i, h[i], i)
		}
	}
}
