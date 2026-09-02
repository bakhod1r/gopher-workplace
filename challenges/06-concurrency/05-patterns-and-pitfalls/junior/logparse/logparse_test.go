package logparse

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

func TestParseStage(t *testing.T) {
	upper := strings.ToUpper

	cases := []struct {
		name  string
		lines []string
		want  []string
	}{
		{"three_lines", []string{"warn disk", "info ok", "err io"}, []string{"WARN DISK", "INFO OK", "ERR IO"}},
		{"single_line", []string{"boot"}, []string{"BOOT"}},
		{"already_upper", []string{"FATAL"}, []string{"FATAL"}},
		{"blank_line", []string{""}, []string{""}},
		{"empty_stream", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			for rec := range ParseStage(send(tc.lines...), upper) {
				got = append(got, rec)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("ParseStage(%v) = %v, want %v", tc.lines, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("ParseStage(%v) = %v, want %v", tc.lines, got, tc.want)
				}
			}
		})
	}
}
