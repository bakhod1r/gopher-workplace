# Take While

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

TakeWhile consumes a prefix, breaking out of the loop as soon as the predicate
fails — order-sensitive, unlike Filter.

## Task

Implement `TakeWhile` in [takewhile.go](takewhile.go).

Do **not** change the function signature or the tests.

## Examples

```go
TakeWhile([]int{1,2,3,-1,4}, pos) // => [1 2 3]
TakeWhile([]int{-1,2}, pos)       // => []
TakeWhile(nil, pos)               // => []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Break on first failure** | Stop, don't skip. |
| 2 | **Prefix semantics** | Only the leading run is taken. |
| 3 | **Predicate arg** | `pred func(int) bool`. |

## Hint

Range `xs`; `if !pred(v) { break }; out = append(out, v)`.

## Validate

```bash
make verify
```
