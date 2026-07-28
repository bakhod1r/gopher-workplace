# Count Distinct

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A set is a map whose keys are the members. Counting distinct values is inserting
each into a set and taking its size.

## Task

Implement `Distinct(xs)`.

## Examples

```go
Distinct([]int{1,2,2,3,3,3}) // => 3
Distinct([]int{5,5,5})       // => 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Set as map** | `map[int]struct{}` keys are members. |
| 2 | **Idempotent insert** | Re-inserting a key is a no-op. |
| 3 | **Size = len** | `len(set)` is the distinct count. |

## Hint

`seen := make(map[int]struct{}); for _, x := range xs { seen[x] = struct{}{} }; return len(seen)`.

## Validate

```bash
make verify
```
