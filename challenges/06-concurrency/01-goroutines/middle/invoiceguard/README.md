# Invoice Render Panic Guard

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

The monthly billing run renders thousands of invoices through a template engine. One customer with a malformed line item makes that template panic — and a panic in a goroutine that nobody recovers kills the entire process, losing every invoice that had already rendered. Each worker therefore owns its own failure.

## Task

Implement the exported function(s) in [invoiceguard.go](invoiceguard.go) so that:

1. Return one `Rendered` per ID, in input order, with `ID` always populated.
2. Render each invoice in its own goroutine, joined with a `sync.WaitGroup`.
3. Recover a panic inside the goroutine that raised it, and record it as that entry's `Err`.
4. The error text must name the invoice ID and include the panic value.
5. Leave `Doc` empty for a failed render; a panic in one invoice must not affect any other.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  RenderInvoices([]string{"INV-1", "INV-2"}, render)
Output: [{INV-1 doc:INV-1 <nil>} {INV-2 doc:INV-2 <nil>}]
```

**Example 2:**

```
Input:  RenderInvoices([]string{"INV-BAD", "INV-2"}, render)
Output: [{INV-BAD "" render INV-BAD panicked: ...} {INV-2 doc:INV-2 <nil>}]
```

**Example 3:**

```
Input:  RenderInvoices(nil, render)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Panic isolation** | `recover` only works in a deferred function running on the *same* goroutine that panicked. |
| 2 | **Deferred ordering** | Defers run last-in-first-out, so the recover must be registered after `wg.Done` to run before it. |
| 3 | **Goroutine ownership** | Each worker owns one slot and one failure; nothing leaks to a sibling. |
| 4 | **`fmt.Errorf` with `%v`** | The recovered value is an `any`; format it into a real error the caller can inspect. |

## Hint

Register `defer wg.Done()` first, then a second `defer` holding the `recover()`. The second one runs first, so `Err` is set before the WaitGroup is released.

## Validate

```bash
make verify
```
