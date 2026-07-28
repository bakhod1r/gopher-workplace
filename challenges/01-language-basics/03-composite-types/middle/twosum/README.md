# Two Sum

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Find two elements adding to a target — the canonical map-as-index-lookup.

## Task

Implement `TwoSum(xs, target)` returning indices `i<j`, or ok=false.

## Examples

```go
TwoSum([]int{2,7,11,15}, 9) // => 0, 1, true
TwoSum([]int{3,2,4}, 6)     // => 1, 2, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Complement lookup** | Need `target - xs[i]`. |
| 2 | **Map value→index** | Remember where each value was. |
| 3 | **Single pass** | Check before inserting. |

## Hint

For each i, if `seen[target-xs[i]]` exists return it; else `seen[xs[i]] = i`.

## Validate

```bash
make verify
```
