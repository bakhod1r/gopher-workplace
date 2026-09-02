package metricdrop

import "testing"

func TestTryRecord(t *testing.T) {
	cases := []struct {
		name string
		make func() chan int
		want bool
	}{
		{"buffer_empty", func() chan int { return make(chan int, 1) }, true},
		{"buffer_full", func() chan int {
			ch := make(chan int, 1)
			ch <- 1
			return ch
		}, false},
		{"unbuffered_no_collector", func() chan int { return make(chan int) }, false},
		{"room_left", func() chan int {
			ch := make(chan int, 2)
			ch <- 1
			return ch
		}, true},
		{"zero_cap", func() chan int { return make(chan int, 0) }, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := tc.make()
			before := len(buffer)
			got := TryRecord(buffer, 5)
			if got != tc.want {
				t.Fatalf("TryRecord() = %t, want %t", got, tc.want)
			}
			after := len(buffer)
			if tc.want && after != before+1 {
				t.Errorf("len(buffer) = %d after a recorded sample, want %d", after, before+1)
			}
			if !tc.want && after != before {
				t.Errorf("len(buffer) = %d after a dropped sample, want %d", after, before)
			}
		})
	}
}
