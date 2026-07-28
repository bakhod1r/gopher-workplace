package escapeslice

import "testing"

func TestItems(t *testing.T) {
	items := Items([]int{1, 2, 3})
	got := []int{items[0].V, items[1].V, items[2].V}
	want := []int{1, 2, 3}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v (all alias one struct?)", got, want)
		}
	}
}
