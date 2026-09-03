package warmupgate

import (
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
)

// warmer fails for any key prefixed with "cold-", returning a distinct error
// per key so the test can tell which failure won.
func warmer(calls *int64) func(string) error {
	return func(key string) error {
		atomic.AddInt64(calls, 1)
		if len(key) >= 5 && key[:5] == "cold-" {
			return fmt.Errorf("cold source: %s", key)
		}
		return nil
	}
}

func TestWarmCaches(t *testing.T) {
	cases := []struct {
		name    string
		keys    []string
		wantErr string
	}{
		{"all_warm", []string{"tags", "prices", "skus"}, ""},
		{"one_cold", []string{"tags", "cold-prices"}, "cold source: cold-prices"},
		{"lowest_index_wins", []string{"cold-a", "tags", "cold-z"}, "cold source: cold-a"},
		{"cold_in_middle", []string{"tags", "cold-mid", "cold-late"}, "cold source: cold-mid"},
		{"all_cold", []string{"cold-x", "cold-y"}, "cold source: cold-x"},
		{"no_keys", nil, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls int64
			err := WarmCaches(tc.keys, warmer(&calls))
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("WarmCaches() = %v, want nil", err)
				}
			} else {
				if err == nil {
					t.Fatalf("WarmCaches() = nil, want %q", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Fatalf("WarmCaches() = %q, want %q", err.Error(), tc.wantErr)
				}
			}
			if int(calls) != len(tc.keys) {
				t.Errorf("warm called %d times, want %d (every key must be attempted)", calls, len(tc.keys))
			}
		})
	}
}

func TestWarmCachesErrorIsUnwrappable(t *testing.T) {
	sentinel := errors.New("boom")
	err := WarmCaches([]string{"a"}, func(string) error { return sentinel })
	if !errors.Is(err, sentinel) {
		t.Fatalf("WarmCaches() = %v, want the warm error itself", err)
	}
}
