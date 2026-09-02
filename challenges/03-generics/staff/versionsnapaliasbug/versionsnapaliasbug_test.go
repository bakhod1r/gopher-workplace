package versionsnapaliasbug

import (
	"reflect"
	"testing"
)

func TestRestoreTwice(t *testing.T) {
	var s Store[int]
	s.Append(1)
	s.Append(2)
	s.Append(3)
	id := s.Snapshot()

	s.Set(0, 99)
	if !s.Restore(id) {
		t.Fatal("first Restore = false, want true")
	}
	if got := s.Items(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("after first restore Items = %v, want [1 2 3]", got)
	}

	s.Set(0, 77)
	if !s.Restore(id) {
		t.Fatal("second Restore = false, want true")
	}
	if got := s.Items(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("after second restore Items = %v, want [1 2 3]", got)
	}
}

func TestRestoreThenAppend(t *testing.T) {
	var s Store[int]
	s.Append(1)
	s.Append(2)
	id := s.Snapshot()
	s.Append(3)
	if !s.Restore(id) {
		t.Fatal("Restore = false, want true")
	}
	s.Append(9)
	if !s.Restore(id) {
		t.Fatal("Restore = false, want true")
	}
	if got := s.Items(); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("Items = %v, want [1 2]", got)
	}
}

func TestRestoreBadID(t *testing.T) {
	var s Store[int]
	s.Append(1)
	if s.Restore(-1) {
		t.Error("Restore(-1) = true, want false")
	}
	if s.Restore(0) {
		t.Error("Restore(0) with no snapshots = true, want false")
	}
}
