package tracespanoverlap

import "testing"

func TestIntersect(t *testing.T) {
	got, ok := Intersect(Span{0, 10}, Span{5, 20})
	if !ok || got != (Span{5, 10}) {
		t.Errorf("Intersect = %v, %v, want {5 10}, true", got, ok)
	}
	got, ok = Intersect(Span{0, 100}, Span{10, 20})
	if !ok || got != (Span{10, 20}) {
		t.Errorf("Intersect = %v, %v, want {10 20}, true", got, ok)
	}
}

func TestIntersectDisjointAndTouching(t *testing.T) {
	if _, ok := Intersect(Span{0, 10}, Span{20, 30}); ok {
		t.Error("disjoint spans reported an overlap")
	}
	if _, ok := Intersect(Span{0, 10}, Span{10, 20}); ok {
		t.Error("touching spans reported an overlap; the intervals are half-open")
	}
	if _, ok := Intersect(Span{5, 5}, Span{0, 10}); ok {
		t.Error("an empty span reported an overlap")
	}
}

func TestConcurrency(t *testing.T) {
	cases := []struct {
		spans []Span
		want  int
	}{
		{[]Span{{0, 10}, {5, 20}, {6, 7}}, 3},
		{[]Span{{0, 10}, {20, 30}}, 1},
		{[]Span{{0, 10}, {10, 20}}, 1},
		{nil, 0},
		{[]Span{{5, 5}}, 0},
		{[]Span{{0, 100}, {0, 100}, {0, 100}}, 3},
	}
	for _, c := range cases {
		if got := Concurrency(c.spans); got != c.want {
			t.Errorf("Concurrency(%v) = %d, want %d", c.spans, got, c.want)
		}
	}
}

func TestBusiestAt(t *testing.T) {
	at, ok := BusiestAt([]Span{{0, 10}, {5, 20}})
	if !ok || at != 5 {
		t.Errorf("BusiestAt = %d, %v, want 5, true", at, ok)
	}
	at, ok = BusiestAt([]Span{{0, 10}})
	if !ok || at != 0 {
		t.Errorf("BusiestAt = %d, %v, want 0, true", at, ok)
	}
	if _, ok := BusiestAt(nil); ok {
		t.Error("BusiestAt(nil) reported a timestamp")
	}
	if _, ok := BusiestAt([]Span{{5, 5}}); ok {
		t.Error("BusiestAt reported a timestamp for an empty span")
	}
}

func TestBusiestAtTakesTheEarliestPeak(t *testing.T) {
	// Peak of 2 is reached at 5 and again at 100.
	at, ok := BusiestAt([]Span{{0, 10}, {5, 20}, {100, 200}, {100, 300}})
	if !ok || at != 5 {
		t.Errorf("BusiestAt = %d, %v, want 5, true", at, ok)
	}
}
