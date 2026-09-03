package benchsink

import "testing"

func TestSumTo(t *testing.T) {
	cases := map[int]int{0: 0, -3: 0, 1: 0, 4: 6, 100: 4950}
	for in, want := range cases {
		if got := SumTo(in); got != want {
			t.Errorf("SumTo(%d) = %d, want %d", in, got, want)
		}
	}
}

func TestConsumeStoresAndReturnsPrevious(t *testing.T) {
	Sink = 0
	if got := Consume(7); got != 0 {
		t.Errorf("Consume(7) = %d, want the previous Sink 0", got)
	}
	if Sink != 7 {
		t.Errorf("Sink = %d, want 7", Sink)
	}
	if got := Consume(9); got != 7 {
		t.Errorf("Consume(9) = %d, want the previous Sink 7", got)
	}
	if Sink != 9 {
		t.Errorf("Sink = %d, want 9", Sink)
	}
}

func TestSinkSeesBenchmarkResult(t *testing.T) {
	Sink = 0
	for i := 0; i < 5; i++ {
		Consume(SumTo(i))
	}
	if Sink != SumTo(4) {
		t.Errorf("Sink = %d, want %d", Sink, SumTo(4))
	}
}
