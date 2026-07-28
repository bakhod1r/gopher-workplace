# Early Return Placement

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

A universal claim ("all positive") is disproved by a single counterexample, so
the loop must return false on the FIRST non-positive and only return true after
checking everything. The bug returns true on the first positive — inverted logic.

## Task

Fix [allpositive.go](allpositive.go) so it checks every element.

Do **not** change the function signature or the tests.

## Examples

```go
AllPositive([]int{1,2,3})  // => true
AllPositive([]int{1,-2,3}) // => false
AllPositive(nil)           // => true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Universal quantifier** | Return false on the first failure, true at the end. |
| 2 | **Early return** | Exit early only for the disqualifying case. |
| 3 | **Empty case** | Vacuously true. |

## Hint

Return false on the first `v <= 0`; keep the final `return true` for when the loop completes.

## Validate

```bash
make verify
```
