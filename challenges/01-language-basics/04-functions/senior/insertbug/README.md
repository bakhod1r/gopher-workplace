# Insert Shifts the Wrong Way

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

To open a gap at index i, the tail `xs[i:]` must shift RIGHT into `xs[i+1:]`.
The bug copies `xs[i+1:]` back onto `xs[i:]`, shifting left and clobbering the
insertion point.

## Task

Fix the shift in [insertbug.go](insertbug.go).

Do **not** change the function signature or the tests.

## Examples

```go
InsertAt([1 2 4], 2, 3) // => [1 2 3 4]
InsertAt([2 3], 0, 1)   // => [1 2 3]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Open a gap** | Shift the tail right by one. |
| 2 | **copy direction** | `copy(xs[i+1:], xs[i:])` moves right. |
| 3 | **Overlapping copy** | Go's copy handles overlap correctly for this direction. |

## Hint

Shift right: `copy(xs[i+1:], xs[i:])`, then set `xs[i] = v`.

## Validate

```bash
make verify
```
