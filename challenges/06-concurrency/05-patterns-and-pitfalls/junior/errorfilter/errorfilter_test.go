package errorfilter

import (
	"strings"
	"testing"
)

func send(vals ...string) <-chan string {
	ch := make(chan string, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestErrorFilter(t *testing.T) {
	isError := func(rec string) bool { return strings.HasPrefix(rec, "ERR ") }

	cases := []struct {
		name    string
		records []string
		want    []string
	}{
		{"mixed_levels", []string{"ERR disk", "INFO ok", "ERR io"}, []string{"ERR disk", "ERR io"}},
		{"no_errors", []string{"INFO ok", "WARN slow"}, nil},
		{"all_errors", []string{"ERR a", "ERR b"}, []string{"ERR a", "ERR b"}},
		{"prefix_only_counts", []string{"NOTERR x", "ERR y"}, []string{"ERR y"}},
		{"empty_stream", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			for rec := range ErrorFilter(send(tc.records...), isError) {
				got = append(got, rec)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("ErrorFilter(%v) = %v, want %v", tc.records, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("ErrorFilter(%v) = %v, want %v", tc.records, got, tc.want)
				}
			}
		})
	}
}
