package combineconstraint

import (
	"reflect"
	"testing"
)

func TestRenderNumbers(t *testing.T) {
	if got := Render([]int{1, 2}); !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Errorf("Render = %v, want [1 2]", got)
	}
	if got := Render([]float64{1.5}); !reflect.DeepEqual(got, []string{"1.5"}) {
		t.Errorf("Render = %v, want [1.5]", got)
	}
}

func TestRenderText(t *testing.T) {
	if got := Render([]string{"a", "b"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Render = %v, want [a b]", got)
	}
}

func TestRenderEmpty(t *testing.T) {
	got := Render([]int{})
	if got == nil || len(got) != 0 {
		t.Errorf("Render(empty) = %v, want an empty non-nil slice", got)
	}
}
