package topnhot

import (
	"reflect"
	"testing"
)

func TestTopN(t *testing.T) {
	flat := map[string]int64{"a": 3, "b": 9, "c": 3, "d": 1}
	got := TopN(flat, 2)
	want := []Entry{{"b", 9}, {"a", 3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopN = %v, want %v", got, want)
	}
}

func TestTopNTieBrokenByName(t *testing.T) {
	flat := map[string]int64{"z": 5, "a": 5, "m": 5}
	got := TopN(flat, 3)
	want := []Entry{{"a", 5}, {"m", 5}, {"z", 5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopN = %v, want %v", got, want)
	}
}

func TestTopNMoreThanAvailable(t *testing.T) {
	got := TopN(map[string]int64{"a": 1}, 10)
	want := []Entry{{"a", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopN = %v, want %v", got, want)
	}
}

func TestTopNNonPositive(t *testing.T) {
	for _, n := range []int{0, -1} {
		got := TopN(map[string]int64{"a": 1}, n)
		if got == nil || len(got) != 0 {
			t.Errorf("TopN(_, %d) = %v, want empty non-nil slice", n, got)
		}
	}
}

func TestTopNIsDeterministic(t *testing.T) {
	flat := map[string]int64{"a": 2, "b": 2, "c": 2, "d": 2, "e": 2}
	first := TopN(flat, 3)
	for i := 0; i < 20; i++ {
		if !reflect.DeepEqual(TopN(flat, 3), first) {
			t.Fatalf("TopN is not deterministic across map iteration orders: %v vs %v", TopN(flat, 3), first)
		}
	}
}
