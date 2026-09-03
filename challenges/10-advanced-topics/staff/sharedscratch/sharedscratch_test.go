package sharedscratch

import (
	"strconv"
	"strings"
	"testing"
)

func TestEncodeAllSmall(t *testing.T) {
	got := EncodeAll([][]int{{1, 2}, {3}})
	if len(got) != 2 || got[0] != "1,2" || got[1] != "3" {
		t.Errorf("EncodeAll = %q, want [1,2 3]", got)
	}
	if got := EncodeAll(nil); len(got) != 0 {
		t.Errorf("EncodeAll = %q, want empty", got)
	}
}

func TestEncodeAllUnderConcurrency(t *testing.T) {
	const n = 64
	batches := make([][]int, n)
	for i := range batches {
		batches[i] = []int{i, i * 2, i * 3}
	}
	for round := 0; round < 20; round++ {
		got := EncodeAll(batches)
		for i := range batches {
			want := strings.Join([]string{
				strconv.Itoa(i), strconv.Itoa(i * 2), strconv.Itoa(i * 3),
			}, ",")
			if got[i] != want {
				t.Fatalf("round %d: result %d = %q, want %q: the goroutines share one buffer", round, i, got[i], want)
			}
		}
	}
}
