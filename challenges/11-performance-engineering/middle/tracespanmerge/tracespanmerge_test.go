package tracespanmerge

import (
	"reflect"
	"testing"
)

func TestMergeOverlapping(t *testing.T) {
	got := Merge([]Span{{0, 10}, {5, 20}})
	if !reflect.DeepEqual(got, []Span{{0, 20}}) {
		t.Errorf("Merge = %v, want [{0 20}]", got)
	}
}

func TestMergeTouchingSpans(t *testing.T) {
	got := Merge([]Span{{0, 10}, {10, 20}})
	if !reflect.DeepEqual(got, []Span{{0, 20}}) {
		t.Errorf("Merge = %v, want [{0 20}] — touching spans join", got)
	}
}

func TestMergeDisjointSpans(t *testing.T) {
	got := Merge([]Span{{20, 30}, {0, 10}})
	if !reflect.DeepEqual(got, []Span{{0, 10}, {20, 30}}) {
		t.Errorf("Merge = %v, want [{0 10} {20 30}] sorted by start", got)
	}
}

func TestMergeNestedSpan(t *testing.T) {
	got := Merge([]Span{{0, 100}, {10, 20}})
	if !reflect.DeepEqual(got, []Span{{0, 100}}) {
		t.Errorf("Merge = %v, want [{0 100}] — a nested span must not shorten the outer one", got)
	}
}

func TestMergeDropsEmptySpans(t *testing.T) {
	got := Merge([]Span{{5, 5}, {10, 3}, {0, 4}})
	if !reflect.DeepEqual(got, []Span{{0, 4}}) {
		t.Errorf("Merge = %v, want [{0 4}]", got)
	}
	empty := Merge([]Span{{5, 5}})
	if empty == nil || len(empty) != 0 {
		t.Errorf("Merge = %v, want empty non-nil slice", empty)
	}
}

func TestMergeDoesNotModifyInput(t *testing.T) {
	in := []Span{{20, 30}, {0, 10}}
	before := append([]Span(nil), in...)
	Merge(in)
	if !reflect.DeepEqual(in, before) {
		t.Errorf("input was sorted in place: %v, want %v", in, before)
	}
}

func TestCovered(t *testing.T) {
	cases := []struct {
		spans []Span
		want  int64
	}{
		{[]Span{{0, 10}, {5, 20}}, 20},
		{[]Span{{0, 10}, {20, 30}}, 20},
		{[]Span{{0, 100}, {10, 20}}, 100},
		{nil, 0},
		{[]Span{{5, 5}}, 0},
	}
	for _, c := range cases {
		if got := Covered(c.spans); got != c.want {
			t.Errorf("Covered(%v) = %d, want %d", c.spans, got, c.want)
		}
	}
}

func TestCoveredIsNotTheSumOfDurations(t *testing.T) {
	spans := []Span{{0, 100}, {0, 100}, {0, 100}}
	if got := Covered(spans); got != 100 {
		t.Errorf("Covered = %d, want 100 — three parallel spans still cover one interval", got)
	}
}
