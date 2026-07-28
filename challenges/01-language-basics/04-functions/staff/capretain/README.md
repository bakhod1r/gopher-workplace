# Sub-slice Retains Backing

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

`xs[:k]` shares the backing array and keeps its full capacity, so appending to
the head writes into `xs[k]`. A full-slice expression `xs[:k:k]` caps capacity
to k, forcing appends to allocate — and lets the rest of the array be freed.

## Task

Fix [capretain.go](capretain.go) so appends to the head don't touch the original.

Do **not** change the function signature or the tests.

## Examples

```go
Head([1 2 3 4 5], 2) then append(h,99) leaves xs[2]==3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Three-index slice** | `xs[:k:k]` sets len=cap=k. |
| 2 | **Capacity sharing** | Default re-slice keeps the parent's cap. |
| 3 | **Append safety + GC** | Capping cap isolates writes and frees the tail. |

## Hint

Cap the capacity: `return xs[:k:k]` (or copy into a fresh slice).

## Validate

```bash
make verify
```
