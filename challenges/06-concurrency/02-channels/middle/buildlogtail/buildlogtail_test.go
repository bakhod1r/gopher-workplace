package buildlogtail

import (
	"reflect"
	"testing"
)

func logOf(lines ...string) <-chan string {
	ch := make(chan string, len(lines))
	for _, l := range lines {
		ch <- l
	}
	close(ch)
	return ch
}

func TestTailBuildLog(t *testing.T) {
	cases := []struct {
		name  string
		lines []string
		keep  int
		want  []string
	}{
		{"more_lines_than_keep", []string{"a", "b", "c"}, 2, []string{"b", "c"}},
		{"fewer_lines_than_keep", []string{"a", "b"}, 5, []string{"a", "b"}},
		{"exactly_keep", []string{"a", "b", "c"}, 3, []string{"a", "b", "c"}},
		{"keep_zero", []string{"a", "b"}, 0, []string{}},
		{"keep_one", []string{"a", "b", "c", "d"}, 1, []string{"d"}},
		{"empty_log", nil, 3, []string{}},
		{"negative_keep", []string{"a"}, -2, []string{}},
		{"long_build", []string{"1", "2", "3", "4", "5", "6", "7"}, 3, []string{"5", "6", "7"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := TailBuildLog(logOf(tc.lines...), tc.keep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TailBuildLog(%v, %d) = %#v, want %#v",
					tc.lines, tc.keep, got, tc.want)
			}
		})
	}
}

func TestTailBuildLogDrainsWholeStream(t *testing.T) {
	log := logOf("a", "b", "c", "d")
	TailBuildLog(log, 2)
	if _, ok := <-log; ok {
		t.Fatal("stream was not drained to close")
	}
}
