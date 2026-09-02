package searchresults

import "testing"

func TestTopResults(t *testing.T) {
	stream := func(vals []string) <-chan string {
		ch := make(chan string)
		go func() {
			defer close(ch)
			for _, v := range vals {
				ch <- v
			}
		}()
		return ch
	}

	cases := []struct {
		name string
		hits []string
		n    int
		want []string
	}{
		{"top_three_of_five", []string{"a", "b", "c", "d", "e"}, 3, []string{"a", "b", "c"}},
		{"n_larger_than_stream", []string{"a", "b"}, 9, []string{"a", "b"}},
		{"top_one", []string{"a", "b", "c"}, 1, []string{"a"}},
		{"zero_requested", []string{"a", "b"}, 0, nil},
		{"empty_stream", nil, 3, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := TopResults(stream(tc.hits), tc.n)
			if len(got) != len(tc.want) {
				t.Fatalf("TopResults(%v, %d) = %v, want %v", tc.hits, tc.n, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("TopResults(%v, %d) = %v, want %v", tc.hits, tc.n, got, tc.want)
				}
			}
		})
	}
}
