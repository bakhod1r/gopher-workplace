package scrapecollect

import (
	"reflect"
	"testing"
	"time"
)

func scrapesOf(closed bool, samples ...Sample) <-chan Sample {
	ch := make(chan Sample, len(samples)+1)
	for _, s := range samples {
		ch <- s
	}
	if closed {
		close(ch)
	}
	return ch
}

var (
	api = Sample{"api", 12}
	db  = Sample{"db", 7}
	web = Sample{"web", 3}
)

func TestCollectScrapes(t *testing.T) {
	cases := []struct {
		name    string
		ready   []Sample
		closed  bool
		want    int
		budget  time.Duration
		wantOut []Sample
		wantOK  bool
	}{
		{"full_set", []Sample{api, db, web}, false, 3, 5 * time.Second, []Sample{api, db, web}, true},
		{"more_ready_than_wanted", []Sample{api, db, web}, false, 2, 5 * time.Second, []Sample{api, db}, true},
		{"pool_closed_early", []Sample{api}, true, 3, 5 * time.Second, []Sample{api}, false},
		{"budget_expired", nil, false, 2, 20 * time.Millisecond, []Sample{}, false},
		{"partial_then_budget", []Sample{api}, false, 3, 20 * time.Millisecond, []Sample{api}, false},
		{"want_zero", []Sample{api}, false, 0, 5 * time.Second, []Sample{}, true},
		{"closed_and_empty", nil, true, 1, 5 * time.Second, []Sample{}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := CollectScrapes(scrapesOf(tc.closed, tc.ready...), tc.want, tc.budget)
			if !reflect.DeepEqual(got, tc.wantOut) || ok != tc.wantOK {
				t.Errorf("CollectScrapes(want=%d) = %#v, %v; want %#v, %v",
					tc.want, got, ok, tc.wantOut, tc.wantOK)
			}
		})
	}
}
