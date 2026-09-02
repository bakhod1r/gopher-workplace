# Billing Run Totals

## Intuition

Two atomic counters would each be safe alone, yet a reader could still catch the sum updated and the count not. One mutex around both fields makes the pair move as a single unit, so every observer sees a consistent report.

## Approach

1. `Add` locks, does `sum += amount` and `count++`, unlocks.
2. `Average` locks and reads both fields into locals.
3. Return 0 when the count is zero, otherwise `sum / count`.

## Solution

```go
// Package billingtotals - Gopher Workplace challenge.
package billingtotals

import "sync"

// Totals accumulates invoice amounts during a billing run.
type Totals struct {
	mu    sync.Mutex
	sum   int64
	count int64
}

// Add records one invoice amount.
//
// Examples:
//
//	var t Totals; t.Add(300); t.Total() => 300
func (t *Totals) Add(amount int64) {
	t.mu.Lock()
	t.sum += amount
	t.count++
	t.mu.Unlock()
}

// Total returns the summed invoice amount.
func (t *Totals) Total() int64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.sum
}

// Count returns how many invoices were added.
func (t *Totals) Count() int64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.count
}

// Average returns the truncated mean amount, or 0 when nothing was added.
//
// Examples:
//
//	var t Totals; t.Add(300); t.Add(100); t.Average() => 200
//	var t Totals; t.Average()                         => 0
func (t *Totals) Average() int64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.count == 0 {
		return 0
	}
	return t.sum / t.count
}
```

## Walkthrough

Two workers add 300 and 100. Each holds the lock for both writes, so after the run `sum == 400` and `count == 2`, and `Average` reports 200 from a single consistent read.

## Pitfalls

- Using two separate locks or two separate atomics, which lets the pair drift apart.
- Reading `Total()` and `Count()` separately in `Average` — that is two lock holds with a gap between them.
- Dividing before the zero-count guard, which panics.
