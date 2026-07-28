# Map Over Slice

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Passing a function as an argument makes a transform generic: the caller decides
the per-element operation.

## Task

Implement `MapInts` in [mapints.go](mapints.go) returning a new slice.

Do **not** change the function signature or the tests.

## Examples

```go
MapInts([]int{1,2,3}, square) // => [1 4 9]
MapInts(nil, f)               // => []
MapInts([]int{2}, inc)        // => [3]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Function as a parameter** | `f func(int) int` is called per element. |
| 2 | **Build new slice** | append the transformed value. |
| 3 | **First-class functions** | Functions are values you can pass around. |

## Hint

Range `xs`, appending `f(v)` to a fresh `out`.

## Validate

```bash
make verify
```
