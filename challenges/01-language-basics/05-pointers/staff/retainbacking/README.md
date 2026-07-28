# Sub-slice Retains Full Array

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _memory-management_

## Context

`xs[:k]` keeps the parent's capacity, so appending to the prefix writes into
`xs[k]`. A full-slice expression `xs[:k:k]` caps capacity to k, isolating appends
and letting the tail be garbage-collected.

## Task

Fix [retainbacking.go](retainbacking.go) so appends to the prefix don't touch the parent.

Do **not** change the function signature or the tests.

## Examples

```go
Prefix([1 2 3 4 5], 2); append(p,99) // xs[2] stays 3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Three-index slice** | `xs[:k:k]` sets cap=k. |
| 2 | **Capacity sharing** | Default re-slice keeps parent cap. |
| 3 | **GC + safety** | Capping isolates writes and frees the tail. |

## Hint

Cap the capacity: `return xs[:k:k]` (or copy into a fresh slice).

## Validate

```bash
make verify
```
