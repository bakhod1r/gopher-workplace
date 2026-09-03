package ordermerge

import (
	"reflect"
	"sort"
	"testing"
)

func feedOf(ids ...string) <-chan string {
	ch := make(chan string, len(ids))
	for _, id := range ids {
		ch <- id
	}
	close(ch)
	return ch
}

func collect(t *testing.T, merged <-chan string) []string {
	t.Helper()
	got := []string{}
	for id := range merged {
		got = append(got, id)
	}
	sort.Strings(got)
	return got
}

func TestMergeOrderFeeds(t *testing.T) {
	cases := []struct {
		name  string
		feeds [][]string
		want  []string
	}{
		{"single_feed", [][]string{{"eu-1", "eu-2"}}, []string{"eu-1", "eu-2"}},
		{"two_regions", [][]string{{"eu-1"}, {"us-1", "us-2"}}, []string{"eu-1", "us-1", "us-2"}},
		{"empty_feed_included", [][]string{{"eu-1"}, {}, {"ap-1"}}, []string{"ap-1", "eu-1"}},
		{"no_feeds", nil, []string{}},
		{"all_empty", [][]string{{}, {}}, []string{}},
		{"four_regions", [][]string{{"a"}, {"b"}, {"c"}, {"d"}}, []string{"a", "b", "c", "d"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			feeds := make([]<-chan string, 0, len(tc.feeds))
			for _, ids := range tc.feeds {
				feeds = append(feeds, feedOf(ids...))
			}
			got := collect(t, MergeOrderFeeds(feeds...))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MergeOrderFeeds(%v) = %#v, want %#v", tc.feeds, got, tc.want)
			}
		})
	}
}

func TestMergeOrderFeedsClosesOnce(t *testing.T) {
	merged := MergeOrderFeeds(feedOf("eu-1"), feedOf("us-1"))
	n := 0
	for range merged {
		n++
	}
	if n != 2 {
		t.Fatalf("received %d ids, want 2", n)
	}
	if _, ok := <-merged; ok {
		t.Fatal("merged channel is not closed after both feeds drained")
	}
}
