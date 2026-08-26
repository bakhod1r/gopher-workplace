package peek

import "testing"

func TestPeek(t *testing.T) {
	cases := []struct {
		name    string
		items   []int
		wantVal int
		wantOK  bool
	}{
		{"three_items", []int{1, 2, 3}, 3, true},
		{"single", []int{42}, 42, true},
		{"empty", nil, 0, false},
		{"empty_slice", []int{}, 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Stack{Items: tc.items}
			v, ok := s.Peek()
			if v != tc.wantVal || ok != tc.wantOK {
				t.Errorf("Peek() = (%d, %v), want (%d, %v)",
					v, ok, tc.wantVal, tc.wantOK)
			}
			// Peek must NOT modify the stack.
			if len(s.Items) != len(tc.items) {
				t.Errorf("Peek mutated stack: len = %d, want %d",
					len(s.Items), len(tc.items))
			}
		})
	}
}
