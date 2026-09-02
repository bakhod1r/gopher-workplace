package sortslice

import (
	"sort"
	"testing"
)

func TestSortOrder(t *testing.T) {
	ps := []Player{{"a", 1}, {"b", 3}, {"c", 2}}
	sort.Sort(ByScore(ps))
	want := []string{"b", "c", "a"}
	for i := range want {
		if ps[i].Name != want[i] {
			t.Fatalf("order = %v, want %v", ps, want)
		}
	}
}

func TestTieBrokenByName(t *testing.T) {
	ps := []Player{{"bob", 5}, {"ann", 5}}
	sort.Sort(ByScore(ps))
	if ps[0].Name != "ann" || ps[1].Name != "bob" {
		t.Errorf("tie order = %v, want ann then bob", ps)
	}
}

func TestTopN(t *testing.T) {
	cases := []struct {
		name string
		ps   []Player
		n    int
		want []string
	}{
		{"top_two", []Player{{"a", 1}, {"b", 3}, {"c", 2}}, 2, []string{"b", "c"}},
		{"n_too_big", []Player{{"a", 1}, {"b", 2}}, 10, []string{"b", "a"}},
		{"zero", []Player{{"a", 1}}, 0, nil},
		{"empty", nil, 3, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := TopN(tc.ps, tc.n)
			if len(got) != len(tc.want) {
				t.Fatalf("TopN = %v, want %v", got, tc.want)
			}
			for i := range tc.want {
				if got[i] != tc.want[i] {
					t.Errorf("TopN = %v, want %v", got, tc.want)
				}
			}
		})
	}
}
