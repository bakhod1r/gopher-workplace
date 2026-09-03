package docindexer

import (
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
)

var errEmptyBody = errors.New("document has no body")

// indexer rejects documents with an empty body.
func indexer(calls *int64) func(Doc) error {
	return func(d Doc) error {
		atomic.AddInt64(calls, 1)
		if d.Body == "" {
			return errEmptyBody
		}
		return nil
	}
}

func TestIndexDocuments(t *testing.T) {
	cases := []struct {
		name        string
		docs        []Doc
		wantIndexed []string
		wantFailed  []string
	}{
		{
			"all_indexed",
			[]Doc{{"a", "x"}, {"b", "y"}},
			[]string{"a", "b"}, []string{},
		},
		{
			"one_rejected",
			[]Doc{{"a", "x"}, {"b", ""}},
			[]string{"a"}, []string{"b"},
		},
		{
			"both_lists_sorted",
			[]Doc{{"zed", "x"}, {"beta", ""}, {"alpha", "y"}, {"yank", ""}},
			[]string{"alpha", "zed"}, []string{"beta", "yank"},
		},
		{
			"all_rejected",
			[]Doc{{"m", ""}, {"n", ""}},
			[]string{}, []string{"m", "n"},
		},
		{
			"single_doc",
			[]Doc{{"solo", "body"}},
			[]string{"solo"}, []string{},
		},
		{"no_docs", nil, []string{}, []string{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls int64
			indexed, failed := IndexDocuments(tc.docs, indexer(&calls))
			if !reflect.DeepEqual(indexed, tc.wantIndexed) {
				t.Errorf("indexed = %v, want %v", indexed, tc.wantIndexed)
			}
			if !reflect.DeepEqual(failed, tc.wantFailed) {
				t.Errorf("failed = %v, want %v", failed, tc.wantFailed)
			}
			if int(calls) != len(tc.docs) {
				t.Errorf("index called %d times, want %d", calls, len(tc.docs))
			}
		})
	}
}
