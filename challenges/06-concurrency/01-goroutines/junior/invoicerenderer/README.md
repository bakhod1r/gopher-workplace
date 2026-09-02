# Invoice Renderer

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A billing run renders a month of invoices as PDFs. Before rendering, each
invoice is totalled from its line items. Invoices are independent documents, so
the run totals them all concurrently and keeps the results in invoice order for
the render step.

## Task

Implement `InvoiceTotals` in [invoicerenderer.go](invoicerenderer.go) so that:

1. Return a slice with one total per invoice, in input order.
2. The total of an invoice is the sum of its `Lines`; an invoice with no lines totals `0`.
3. Total each invoice in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  InvoiceTotals([]Invoice{{[]int{100, 250}}})
Output: [350]
```

**Example 2:**

```
Input:  InvoiceTotals([]Invoice{{nil}})
Output: [0]
```

**Example 3:**

```
Input:  InvoiceTotals(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Concurrent reads are safe** | Each goroutine reads one invoice's lines and writes one index — reads never need locking. |

## Hint

Accumulate into a local `total` declared inside the goroutine, then write
`out[i]` once at the end.

## Validate

```bash
make verify
```
