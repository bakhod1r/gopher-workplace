package sessiongc

import (
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func TestExpire(t *testing.T) {
	cases := []struct {
		name    string
		touches map[string]int64
		cutoff  int64
		want    []string
		left    int
	}{
		{"one_expires", map[string]int64{"u1": 1, "u2": 9}, 5, []string{"u1"}, 1},
		{"none_expire", map[string]int64{"u1": 7, "u2": 9}, 5, []string{}, 2},
		{"all_expire", map[string]int64{"u1": 1, "u2": 2}, 5, []string{"u1", "u2"}, 0},
		{"empty_store", nil, 5, []string{}, 0},
		{"cutoff_is_exclusive", map[string]int64{"u1": 5}, 5, []string{}, 1},
		{"sorted_output", map[string]int64{"c": 1, "a": 1, "b": 1}, 2, []string{"a", "b", "c"}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s Store
			for id, tick := range tc.touches {
				s.Touch(id, tick)
			}
			got := s.Expire(tc.cutoff)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expire(%d) = %v, want %v", tc.cutoff, got, tc.want)
			}
			if left := s.Active(); left != tc.left {
				t.Errorf("Active() = %d, want %d", left, tc.left)
			}
		})
	}
}

func TestTouchUpdatesLastSeen(t *testing.T) {
	var s Store
	s.Touch("u1", 3)
	s.Touch("u1", 11)

	got, ok := s.LastSeen("u1")
	if !ok || got != 11 {
		t.Errorf("LastSeen(u1) = %d, %v; want 11, true", got, ok)
	}
	if _, ok := s.LastSeen("nobody"); ok {
		t.Error("LastSeen(unknown) reported a live session")
	}
	if n := s.Active(); n != 1 {
		t.Errorf("Active() = %d, want 1", n)
	}
}

func TestTouchWhileSweeping(t *testing.T) {
	var s Store
	const users = 200

	for i := range users {
		s.Touch("u"+strconv.Itoa(i), 1)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range users {
			s.Touch("fresh"+strconv.Itoa(i), 100)
		}
	}()
	go func() {
		defer wg.Done()
		s.Expire(50)
	}()
	wg.Wait()

	for i := range users {
		if _, ok := s.LastSeen("fresh" + strconv.Itoa(i)); !ok {
			t.Fatalf("fresh session %d was swept", i)
		}
	}
}
