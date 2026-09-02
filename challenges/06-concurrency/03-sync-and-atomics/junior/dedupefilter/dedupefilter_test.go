package dedupefilter

import (
	"strconv"
	"sync"
	"testing"
)

func TestDedupeFilter(t *testing.T) {
	cases := []struct {
		name      string
		delivered []string
		eventID   string
		wantFirst bool
		wantLen   int
	}{
		{"first_delivery", []string{}, "evt-1", true, 1},
		{"redelivery", []string{"evt-1"}, "evt-1", false, 1},
		{"distinct_event", []string{"evt-1"}, "evt-2", true, 2},
		{"new_after_many", []string{"evt-1", "evt-2", "evt-3"}, "evt-4", true, 4},
		{"redelivery_of_many", []string{"evt-1", "evt-2", "evt-3"}, "evt-2", false, 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var d DedupeFilter
			for _, id := range tc.delivered {
				d.Accept(id)
			}
			if got := d.Accept(tc.eventID); got != tc.wantFirst {
				t.Errorf("Accept(%q) = %v, want %v", tc.eventID, got, tc.wantFirst)
			}
			if !d.Seen(tc.eventID) {
				t.Errorf("Seen(%q) = false, want true", tc.eventID)
			}
			if got := d.Len(); got != tc.wantLen {
				t.Errorf("Len() = %d, want %d", got, tc.wantLen)
			}
		})
	}
}

func TestDedupeFilterConcurrent(t *testing.T) {
	var d DedupeFilter
	const connections = 16
	const events = 25
	accepted := make(chan int, connections*events)
	var wg sync.WaitGroup
	wg.Add(connections)
	for i := 0; i < connections; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < events; j++ {
				if d.Accept(strconv.Itoa(j)) {
					accepted <- j
				}
			}
		}()
	}
	wg.Wait()
	close(accepted)

	if got := len(accepted); got != events {
		t.Errorf("processed events = %d, want %d", got, events)
	}
	if got := d.Len(); got != events {
		t.Errorf("Len() = %d, want %d", got, events)
	}
}
