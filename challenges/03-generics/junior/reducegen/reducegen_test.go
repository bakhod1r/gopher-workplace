package reducegen

import "testing"

func TestReduceSum(t *testing.T) {
	add := func(a, n int) int { return a + n }
	if got := Reduce([]int{1, 2, 3}, 0, add); got != 6 {
		t.Errorf("Reduce([]int{1, 2, 3}, 0, add) = %v, want 6", got)
	}
	if got := Reduce([]int{}, 5, add); got != 5 {
		t.Errorf("Reduce([]int{}, 5, add) = %v, want 5", got)
	}
}

func TestReduceChangesType(t *testing.T) {
	concat := func(a string, s string) string { return a + s }
	if got := Reduce([]string{"a", "b"}, "", concat); got != "ab" {
		t.Errorf("Reduce([]string{\"a\", \"b\"}, \"\", concat) = %q, want \"ab\"", got)
	}
	count := func(a int, s string) int { return a + len(s) }
	if got := Reduce([]string{"ab", "c"}, 0, count); got != 3 {
		t.Errorf("Reduce([]string{\"ab\", \"c\"}, 0, count) = %v, want 3", got)
	}
}
