package indexbygen

import "testing"

func TestIndexBy(t *testing.T) {
	type row struct {
		id   int
		name string
	}
	rows := []row{{1, "a"}, {2, "b"}}
	got := IndexBy(rows, func(r row) int { return r.id })
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[1].name != "a" || got[2].name != "b" {
		t.Errorf("index is wrong: %v", got)
	}
}

func TestIndexByLastWins(t *testing.T) {
	got := IndexBy([]int{1, 11}, func(n int) int { return n % 10 })
	if len(got) != 1 || got[1] != 11 {
		t.Errorf("IndexBy = %v, want {1: 11} (last duplicate wins)", got)
	}
}

func TestIndexByEmpty(t *testing.T) {
	got := IndexBy([]int{}, func(n int) int { return n })
	if got == nil || len(got) != 0 {
		t.Errorf("IndexBy(empty) = %v, want an empty non-nil map", got)
	}
}
