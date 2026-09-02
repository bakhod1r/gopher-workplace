package tempavg

import "testing"

func chanOf(vals ...float64) <-chan float64 {
	ch := make(chan float64, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestAverageReading(t *testing.T) {
	cases := []struct {
		name     string
		readings []float64
		wantV    float64
		wantOK   bool
	}{
		{"three_readings", []float64{1, 2, 3}, 2, true},
		{"empty_window", nil, 0, false},
		{"single", []float64{5}, 5, true},
		{"halves", []float64{1, 2}, 1.5, true},
		{"below_and_above", []float64{-2, 2}, 0, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotV, gotOK := AverageReading(chanOf(tc.readings...))
			if gotV != tc.wantV || gotOK != tc.wantOK {
				t.Errorf("AverageReading(%v) = %g, %t, want %g, %t",
					tc.readings, gotV, gotOK, tc.wantV, tc.wantOK)
			}
		})
	}
}
