# Drop While

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

DropWhile finds the first element failing the predicate and returns everything
from there — the complement of TakeWhile.

## Task

Implement `DropWhile` in [dropwhile.go](dropwhile.go). Return a fresh slice (copy the tail).

Do **not** change the function signature or the tests.

## Examples

```go
DropWhile([]int{1,2,-1,3}, pos) // => [-1 3]
DropWhile([]int{1,2}, pos)      // => []
DropWhile([]int{-1}, pos)       // => [-1]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Find boundary** | Advance while pred holds. |
| 2 | **Copy the tail** | Return a new slice to avoid aliasing. |
| 3 | **Complements TakeWhile** | Drops the prefix TakeWhile keeps. |

## Hint

Advance `i` while `i < len(xs) && pred(xs[i])`; return a copy of `xs[i:]`.

## Validate

```bash
make verify
```
