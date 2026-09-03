package webhookenqueue

import (
	"reflect"
	"testing"
)

func batchOf(ids ...string) []Delivery {
	out := make([]Delivery, 0, len(ids))
	for _, id := range ids {
		out = append(out, Delivery{ID: id, Endpoint: "https://hooks.example/" + id})
	}
	return out
}

func TestEnqueueDeliveries(t *testing.T) {
	cases := []struct {
		name         string
		capacity     int
		prefill      int
		ids          []string
		wantAccepted []string
		wantDropped  []string
	}{
		{"queue_fills_up", 2, 0, []string{"a", "b", "c"}, []string{"a", "b"}, []string{"c"}},
		{"room_for_all", 4, 0, []string{"a"}, []string{"a"}, []string{}},
		{"unbuffered_no_receiver", 0, 0, []string{"a"}, []string{}, []string{"a"}},
		{"already_partly_full", 3, 2, []string{"a", "b"}, []string{"a"}, []string{"b"}},
		{"empty_batch", 2, 0, nil, []string{}, []string{}},
		{"already_full", 1, 1, []string{"a", "b"}, []string{}, []string{"a", "b"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			queue := make(chan Delivery, tc.capacity)
			for i := 0; i < tc.prefill; i++ {
				queue <- Delivery{ID: "old"}
			}

			accepted, dropped := EnqueueDeliveries(queue, batchOf(tc.ids...))
			if !reflect.DeepEqual(accepted, tc.wantAccepted) {
				t.Errorf("accepted = %#v, want %#v", accepted, tc.wantAccepted)
			}
			if !reflect.DeepEqual(dropped, tc.wantDropped) {
				t.Errorf("dropped = %#v, want %#v", dropped, tc.wantDropped)
			}
			if got := len(queue); got != tc.prefill+len(tc.wantAccepted) {
				t.Errorf("queue holds %d, want %d", got, tc.prefill+len(tc.wantAccepted))
			}
		})
	}
}
