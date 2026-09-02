package batchwriter

import "testing"

func TestBatchInserts(t *testing.T) {
	cases := []struct {
		name string
		rows []string
		size int
		want [][]string
	}{
		{"short_final_batch", []string{"a", "b", "c"}, 2, [][]string{{"a", "b"}, {"c"}}},
		{"exact_multiple", []string{"a", "b"}, 2, [][]string{{"a", "b"}}},
		{"size_one", []string{"a", "b"}, 1, [][]string{{"a"}, {"b"}}},
		{"size_larger_than_stream", []string{"a", "b"}, 9, [][]string{{"a", "b"}}},
		{"zero_size", []string{"a"}, 0, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rows := make(chan string)
			go func() {
				defer close(rows)
				for _, r := range tc.rows {
					rows <- r
				}
			}()

			got := BatchInserts(rows, tc.size)
			if len(got) != len(tc.want) {
				t.Fatalf("BatchInserts(%v, %d) = %v, want %v", tc.rows, tc.size, got, tc.want)
			}
			for i := range got {
				if len(got[i]) != len(tc.want[i]) {
					t.Fatalf("BatchInserts(%v, %d) = %v, want %v", tc.rows, tc.size, got, tc.want)
				}
				for j := range got[i] {
					if got[i][j] != tc.want[i][j] {
						t.Fatalf("BatchInserts(%v, %d) = %v, want %v", tc.rows, tc.size, got, tc.want)
					}
				}
			}
		})
	}
}
