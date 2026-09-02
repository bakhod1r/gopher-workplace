package deadletters

import "testing"

func TestDrainDeadLetters(t *testing.T) {
	cases := []struct {
		name string
		msgs []string
		want int
	}{
		{"three_messages", []string{"a", "b", "c"}, 3},
		{"one_message", []string{"a"}, 1},
		{"five_messages", []string{"a", "b", "c", "d", "e"}, 5},
		{"duplicates_counted", []string{"a", "a"}, 2},
		{"empty_queue", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			msgs := make(chan string)
			go func() {
				defer close(msgs)
				for _, m := range tc.msgs {
					msgs <- m
				}
			}()

			if got := DrainDeadLetters(msgs); got != tc.want {
				t.Errorf("DrainDeadLetters(%v) = %d, want %d", tc.msgs, got, tc.want)
			}
		})
	}
}
