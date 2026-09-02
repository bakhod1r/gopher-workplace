package bucketsgen

import (
	"reflect"
	"testing"
)

func TestBuckets(t *testing.T) {
	if got := Buckets([]int{1, 5, 9}, []int{5}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Buckets = %v, want [1 2] (5 lands in the upper bucket)", got)
	}
	if got := Buckets([]int{1, 4, 6, 11}, []int{5, 10}); !reflect.DeepEqual(got, []int{2, 1, 1}) {
		t.Errorf("Buckets = %v, want [2 1 1]", got)
	}
	if got := Buckets([]string{"a", "m"}, []string{"g"}); !reflect.DeepEqual(got, []int{1, 1}) {
		t.Errorf("Buckets = %v, want [1 1]", got)
	}
}

func TestBucketsEdges(t *testing.T) {
	if got := Buckets([]int{1, 2}, []int{}); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("Buckets(no edges) = %v, want [2]", got)
	}
	if got := Buckets([]int{}, []int{5}); !reflect.DeepEqual(got, []int{0, 0}) {
		t.Errorf("Buckets(empty) = %v, want [0 0]", got)
	}
}
