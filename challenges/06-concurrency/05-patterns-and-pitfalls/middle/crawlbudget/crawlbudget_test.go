package crawlbudget

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"
)

var errFetch = errors.New("fetch failed")

func cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func fetch(_ context.Context, url string) (Page, error) {
	if strings.HasPrefix(url, "bad") {
		return Page{}, errFetch
	}
	return Page{URL: url, Bytes: len(url)}, nil
}

// tracker wraps fetch and records the highest number of concurrent fetches.
type tracker struct {
	mu      sync.Mutex
	inFlite int
	max     int
}

func (tr *tracker) fetch(ctx context.Context, url string) (Page, error) {
	tr.mu.Lock()
	tr.inFlite++
	if tr.inFlite > tr.max {
		tr.max = tr.inFlite
	}
	tr.mu.Unlock()

	page, err := fetch(ctx, url)

	tr.mu.Lock()
	tr.inFlite--
	tr.mu.Unlock()
	return page, err
}

func TestCrawlPages(t *testing.T) {
	live := context.Background()

	cases := []struct {
		name    string
		ctx     context.Context
		urls    []string
		limit   int
		want    []string
		wantErr error
	}{
		{"budget_two", live, []string{"a", "bb", "ccc", "dddd"}, 2, []string{"a", "bb", "ccc", "dddd"}, nil},
		{"budget_larger_than_batch", live, []string{"x", "yy"}, 10, []string{"x", "yy"}, nil},
		{"serial_budget", live, []string{"p", "q", "r"}, 1, []string{"p", "q", "r"}, nil},
		{"zero_budget_means_one", live, []string{"m", "n"}, 0, []string{"m", "n"}, nil},
		{"one_bad_url", live, []string{"a", "bad-1", "c"}, 2, nil, errFetch},
		{"empty_frontier", live, nil, 3, nil, nil},
		{"crawl_abandoned", cancelled(), []string{"a", "b"}, 2, nil, context.Canceled},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CrawlPages(tc.ctx, tc.urls, tc.limit, fetch)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("CrawlPages() error = %v, want %v", err, tc.wantErr)
			}
			if err != nil {
				if got != nil {
					t.Fatalf("CrawlPages() = %v, want nil pages on error", got)
				}
				return
			}
			if len(got) != len(tc.want) {
				t.Fatalf("CrawlPages() = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i].URL != tc.want[i] {
					t.Fatalf("CrawlPages() = %v, want %v", got, tc.want)
				}
			}
		})
	}
}

func TestCrawlPagesRespectsBudget(t *testing.T) {
	urls := make([]string, 40)
	for i := range urls {
		urls[i] = strings.Repeat("u", i+1)
	}
	for _, limit := range []int{1, 2, 5} {
		tr := &tracker{}
		if _, err := CrawlPages(context.Background(), urls, limit, tr.fetch); err != nil {
			t.Fatalf("CrawlPages() error = %v", err)
		}
		if tr.max > limit {
			t.Errorf("limit %d: saw %d concurrent fetches, want at most %d", limit, tr.max, limit)
		}
	}
}
