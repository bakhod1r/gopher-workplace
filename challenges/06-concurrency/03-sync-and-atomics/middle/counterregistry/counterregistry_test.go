package counterregistry

import (
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name  string
		steps []int64
		want  int64
	}{
		{"single", []int64{1}, 1},
		{"accumulates", []int64{1, 2}, 3},
		{"zero", []int64{0}, 0},
		{"negative", []int64{5, -2}, 3},
		{"many", []int64{1, 1, 1, 1, 1}, 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := NewRegistry()
			var got int64
			for _, n := range tc.steps {
				got = r.Add("http_requests", n)
			}
			if got != tc.want {
				t.Errorf("Add total = %d, want %d", got, tc.want)
			}
			if v := r.Value("http_requests"); v != tc.want {
				t.Errorf("Value() = %d, want %d", v, tc.want)
			}
		})
	}
}

func TestUnknownNameIsZero(t *testing.T) {
	r := NewRegistry()
	if got := r.Value("unknown"); got != 0 {
		t.Errorf("Value(unknown) = %d, want 0", got)
	}
}

func TestSnapshotIsACopy(t *testing.T) {
	r := NewRegistry()
	r.Add("a", 2)
	r.Add("b", 5)

	snap := r.Snapshot()
	want := map[string]int64{"a": 2, "b": 5}
	if !reflect.DeepEqual(snap, want) {
		t.Fatalf("Snapshot() = %v, want %v", snap, want)
	}

	snap["a"] = 99
	r.Add("c", 1)
	if got := r.Value("a"); got != 2 {
		t.Errorf("mutating the snapshot changed the registry: a = %d", got)
	}
	if _, ok := snap["c"]; ok {
		t.Error("snapshot saw a counter registered after it was taken")
	}
}

func TestConcurrentRegistrationAndIncrement(t *testing.T) {
	r := NewRegistry()
	const names, perName = 8, 250

	var wg sync.WaitGroup
	for i := range names {
		for range perName {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				r.Add("metric_"+strconv.Itoa(i), 1)
			}(i)
		}
	}
	wg.Wait()

	for i := range names {
		name := "metric_" + strconv.Itoa(i)
		if got := r.Value(name); got != perName {
			t.Errorf("Value(%q) = %d, want %d", name, got, perName)
		}
	}
	if got := len(r.Snapshot()); got != names {
		t.Errorf("len(Snapshot()) = %d, want %d", got, names)
	}
}
