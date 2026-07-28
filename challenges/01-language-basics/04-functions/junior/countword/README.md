# Count Matching

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

Counting matches is a loop with a conditional increment — the smallest useful
filter.

## Task

Implement `CountEqual` in [countword.go](countword.go).

Do **not** change the function signature or the tests.

## Examples

```go
CountEqual([]int{1,2,1,3,1}, 1) // => 3
CountEqual(nil, 1)              // => 0
CountEqual([]int{2,2}, 5)       // => 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Loop + condition** | Increment when the element matches. |
| 2 | **Equality** | `v == target`. |
| 3 | **Counter** | Starts at 0. |

## Hint

Range `xs`; `if v == target { count++ }`.

## Validate

```bash
make verify
```
