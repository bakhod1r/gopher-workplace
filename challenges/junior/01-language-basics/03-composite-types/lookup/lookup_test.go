package lookup

import "testing"

func TestLookup(t *testing.T) {
	cases := []struct {
		name   string
		m      map[string]int
		key    string
		wantV  int
		wantOK bool
	}{
		{"present nonzero", map[string]int{"a": 5}, "a", 5, true},
		{"absent", map[string]int{"a": 5}, "z", 0, false},
		{"present zero value", map[string]int{"z": 0}, "z", 0, true},
		{"nil map", nil, "a", 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotV, gotOK := Lookup(tc.m, tc.key)
			if gotV != tc.wantV || gotOK != tc.wantOK {
				t.Errorf("Lookup(%v, %q) = %d, %v; want %d, %v",
					tc.m, tc.key, gotV, gotOK, tc.wantV, tc.wantOK)
			}
		})
	}
}
