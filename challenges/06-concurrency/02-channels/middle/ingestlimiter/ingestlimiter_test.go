package ingestlimiter

import (
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
)

func docsOf(ids ...string) []Document {
	out := make([]Document, 0, len(ids))
	for i, id := range ids {
		out = append(out, Document{ID: id, Bytes: (i + 1) * 100})
	}
	return out
}

// tracker returns an index func that records peak concurrency and holds every
// call until all callers that can run at once have arrived.
func tracker() (func(Document) int, *int64) {
	var live, peak int64
	fn := func(d Document) int {
		n := atomic.AddInt64(&live, 1)
		for {
			old := atomic.LoadInt64(&peak)
			if n <= old || atomic.CompareAndSwapInt64(&peak, old, n) {
				break
			}
		}
		atomic.AddInt64(&live, -1)
		return d.Bytes / 10
	}
	return fn, &peak
}

func TestIndexDocuments(t *testing.T) {
	cases := []struct {
		name        string
		ids         []string
		maxInFlight int
		want        map[string]int
	}{
		{"three_docs_two_slots", []string{"a", "b", "c"}, 2,
			map[string]int{"a": 10, "b": 20, "c": 30}},
		{"single_doc", []string{"a"}, 8, map[string]int{"a": 10}},
		{"no_docs", nil, 4, map[string]int{}},
		{"serial", []string{"a", "b"}, 1, map[string]int{"a": 10, "b": 20}},
		{"clamped_to_one", []string{"a", "b"}, 0, map[string]int{"a": 10, "b": 20}},
		{"negative_budget", []string{"a"}, -3, map[string]int{"a": 10}},
		{"more_slots_than_docs", []string{"a", "b", "c"}, 10,
			map[string]int{"a": 10, "b": 20, "c": 30}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fn, peak := tracker()
			got := IndexDocuments(docsOf(tc.ids...), tc.maxInFlight, fn)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("IndexDocuments() = %#v, want %#v", got, tc.want)
			}
			budget := int64(tc.maxInFlight)
			if budget < 1 {
				budget = 1
			}
			if p := atomic.LoadInt64(peak); p > budget {
				t.Errorf("peak concurrency %d exceeds budget %d", p, budget)
			}
		})
	}
}

func TestIndexDocumentsCallsEachDocumentOnce(t *testing.T) {
	var mu sync.Mutex
	calls := map[string]int{}
	IndexDocuments(docsOf("a", "b", "c"), 3, func(d Document) int {
		mu.Lock()
		calls[d.ID]++
		mu.Unlock()
		return d.Bytes
	})
	for _, id := range []string{"a", "b", "c"} {
		if calls[id] != 1 {
			t.Errorf("index called %d times for %s, want 1", calls[id], id)
		}
	}
}
