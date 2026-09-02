package escapeaudit

import "testing"

func TestSumValues(t *testing.T) {
	cases := []struct {
		name string
		vs   []int
		want int
	}{
		{"basic", []int{1, 2, 3}, 6},
		{"negatives", []int{-1, 1}, 0},
		{"single", []int{42}, 42},
		{"empty", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumValues(tc.vs); got != tc.want {
				t.Errorf("SumValues = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestStackAgg(t *testing.T) {
	var s Summer = StackAgg{}
	if got := s.Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum = %d, want 6", got)
	}
}

func TestSumBoxed(t *testing.T) {
	if got := SumBoxed([]any{1, 2, 3}); got != 6 {
		t.Errorf("SumBoxed = %d, want 6", got)
	}
	if got := SumBoxed([]any{1, "x", 2.5, 2}); got != 3 {
		t.Errorf("SumBoxed = %d, want 3 (non-ints skipped)", got)
	}
	if got := SumBoxed(nil); got != 0 {
		t.Errorf("SumBoxed(nil) = %d, want 0", got)
	}
}

func TestBoxAll(t *testing.T) {
	got := BoxAll([]int{1, 2})
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if n, ok := got[0].(int); !ok || n != 1 {
		t.Errorf("got[0] = %v", got[0])
	}
	if SumBoxed(got) != 3 {
		t.Errorf("round trip failed: %v", got)
	}
}

func TestSumValuesDoesNotAllocate(t *testing.T) {
	vs := []int{1, 2, 3, 4, 5, 6, 7, 8}

	avg := testing.AllocsPerRun(1000, func() {
		_ = SumValues(vs)
	})
	if avg > 0 {
		t.Errorf("SumValues allocated %.2f times per call, want 0", avg)
	}
}

func TestStackAggThroughInterfaceDoesNotAllocate(t *testing.T) {
	vs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var s Summer = StackAgg{}

	avg := testing.AllocsPerRun(1000, func() {
		_ = s.Sum(vs)
	})
	if avg > 0 {
		t.Errorf("Sum allocated %.2f times per call, want 0", avg)
	}
}

func BenchmarkSumValues(b *testing.B) {
	vs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SumValues(vs)
	}
}

func BenchmarkSumBoxed(b *testing.B) {
	boxed := BoxAll([]int{1, 2, 3, 4, 5, 6, 7, 8})
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SumBoxed(boxed)
	}
}
