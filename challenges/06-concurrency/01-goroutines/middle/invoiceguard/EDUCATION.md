# Invoice Render Panic Guard

## Intuition

A `recover` in the parent cannot catch a child goroutine's panic — there is no stack relationship between them. The runtime unwinds only the panicking goroutine, and if nothing on *that* stack recovers, the whole process dies. Isolation therefore has to be installed at the top of every worker.

## Approach

1. Allocate `out := make([]Rendered, len(ids))`.
2. Launch a goroutine per ID with `i` and `id` as parameters.
3. `defer wg.Done()`, set `out[i].ID = id`, then register a deferred `recover` that fills `Doc` and `Err` on failure.
4. Call `render(id)` and assign the result to `out[i].Doc`.
5. `wg.Wait()` and return `out`.

## Solution

```go
// Package invoiceguard — Gopher Workplace challenge.
package invoiceguard

import (
	"fmt"
	"sync"
)

// Rendered is the outcome of rendering one invoice.
type Rendered struct {
	ID  string
	Doc string
	Err error
}

// RenderInvoices renders every invoice in its own goroutine and returns one
// Rendered per ID, in input order. A template that panics on one malformed
// invoice must not take the whole billing run down with it: the worker recovers
// the panic and reports it as that invoice's Err.
//
// Examples:
//
//	RenderInvoices([]string{"INV-1"}, render)         => [{INV-1 doc:INV-1 <nil>}]
//	RenderInvoices([]string{"INV-BAD"}, render)       => [{INV-BAD  panic}]
//	RenderInvoices(nil, render)                       => []
func RenderInvoices(ids []string, render func(id string) string) []Rendered {
	out := make([]Rendered, len(ids))
	var wg sync.WaitGroup
	for i, id := range ids {
		wg.Add(1)
		go func(i int, id string) {
			defer wg.Done()
			out[i].ID = id
			defer func() {
				if r := recover(); r != nil {
					out[i].Doc = ""
					out[i].Err = fmt.Errorf("render %s panicked: %v", id, r)
				}
			}()
			out[i].Doc = render(id)
		}(i, id)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- For `INV-1` no panic fires, the recover defer sees a nil value and does nothing, and `Doc` keeps `doc:INV-1`.
- For `INV-BAD` the assignment never happens, the deferred recover catches `"nil line items for INV-BAD"` and builds an error naming the invoice.
- In `first_panics_rest_survive` the run still returns three entries — the panic was confined to its own goroutine.
- Setting `ID` before the risky call means even a failed entry is attributable.

## Pitfalls

- Putting the `recover` in the parent around `wg.Wait()`: it catches nothing and the process still dies.
- Registering `defer wg.Done()` *after* the recover defer — `Done` then runs first and `Wait` can return before `Err` is written.
- Recovering but discarding the value, which turns a crash into a silent empty invoice.
- Letting a partially written `Doc` survive alongside a non-nil `Err`; a failed render must not ship a half-built document.
