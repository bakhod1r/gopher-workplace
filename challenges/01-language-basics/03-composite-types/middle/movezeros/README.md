# Move Zeros

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Compacting non-zero values to the front (stable), zeros to the back — a common
array-shuffling task.

## Task

Implement `MoveZeros(xs)` (stable order of non-zeros).

## Examples

```go
MoveZeros([]int{0,1,0,3,12}) // => [1 3 12 0 0]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Stable compaction** | Keep non-zero order. |
| 2 | **Two-phase build** | Non-zeros, then zeros. |
| 3 | **Length preserved** | Output same length. |

## Hint

Append non-zeros, then append `len(xs)-count` zeros.

## Validate

```bash
make verify
```
