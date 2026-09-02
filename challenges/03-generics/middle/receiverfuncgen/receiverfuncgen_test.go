package receiverfuncgen

import (
	"strconv"
	"testing"
)

func TestUpdate(t *testing.T) {
	b := &Box[int]{value: 2}
	Update(b, func(n int) int { return n * 3 })
	if got := b.Get(); got != 6 {
		t.Errorf("Get() = %v, want 6", got)
	}
}

func TestConvert(t *testing.T) {
	b := &Box[int]{value: 7}
	s := Convert(b, strconv.Itoa)
	if got := s.Get(); got != "7" {
		t.Errorf("Get() = %q, want 7", got)
	}
	if got := b.Get(); got != 7 {
		t.Errorf("source box changed: %v, want 7", got)
	}
}

func TestUpdateThenConvert(t *testing.T) {
	b := &Box[int]{value: 1}
	Update(b, func(n int) int { return n + 1 })
	if got := Convert(b, strconv.Itoa).Get(); got != "2" {
		t.Errorf("Get() = %q, want 2", got)
	}
}
