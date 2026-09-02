package ctxsource

import (
	"context"
	"errors"
	"testing"
)

type sliceSource struct {
	data []int
	pos  int
}

func (s *sliceSource) Next() (int, bool) {
	if s.pos >= len(s.data) {
		return 0, false
	}
	v := s.data[s.pos]
	s.pos++
	return v, true
}

// cancelAt cancels the context after n reads.
type cancelAt struct {
	inner  Source
	n      int
	reads  int
	cancel context.CancelFunc
}

func (c *cancelAt) Next() (int, bool) {
	c.reads++
	if c.reads > c.n {
		c.cancel()
	}
	return c.inner.Next()
}

func TestSumLiveContext(t *testing.T) {
	got, err := SumWithContext(context.Background(), &sliceSource{data: []int{1, 2, 3}})
	if err != nil || got != 6 {
		t.Errorf("Sum = %d, %v; want 6, nil", got, err)
	}
}

func TestSumEmpty(t *testing.T) {
	got, err := SumWithContext(context.Background(), &sliceSource{})
	if err != nil || got != 0 {
		t.Errorf("Sum = %d, %v; want 0, nil", got, err)
	}
}

func TestCancelledUpFront(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	src := &sliceSource{data: []int{1, 2, 3}}
	got, err := SumWithContext(ctx, src)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	if got != 0 {
		t.Errorf("sum = %d, want 0", got)
	}
	if src.pos != 0 {
		t.Errorf("read %d values from a cancelled scan, want 0", src.pos)
	}
}

func TestCancelMidScan(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	src := &cancelAt{inner: &RangeSource{N: 10_000_000}, n: 5, cancel: cancel}

	_, err := SumWithContext(ctx, src)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	if src.reads > 100 {
		t.Errorf("kept reading after cancellation: %d reads", src.reads)
	}
}

func TestRangeSource(t *testing.T) {
	r := &RangeSource{N: 2}
	if v, ok := r.Next(); v != 1 || !ok {
		t.Errorf("Next = %d, %v", v, ok)
	}
	if v, ok := r.Next(); v != 2 || !ok {
		t.Errorf("Next = %d, %v", v, ok)
	}
	if _, ok := r.Next(); ok {
		t.Error("drained Next should report false")
	}
}
