package paymentworkers

import (
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
)

func TestCaptureAll(t *testing.T) {
	capture := func(charge string) string {
		if strings.HasSuffix(charge, "_bad") {
			return "declined"
		}
		return "captured"
	}

	cases := []struct {
		name    string
		charges []string
		workers int
		want    map[string]string
	}{
		{"single", []string{"ch_1"}, 2, map[string]string{"ch_1": "captured"}},
		{"one_worker", []string{"ch_1", "ch_2"}, 1, map[string]string{"ch_1": "captured", "ch_2": "captured"}},
		{"mixed", []string{"ch_1", "ch_2_bad"}, 3, map[string]string{"ch_1": "captured", "ch_2_bad": "declined"}},
		{"empty", nil, 4, map[string]string{}},
		{"more_workers_than_jobs", []string{"ch_1"}, 8, map[string]string{"ch_1": "captured"}},
		{"zero_workers_means_one", []string{"ch_1", "ch_2"}, 0, map[string]string{"ch_1": "captured", "ch_2": "captured"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CaptureAll(tc.charges, tc.workers, capture)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CaptureAll(%v, %d) = %v, want %v", tc.charges, tc.workers, got, tc.want)
			}
		})
	}
}

func TestEveryChargeIsCapturedExactlyOnce(t *testing.T) {
	const n = 200
	charges := make([]string, n)
	for i := range charges {
		charges[i] = "ch_" + strconv.Itoa(i)
	}

	var calls atomic.Int64
	got := CaptureAll(charges, 8, func(charge string) string {
		calls.Add(1)
		return "captured:" + charge
	})

	if calls.Load() != n {
		t.Errorf("capture called %d times, want %d", calls.Load(), n)
	}
	if len(got) != n {
		t.Fatalf("len(results) = %d, want %d", len(got), n)
	}
	for _, c := range charges {
		if got[c] != "captured:"+c {
			t.Fatalf("result[%q] = %q, want %q", c, got[c], "captured:"+c)
		}
	}
}

func TestWorkerLimitIsRespected(t *testing.T) {
	const n, workers = 120, 4

	charges := make([]string, n)
	for i := range charges {
		charges[i] = "ch_" + strconv.Itoa(i)
	}

	var live, peak atomic.Int64
	CaptureAll(charges, workers, func(charge string) string {
		cur := live.Add(1)
		for {
			old := peak.Load()
			if cur <= old || peak.CompareAndSwap(old, cur) {
				break
			}
		}
		live.Add(-1)
		return "captured"
	})

	if peak.Load() > workers {
		t.Errorf("peak concurrency = %d, want <= %d", peak.Load(), workers)
	}
}
