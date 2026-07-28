# Filter by Predicate

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A predicate function passed in lets one filter serve every condition.

## Task

Implement `Filter` in [filterints.go](filterints.go).

Do **not** change the function signature or the tests.

## Examples

```go
Filter([]int{1,2,3,4}, even) // => [2 4]
Filter(nil, p)               // => []
Filter([]int{1,3}, even)     // => []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Predicate parameter** | `keep func(int) bool` decides membership. |
| 2 | **Conditional append** | Only kept elements go to the result. |
| 3 | **Fresh slice** | Input is untouched. |

## Hint

Range `xs`; `if keep(v) { out = append(out, v) }`.

## Validate

```bash
make verify
```
