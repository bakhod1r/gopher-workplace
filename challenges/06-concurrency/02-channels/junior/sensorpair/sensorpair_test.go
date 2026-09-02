package sensorpair

import "testing"

// one returns a buffered channel holding exactly one reading. It is left open
// on purpose: a closed channel would always be ready in a select.
func one(v int) <-chan int {
	ch := make(chan int, 1)
	ch <- v
	return ch
}

func TestCombinedReading(t *testing.T) {
	cases := []struct {
		name           string
		temp, humidity int
		want           int
	}{
		{"typical", 21, 40, 61},
		{"zero_temp", 0, 55, 55},
		{"cancels_out", -5, 5, 0},
		{"both_zero", 0, 0, 0},
		{"large", 100, 250, 350},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CombinedReading(one(tc.temp), one(tc.humidity))
			if got != tc.want {
				t.Errorf("CombinedReading(%d, %d) = %d, want %d",
					tc.temp, tc.humidity, got, tc.want)
			}
		})
	}
}
