package shipmentpipeline

import (
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
)

func label(order string) string { return "LABEL-" + order }

func TestLabels(t *testing.T) {
	all := func(string) bool { return true }
	notBad := func(o string) bool { return !strings.HasPrefix(o, "bad") }

	cases := []struct {
		name   string
		orders []string
		keep   func(string) bool
		want   []string
	}{
		{"all_kept", []string{"o1", "o2"}, all, []string{"LABEL-o1", "LABEL-o2"}},
		{"one_dropped", []string{"o1", "bad1"}, notBad, []string{"LABEL-o1"}},
		{"all_dropped", []string{"bad1", "bad2"}, notBad, []string{}},
		{"empty", nil, all, []string{}},
		{"order_preserved", []string{"o3", "o1", "o2"}, all, []string{"LABEL-o3", "LABEL-o1", "LABEL-o2"}},
		{"drop_in_the_middle", []string{"o1", "bad", "o2"}, notBad, []string{"LABEL-o1", "LABEL-o2"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Labels(tc.orders, tc.keep, label)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Labels(%v) = %v, want %v", tc.orders, got, tc.want)
			}
		})
	}
}

func TestRejectedOrdersNeverReachRender(t *testing.T) {
	var rendered atomic.Int64
	got := Labels(
		[]string{"o1", "bad1", "bad2", "o2"},
		func(o string) bool { return !strings.HasPrefix(o, "bad") },
		func(o string) string {
			rendered.Add(1)
			return label(o)
		},
	)

	if rendered.Load() != 2 {
		t.Errorf("render called %d times, want 2", rendered.Load())
	}
	if want := []string{"LABEL-o1", "LABEL-o2"}; !reflect.DeepEqual(got, want) {
		t.Errorf("labels = %v, want %v", got, want)
	}
}

func TestLargeBatchStaysInOrder(t *testing.T) {
	const n = 300
	orders := make([]string, n)
	for i := range orders {
		orders[i] = "o" + strconv.Itoa(i)
	}

	got := Labels(orders, func(o string) bool { return true }, label)
	if len(got) != n {
		t.Fatalf("len = %d, want %d", len(got), n)
	}
	for i := range got {
		if want := "LABEL-o" + strconv.Itoa(i); got[i] != want {
			t.Fatalf("labels[%d] = %q, want %q", i, got[i], want)
		}
	}
}
