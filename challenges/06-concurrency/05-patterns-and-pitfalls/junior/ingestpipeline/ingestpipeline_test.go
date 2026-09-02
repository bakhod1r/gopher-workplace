package ingestpipeline

import (
	"strings"
	"testing"
)

func TestIngestPipeline(t *testing.T) {
	upper := strings.ToUpper
	isError := func(rec string) bool { return strings.HasPrefix(rec, "ERR ") }

	cases := []struct {
		name  string
		lines []string
		want  []string
	}{
		{"mixed", []string{"err disk", "info ok", "err io"}, []string{"ERR DISK", "ERR IO"}},
		{"no_errors", []string{"info ok"}, nil},
		{"all_errors", []string{"err a", "err b"}, []string{"ERR A", "ERR B"}},
		{"order_preserved", []string{"err z", "info m", "err a"}, []string{"ERR Z", "ERR A"}},
		{"empty", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IngestPipeline(tc.lines, upper, isError)
			if len(got) != len(tc.want) {
				t.Fatalf("IngestPipeline(%v) = %v, want %v", tc.lines, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("IngestPipeline(%v) = %v, want %v", tc.lines, got, tc.want)
				}
			}
		})
	}
}
