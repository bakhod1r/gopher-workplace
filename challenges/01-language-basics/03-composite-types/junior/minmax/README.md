# Min and Max

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A single pass finds both extremes — seed from the first element, then compare.

## Task

Implement `MinMax(xs)` returning min, max, and `ok=false` for empty.

## Examples

```go
MinMax([]int{3,1,4,1,5}) // => 1, 5, true
MinMax(nil)              // => 0, 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Seed from first** | Initialize both to `xs[0]`. |
| 2 | **Single pass** | Update min/max together. |
| 3 | **Empty guard** | Return ok=false, avoid indexing xs[0]. |

## Hint

Guard empty; set `min=max=xs[0]`; loop the rest updating both.

## Validate

```bash
make verify
```
