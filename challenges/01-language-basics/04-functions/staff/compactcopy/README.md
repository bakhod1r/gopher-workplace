# Overlapping Copy Shift

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

To drop index 0, the tail `xs[1:]` must move LEFT into `xs[:len-1]`, i.e.
`copy(xs, xs[1:])`. The bug copies `xs[:len-1]` into `xs[1:]`, shifting right and
duplicating the first element.

## Task

Fix the copy in [compactcopy.go](compactcopy.go).

Do **not** change the function signature or the tests.

## Examples

```go
DropFirst([10 20 30 40]) // => [20 30 40]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Compacting shift** | Move the tail left with `copy(xs, xs[1:])`. |
| 2 | **copy overlap direction** | Source and destination overlap; direction matters. |
| 3 | **Trim length** | Re-slice to drop the last slot. |

## Hint

Shift left: `copy(xs, xs[1:])`, then return `xs[:len(xs)-1]`.

## Validate

```bash
make verify
```
