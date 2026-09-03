package mailsemaphore

import (
	"reflect"
	"strconv"
	"sync/atomic"
	"testing"
)

func TestSendAll(t *testing.T) {
	send := func(m string) string { return "sent:" + m }

	cases := []struct {
		name     string
		messages []string
		limit    int
		want     []string
	}{
		{"serial", []string{"m1", "m2"}, 1, []string{"sent:m1", "sent:m2"}},
		{"limit_above_count", []string{"m1"}, 5, []string{"sent:m1"}},
		{"empty", nil, 2, []string{}},
		{"unlimited", []string{"m1", "m2", "m3"}, 0, []string{"sent:m1", "sent:m2", "sent:m3"}},
		{"negative_limit", []string{"m1", "m2"}, -3, []string{"sent:m1", "sent:m2"}},
		{"order_preserved", []string{"b", "a", "c"}, 2, []string{"sent:b", "sent:a", "sent:c"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SendAll(tc.messages, tc.limit, send)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SendAll(%v, %d) = %v, want %v", tc.messages, tc.limit, got, tc.want)
			}
		})
	}
}

func TestEveryMessageIsSentOnce(t *testing.T) {
	const n = 250
	msgs := make([]string, n)
	for i := range msgs {
		msgs[i] = "m" + strconv.Itoa(i)
	}

	var calls atomic.Int64
	got := SendAll(msgs, 6, func(m string) string {
		calls.Add(1)
		return "sent:" + m
	})

	if calls.Load() != n {
		t.Errorf("send called %d times, want %d", calls.Load(), n)
	}
	for i, s := range got {
		if want := "sent:m" + strconv.Itoa(i); s != want {
			t.Fatalf("results[%d] = %q, want %q", i, s, want)
		}
	}
}

func TestConcurrencyNeverExceedsLimit(t *testing.T) {
	const n, limit = 200, 5

	msgs := make([]string, n)
	for i := range msgs {
		msgs[i] = "m" + strconv.Itoa(i)
	}

	var live, peak atomic.Int64
	SendAll(msgs, limit, func(m string) string {
		cur := live.Add(1)
		for {
			old := peak.Load()
			if cur <= old || peak.CompareAndSwap(old, cur) {
				break
			}
		}
		live.Add(-1)
		return "sent:" + m
	})

	if got := peak.Load(); got > limit {
		t.Errorf("peak concurrency = %d, want <= %d", got, limit)
	}
}
