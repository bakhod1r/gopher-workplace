package streamfilt

import "testing"

func src(lines ...string) *SliceSource { return &SliceSource{Lines: lines} }

func eq(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestPredicates(t *testing.T) {
	if !(Contains{Sub: "err"}).Match("error") {
		t.Error("Contains should match")
	}
	if (Contains{Sub: "err"}).Match("ok") {
		t.Error("Contains should not match")
	}
	if !(MinLen{N: 3}).Match("abc") {
		t.Error("MinLen exact should match")
	}
	if (MinLen{N: 3}).Match("ab") {
		t.Error("MinLen short should not match")
	}
	if (Not{Inner: Contains{Sub: "err"}}).Match("error") {
		t.Error("Not should invert")
	}
	if !(Not{Inner: Not{Inner: Contains{Sub: "err"}}}).Match("error") {
		t.Error("double Not should match")
	}
}

func TestFilterStream(t *testing.T) {
	cases := []struct {
		name  string
		lines []string
		p     Predicate
		want  []string
	}{
		{"contains", []string{"error x", "ok", "err"}, Contains{Sub: "err"}, []string{"error x", "err"}},
		{"minlen", []string{"a", "abc", "abcd"}, MinLen{N: 3}, []string{"abc", "abcd"}},
		{"not", []string{"a", "abc"}, Not{Inner: MinLen{N: 3}}, []string{"a"}},
		{"none_match", []string{"a"}, MinLen{N: 9}, nil},
		{"empty_source", nil, MinLen{N: 0}, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FilterStream(src(tc.lines...), tc.p)
			if !eq(got, tc.want) {
				t.Errorf("FilterStream = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSourceDrained(t *testing.T) {
	s := src("a")
	FilterStream(s, MinLen{N: 0})
	if _, ok := s.Next(); ok {
		t.Error("source should be drained")
	}
}
