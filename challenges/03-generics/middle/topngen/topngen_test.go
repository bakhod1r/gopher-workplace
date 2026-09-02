package topngen

import (
	"reflect"
	"testing"
)

func TestTopN(t *testing.T) {
	got := TopN([]int{1, 2, 2, 3, 3, 3}, 2)
	if !reflect.DeepEqual(got, []int{3, 2}) {
		t.Errorf("TopN = %v, want [3 2]", got)
	}
}

func TestTopNTiesUseFirstAppearance(t *testing.T) {
	for i := 0; i < 30; i++ {
		got := TopN([]string{"b", "a", "a", "b"}, 2)
		if !reflect.DeepEqual(got, []string{"b", "a"}) {
			t.Fatalf("TopN = %v, want [b a] on every run", got)
		}
	}
}

func TestTopNEdges(t *testing.T) {
	if got := TopN([]int{1, 2}, 5); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("TopN(n > distinct) = %v, want [1 2]", got)
	}
	if got := TopN([]int{1}, 0); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("TopN(0) = %v, want []", got)
	}
	if got := TopN([]int{}, 3); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("TopN(empty) = %v, want []", got)
	}
}
