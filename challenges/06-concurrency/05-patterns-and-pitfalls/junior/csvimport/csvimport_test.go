package csvimport

import (
	"strings"
	"testing"
)

func TestImportCSV(t *testing.T) {
	clean := strings.TrimSpace
	valid := func(row string) bool { return row != "" }

	cases := []struct {
		name string
		rows []string
		want []string
	}{
		{"drops_blank_row", []string{" a ", "  "}, []string{"a"}},
		{"keeps_all", []string{" a ", " b "}, []string{"a", "b"}},
		{"order_preserved", []string{" z", "", "a "}, []string{"z", "a"}},
		{"all_blank", []string{"", "   "}, nil},
		{"empty_file", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ImportCSV(tc.rows, clean, valid)
			if len(got) != len(tc.want) {
				t.Fatalf("ImportCSV(%q) = %q, want %q", tc.rows, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("ImportCSV(%q) = %q, want %q", tc.rows, got, tc.want)
				}
			}
		})
	}
}
