# Invoice Renderer

## Intuition

Reading shared data from many goroutines is safe; only writes need coordination.
Every goroutine here reads its own invoice and writes its own slot.

## Approach

1. Allocate `out := make([]int, len(invoices))`.
2. Launch one goroutine per invoice, passing `i` and the `Invoice`.
3. Sum `inv.Lines` into a local accumulator and store it at `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package invoicerenderer — Gopher Workplace challenge.
package invoicerenderer

import (
	"sync"
)

// Invoice is a billing document made of line-item amounts in cents.
type Invoice struct {
	Lines []int
}

// InvoiceTotals returns the total of every invoice, in input order.
//
// Examples:
//
//	InvoiceTotals([]Invoice{{[]int{100, 250}}})  => [350]
//	InvoiceTotals([]Invoice{{nil}})              => [0]
//	InvoiceTotals(nil)                           => []
func InvoiceTotals(invoices []Invoice) []int {
	out := make([]int, len(invoices))
	var wg sync.WaitGroup
	for i, inv := range invoices {
		wg.Add(1)
		go func(i int, inv Invoice) {
			defer wg.Done()
			total := 0
			for _, amount := range inv.Lines {
				total += amount
			}
			out[i] = total
		}(i, inv)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- An invoice with lines `100` and `250` totals `350`.
- An invoice with a `nil` line slice never enters the loop and totals `0`.
- A credit note of `-200` against `500` totals `300`.

## Pitfalls

- Summing into one shared variable — a race, and it answers the wrong question.
- Capturing `inv` from the loop variable instead of passing it as a parameter.
- Assuming every invoice has at least one line.
