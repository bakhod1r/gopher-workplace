# Dedupe Consecutive

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Collapsing runs of equal values (e.g. deduping a sorted list) in one pass.

## Task

Implement `Dedupe(xs)` collapsing consecutive duplicates.

## Examples

```go
Dedupe([]int{1,1,2,3,3,3,4}) // => [1 2 3 4]
Dedupe([]int{5,5,5})         // => [5]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Track previous** | Keep the last emitted value. |
| 2 | **Compare & skip** | Skip when equal to previous. |
| 3 | **First element** | Always keep the first. |

## Hint

Keep the first; append `xs[i]` only when it differs from `xs[i-1]`.

## Validate

```bash
make verify
```
