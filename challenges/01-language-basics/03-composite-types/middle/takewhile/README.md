# Take While Positive

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Reading a leading run until a condition breaks — like parsing a prefix.

## Task

Implement `TakePositive(xs)` returning the longest all-positive prefix.

## Examples

```go
TakePositive([]int{1,2,3,-1,4}) // => [1 2 3]
TakePositive([]int{-1,2})       // => []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Prefix scan** | Stop at first failure. |
| 2 | **Slice up to i** | `xs[:i]` is the taken prefix. |
| 3 | **Empty result** | Return non-nil empty. |

## Hint

Find the first index where `xs[i] <= 0`, return `xs[:i]` (or all).

## Validate

```bash
make verify
```
