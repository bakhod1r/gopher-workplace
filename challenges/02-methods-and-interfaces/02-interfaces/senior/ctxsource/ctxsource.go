// Package ctxsource — Gopher Workplace challenge.
package ctxsource

import "context"

// Source yields values until drained.
type Source interface {
	Next() (int, bool)
}

// RangeSource yields 1..N.
type RangeSource struct {
	N   int
	pos int
}

// Next yields the next value.
func (r *RangeSource) Next() (int, bool) {
	// TODO(candidate): yield 1..N, then report drained.
	panic("not implemented")
}

// SumWithContext sums src, aborting as soon as ctx is done.
//
// Examples:
//
//	live context over 1..3   => 6, nil
//	cancelled context        => 0, context.Canceled
func SumWithContext(ctx context.Context, src Source) (int, error) {
	// TODO(candidate): check cancellation on every iteration.
	panic("not implemented")
}
