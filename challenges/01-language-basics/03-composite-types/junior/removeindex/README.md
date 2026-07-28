# Remove by Index

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Deleting an element while keeping order uses the append-of-two-halves idiom.

## Task

Implement `RemoveAt(xs, i)`; out-of-range `i` returns `xs` unchanged.

## Examples

```go
RemoveAt([]int{1,2,3,4}, 1) // => [1 3 4]
RemoveAt([]int{1,2,3}, 5)   // => [1 2 3]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two-half append** | `append(xs[:i], xs[i+1:]...)`. |
| 2 | **Bounds guard** | Check `0 <= i < len`. |
| 3 | **Variadic spread** | `...` expands the tail. |

## Hint

Guard bounds, then `return append(xs[:i], xs[i+1:]...)`.

## Validate

```bash
make verify
```
