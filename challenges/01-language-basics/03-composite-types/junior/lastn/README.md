# Last N Elements

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Tailing a log or history buffer: return the last N entries, clamping when N
exceeds the length.

## Task

Implement `Last(xs, n)` — last `n` in order; clamp `n` to `[0, len]`.

## Examples

```go
Last([]int{1,2,3,4,5}, 2)  // => [4 5]
Last([]int{1,2,3}, 10)     // => [1 2 3]
Last([]int{1}, 0)          // => []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice expression** | `xs[len(xs)-n:]` takes a tail. |
| 2 | **Clamp bounds** | Guard n>len and n<=0. |
| 3 | **Empty result** | Return `[]int{}`, not a panic. |

## Hint

Clamp `n` to `[0, len(xs)]`, then `return xs[len(xs)-n:]`.

## Validate

```bash
make verify
```
