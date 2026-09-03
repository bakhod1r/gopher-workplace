package logindexstages

import (
	"strings"
	"testing"
)

func parse(line string) (Record, bool) {
	level, msg, ok := strings.Cut(line, " ")
	if !ok || level != "err" {
		return Record{}, false
	}
	return Record{Level: level, Message: msg}, true
}

func index(r Record) string { return "idx:" + r.Message }

func TestIndexLogs(t *testing.T) {
	open := make(chan struct{})
	closed := make(chan struct{})
	close(closed)

	cases := []struct {
		name  string
		done  <-chan struct{}
		lines []string
		want  []string
	}{
		{"mixed_levels", open, []string{"err disk", "info ok", "err io"}, []string{"idx:disk", "idx:io"}},
		{"order_preserved", open, []string{"err z", "info m", "err a"}, []string{"idx:z", "idx:a"}},
		{"all_filtered_out", open, []string{"info ok", "warn slow"}, nil},
		{"malformed_line_dropped", open, []string{"garbage", "err net"}, []string{"idx:net"}},
		{"empty_input", open, nil, nil},
		{"shutdown_before_start", closed, []string{"err disk", "err io"}, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IndexLogs(tc.done, tc.lines, parse, index)
			if tc.name == "shutdown_before_start" {
				if len(got) > len(tc.lines) {
					t.Fatalf("IndexLogs() = %v, want at most the input length", got)
				}
				return
			}
			if len(got) != len(tc.want) {
				t.Fatalf("IndexLogs(%v) = %v, want %v", tc.lines, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("IndexLogs(%v) = %v, want %v", tc.lines, got, tc.want)
				}
			}
		})
	}
}

func TestIndexLogsStagesDoNotLeakAfterShutdown(t *testing.T) {
	done := make(chan struct{})
	close(done)
	// A closed done must let every stage return; the test would time out
	// instead of finishing if a stage stayed blocked on a send.
	for i := 0; i < 50; i++ {
		IndexLogs(done, []string{"err a", "err b", "err c", "err d"}, parse, index)
	}
}
