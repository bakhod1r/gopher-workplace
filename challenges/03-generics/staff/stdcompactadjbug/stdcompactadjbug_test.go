package stdcompactadjbug

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDistinctFoldAdjacent(t *testing.T) {
	got := DistinctFold([]string{"b", "B", "a"})
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("DistinctFold = %v, want [a b]", got)
	}
}

func TestDistinctFoldScattered(t *testing.T) {
	got := DistinctFold([]string{"a", "b", "A", "B", "a"})
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("DistinctFold = %v, want [a b]", got)
	}
}

func TestDistinctFoldDoesNotTouchInput(t *testing.T) {
	in := []string{"b", "a", "B"}
	DistinctFold(in)
	if !reflect.DeepEqual(in, []string{"b", "a", "B"}) {
		t.Errorf("input mutated: %v", in)
	}
	if got := DistinctFold(nil); len(got) != 0 {
		t.Errorf("DistinctFold(nil) = %v, want empty", got)
	}
}

func TestDistinctFoldAtScale(t *testing.T) {
	const distinct = 20000
	const copies = 10
	in := make([]string, 0, distinct*copies)
	for c := 0; c < copies; c++ {
		for i := 0; i < distinct; i++ {
			in = append(in, fmt.Sprintf("tag-%05d", i))
		}
	}
	start := time.Now()
	got := DistinctFold(in)
	if len(got) != distinct {
		t.Fatalf("DistinctFold returned %d values, want %d", len(got), distinct)
	}
	if d := time.Since(start); d > 5*time.Second {
		t.Fatalf("scale dedupe took %v, want under 5s", d)
	}
}
