package orderleakbug

import (
	"reflect"
	"testing"
)

func TestRankByCountOrder(t *testing.T) {
	got := RankByCount([]string{"b", "a", "b", "c"})
	if !reflect.DeepEqual(got, []string{"b", "a", "c"}) {
		t.Errorf("RankByCount = %v, want [b a c]", got)
	}
}

func TestRankByCountEmpty(t *testing.T) {
	if got := RankByCount([]int{}); len(got) != 0 {
		t.Errorf("RankByCount = %v, want []", got)
	}
}

func TestRankByCountTiesAreDeterministic(t *testing.T) {
	in := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}
	want := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}
	for i := 0; i < 200; i++ {
		got := RankByCount(in)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("run %d: RankByCount = %v, want %v (ties must break on the value)", i, got, want)
		}
	}
}

func TestRankByCountMixedTies(t *testing.T) {
	in := []int{3, 3, 1, 2, 4, 5}
	want := []int{3, 1, 2, 4, 5}
	for i := 0; i < 200; i++ {
		if got := RankByCount(in); !reflect.DeepEqual(got, want) {
			t.Fatalf("run %d: RankByCount = %v, want %v", i, got, want)
		}
	}
}
