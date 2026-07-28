package hoistedaddr

import "testing"

func TestPointers(t *testing.T) {
	ps := Pointers([]int{10, 20, 30})
	got := []int{*ps[0], *ps[1], *ps[2]}
	want := []int{10, 20, 30}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v (all alias one var?)", got, want)
		}
	}
}
