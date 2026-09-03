package escapecallback

import "testing"

var sink int64

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum = %d, want 6", got)
	}
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
	if got := Sum([]int{-5, 5}); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
}

func TestSumLarge(t *testing.T) {
	s := make([]int, 1000)
	var want int64
	for i := range s {
		s[i] = i
		want += int64(i)
	}
	if got := Sum(s); got != want {
		t.Errorf("Sum = %d, want %d", got, want)
	}
}

func TestSumAllocatesNothing(t *testing.T) {
	s := make([]int, 512)
	for i := range s {
		s[i] = i
	}
	if n := testing.AllocsPerRun(100, func() { sink = Sum(s) }); n != 0 {
		t.Errorf("Sum made %v allocations, want 0: the accumulator is escaping through the callback", n)
	}
}
