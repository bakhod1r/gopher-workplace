package ordership

import "testing"

func TestShipOrderIDs(t *testing.T) {
	cases := []struct {
		name string
		ids  []int
	}{
		{"pair", []int{101, 102}},
		{"empty_batch", nil},
		{"single", []int{7}},
		{"five", []int{5, 4, 3, 2, 1}},
		{"zero_ids", []int{0, 0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := make(chan int, len(tc.ids)+1)
			ShipOrderIDs(out, tc.ids)

			var got []int
			for id := range out {
				got = append(got, id)
			}
			if len(got) != len(tc.ids) {
				t.Fatalf("received %v, want %v", got, tc.ids)
			}
			for i := range tc.ids {
				if got[i] != tc.ids[i] {
					t.Fatalf("received %v, want %v", got, tc.ids)
				}
			}
		})
	}
}
