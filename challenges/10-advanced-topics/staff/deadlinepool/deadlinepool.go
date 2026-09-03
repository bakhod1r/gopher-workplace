// Package deadlinepool — Gopher Workplace challenge.
package deadlinepool

import (
	"context"
	"sync"
)

// Process doubles every item using workers goroutines and returns the
// results in input order.
//
// If ctx is cancelled or its deadline passes first, Process returns the
// context's error, and every goroutine it started must have exited.
//
// Examples:
//
//	Process(ctx, []int{1, 2}, 2) => []int{2, 4}, nil
func Process(ctx context.Context, items []int, workers int) ([]int, error) {
	panic("not implemented")
}
