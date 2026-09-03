package localarray

import "testing"

var sink int

func TestDigits(t *testing.T) {
	cases := map[int]int{0: 1, 7: 1, 10: 2, 1234: 4, -99: 2, 1000000: 7}
	for in, want := range cases {
		if got := Digits(in); got != want {
			t.Errorf("Digits(%d) = %d, want %d", in, got, want)
		}
	}
}

func TestDigitsAllocatesNothing(t *testing.T) {
	if n := testing.AllocsPerRun(200, func() { sink = Digits(987654321) }); n != 0 {
		t.Errorf("Digits made %v allocations, want 0: the scratch array must not escape", n)
	}
}
