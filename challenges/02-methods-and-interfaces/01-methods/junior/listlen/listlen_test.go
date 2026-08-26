package listlen

import "testing"

func TestLen(t *testing.T) {
	cases := []struct {
		name string
		head *Node
		want int
	}{
		{"three", &Node{1, &Node{2, &Node{3, nil}}}, 3},
		{"one", &Node{42, nil}, 1},
		{"nil", nil, 0},
		{"two", &Node{10, &Node{20, nil}}, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.head.Len(); got != tc.want {
				t.Errorf("Len() = %d, want %d", got, tc.want)
			}
		})
	}
}
