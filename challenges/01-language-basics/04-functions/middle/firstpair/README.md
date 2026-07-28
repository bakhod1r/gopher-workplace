# Labeled Break Search

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

A labeled `break` exits an outer loop from inside a nested one — cleaner than a
found-flag checked at every level.

## Task

Implement `FindPairSum` in [firstpair.go](firstpair.go) using labeled break.

Do **not** change the function signature or the tests.

## Examples

```go
FindPairSum([]int{1,2,3,4}, 6) // => 1, 3, true
FindPairSum([]int{1,2}, 100)   // => 0, 0, false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Labeled break** | `Outer: for ... { for ... { break Outer } }`. |
| 2 | **Nested search** | Scan all i<j pairs. |
| 3 | **Result flags** | Set i,j,ok then break out. |

## Hint

Label the outer loop; on a match set `i,j,ok` and `break Outer`.

## Validate

```bash
make verify
```
