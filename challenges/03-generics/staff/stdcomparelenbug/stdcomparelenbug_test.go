package stdcomparelenbug

import "testing"

func TestComparePathsLengthWins(t *testing.T) {
	if got := ComparePaths([]int{9}, []int{1, 2}); got != -1 {
		t.Errorf("ComparePaths([9],[1 2]) = %d, want -1", got)
	}
	if got := ComparePaths([]int{1, 2}, []int{9}); got != 1 {
		t.Errorf("ComparePaths([1 2],[9]) = %d, want 1", got)
	}
}

func TestComparePathsTieBreak(t *testing.T) {
	if got := ComparePaths([]int{1, 2}, []int{1, 3}); got != -1 {
		t.Errorf("ComparePaths = %d, want -1", got)
	}
	if got := ComparePaths([]int{1, 3}, []int{1, 2}); got != 1 {
		t.Errorf("ComparePaths = %d, want 1", got)
	}
}

func TestComparePathsEqual(t *testing.T) {
	if got := ComparePaths([]int{1, 2}, []int{1, 2}); got != 0 {
		t.Errorf("ComparePaths = %d, want 0", got)
	}
	if got := ComparePaths(nil, nil); got != 0 {
		t.Errorf("ComparePaths(nil,nil) = %d, want 0", got)
	}
}
