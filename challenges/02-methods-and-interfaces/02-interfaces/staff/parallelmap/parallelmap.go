// Package parallelmap — Gopher Workplace challenge.
package parallelmap

import (
	"runtime"
	"sync"
)

// Op transforms one value.
type Op interface {
	Apply(v int) int
}

// SquareOp squares its input.
type SquareOp struct{}

// Apply returns v*v.
func (SquareOp) Apply(v int) int {
	// TODO(candidate): square it.
	panic("not implemented")
}

// MapSeq applies op to every element sequentially.
func MapSeq(op Op, vs []int) []int {
	out := make([]int, len(vs))
	for i, v := range vs {
		out[i] = op.Apply(v)
	}
	return out
}

// MapParallel applies op across GOMAXPROCS workers, preserving order.
//
// Examples:
//
//	MapParallel(SquareOp{}, []int{1, 2, 3}) => [1 4 9]
func MapParallel(op Op, vs []int) []int {
	// TODO(candidate): chunk the range, one worker per chunk.
	panic("not implemented")
}

var (
	_ = runtime.GOMAXPROCS
	_ = sync.WaitGroup{}
)
