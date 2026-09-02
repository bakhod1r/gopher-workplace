package dedupstream

import "testing"

func TestExactSet(t *testing.T) {
	e := NewExactSet()
	if e.Seen("a") {
		t.Error("first Seen should be false")
	}
	if !e.Seen("a") {
		t.Error("second Seen should be true")
	}
	if e.Seen("b") {
		t.Error("new id should be false")
	}
}

func TestDedupExact(t *testing.T) {
	cases := []struct {
		name string
		ids  []string
		want int
	}{
		{"dupes", []string{"a", "a", "b"}, 2},
		{"all_same", []string{"a", "a", "a"}, 1},
		{"all_unique", []string{"a", "b", "c"}, 3},
		{"empty", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Dedup(&SliceSource{IDs: tc.ids}, NewExactSet())
			if got != tc.want {
				t.Errorf("Dedup = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestWindowSetForgets(t *testing.T) {
	w := NewWindowSet(1)
	if got := Dedup(&SliceSource{IDs: []string{"a", "b", "a"}}, w); got != 3 {
		t.Errorf("Dedup = %d, want 3 (a fell out of the window)", got)
	}

	w2 := NewWindowSet(2)
	if got := Dedup(&SliceSource{IDs: []string{"a", "b", "a"}}, w2); got != 2 {
		t.Errorf("Dedup = %d, want 2 (a still in window)", got)
	}
}

func TestWindowIsBounded(t *testing.T) {
	w := NewWindowSet(16)
	ids := make([]string, 0, 100000)
	for i := 0; i < 100000; i++ {
		ids = append(ids, string(rune('a'+i%26))+string(rune('a'+(i/26)%26))+string(rune('a'+i/676)))
	}

	Dedup(&SliceSource{IDs: ids}, w)
	if len(w.seen) > 16 {
		t.Errorf("window holds %d ids, want at most 16", len(w.seen))
	}
	if len(w.order) > 16 {
		t.Errorf("order holds %d ids, want at most 16", len(w.order))
	}
}

func TestZeroWindowCountsEverything(t *testing.T) {
	w := NewWindowSet(0)
	if got := Dedup(&SliceSource{IDs: []string{"a", "a"}}, w); got != 2 {
		t.Errorf("Dedup = %d, want 2", got)
	}
	if len(w.seen) != 0 {
		t.Errorf("seen holds %d ids, want 0", len(w.seen))
	}
}

func BenchmarkDedupWindow(b *testing.B) {
	ids := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		ids = append(ids, string(rune('a'+i%26)))
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Dedup(&SliceSource{IDs: ids}, NewWindowSet(64))
	}
}
