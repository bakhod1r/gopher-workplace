package logbatch

import (
	"reflect"
	"testing"
)

func chanOf(vals ...string) <-chan string {
	ch := make(chan string, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestCollectLines(t *testing.T) {
	cases := []struct {
		name  string
		lines []string
		want  []string
	}{
		{"pair", []string{"a", "b"}, []string{"a", "b"}},
		{"silent_container", nil, []string{}},
		{"single", []string{"x"}, []string{"x"}},
		{"order_kept", []string{"c", "b", "a"}, []string{"c", "b", "a"}},
		{"blank_line", []string{""}, []string{""}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CollectLines(chanOf(tc.lines...))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CollectLines(%v) = %#v, want %#v", tc.lines, got, tc.want)
			}
		})
	}
}
