package sensorlast

import "testing"

func chanOf(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestLastReading(t *testing.T) {
	cases := []struct {
		name     string
		readings []int
		wantV    int
		wantOK   bool
	}{
		{"three_readings", []int{1, 2, 3}, 3, true},
		{"silent_device", nil, 0, false},
		{"single", []int{7}, 7, true},
		{"ends_with_zero", []int{5, 0}, 0, true},
		{"negatives", []int{-1, -9}, -9, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotV, gotOK := LastReading(chanOf(tc.readings...))
			if gotV != tc.wantV || gotOK != tc.wantOK {
				t.Errorf("LastReading(%v) = %d, %t, want %d, %t",
					tc.readings, gotV, gotOK, tc.wantV, tc.wantOK)
			}
		})
	}
}
