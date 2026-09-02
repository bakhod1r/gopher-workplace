package stddeletediscardbug

import (
	"reflect"
	"testing"
	"time"
)

func even(n int) bool { return n%2 == 0 }

func TestPurgeKeepsOrder(t *testing.T) {
	got := Purge([]int{1, 2, 3, 4}, even)
	if !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("Purge = %v, want [1 3]", got)
	}
}

func TestPurgeAllAndNone(t *testing.T) {
	if got := Purge([]int{2, 4}, even); len(got) != 0 {
		t.Errorf("Purge = %v, want []", got)
	}
	if got := Purge([]int{1, 3}, even); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("Purge = %v, want [1 3]", got)
	}
}

func TestPurgeAtScale(t *testing.T) {
	const n = 300000
	in := make([]int, n)
	for i := range in {
		in[i] = i
	}
	start := time.Now()
	got := Purge(in, even)
	if len(got) != n/2 {
		t.Fatalf("Purge returned %d elements, want %d", len(got), n/2)
	}
	for i, v := range got {
		if v != 2*i+1 {
			t.Fatalf("got[%d] = %d, want %d", i, v, 2*i+1)
		}
	}
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("scale purge took %v, want under 2s", d)
	}
}
