package shardids

import "testing"

func TestStreamShardIDs(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want []int
	}{
		{"three", 3, []int{0, 1, 2}},
		{"one", 1, []int{0}},
		{"zero", 0, nil},
		{"negative", -2, nil},
		{"five", 5, []int{0, 1, 2, 3, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			for id := range StreamShardIDs(tc.n) {
				got = append(got, id)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("StreamShardIDs(%d) = %v, want %v", tc.n, got, tc.want)
			}
			for i := range tc.want {
				if got[i] != tc.want[i] {
					t.Fatalf("StreamShardIDs(%d) = %v, want %v", tc.n, got, tc.want)
				}
			}
		})
	}
}
