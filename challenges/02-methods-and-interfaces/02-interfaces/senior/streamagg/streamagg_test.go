package streamagg

import "testing"

type sliceSource struct {
	data []int
	pos  int
}

func (s *sliceSource) Next() (int, bool) {
	if s.pos >= len(s.data) {
		return 0, false
	}
	v := s.data[s.pos]
	s.pos++
	return v, true
}

func TestMean(t *testing.T) {
	cases := []struct {
		name string
		data []int
		want int
	}{
		{"simple", []int{1, 2, 3}, 2},
		{"truncates", []int{1, 2}, 1},
		{"negatives", []int{-4, 2}, -1},
		{"empty", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Aggregate(&sliceSource{data: tc.data}, &MeanAgg{}); got != tc.want {
				t.Errorf("mean = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		name string
		data []int
		want int
	}{
		{"simple", []int{1, 5, 2}, 5},
		{"all_negative", []int{-5, -2}, -2},
		{"single", []int{7}, 7},
		{"empty", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Aggregate(&sliceSource{data: tc.data}, &MaxAgg{}); got != tc.want {
				t.Errorf("max = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestLargeStreamIsConstantMemory(t *testing.T) {
	const n = 1_000_000

	agg := &MeanAgg{}
	avg := testing.AllocsPerRun(1, func() {
		Aggregate(&RangeSource{N: n}, agg)
	})
	if avg > 8 {
		t.Errorf("Aggregate allocated %.0f times over %d readings; must be bounded", avg, n)
	}
}

func TestLargeStreamResult(t *testing.T) {
	const n = 1_000_000
	got := Aggregate(&RangeSource{N: n}, &MeanAgg{})
	want := (n + 1) / 2
	if got != want {
		t.Errorf("mean of 1..%d = %d, want %d", n, got, want)
	}
}

func BenchmarkAggregate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Aggregate(&RangeSource{N: 10000}, &MeanAgg{})
	}
}
