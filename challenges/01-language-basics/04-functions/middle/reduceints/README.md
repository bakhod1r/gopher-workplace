# Reduce/Fold

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Reduce threads an accumulator through the slice, combining each element with a
caller-supplied binary function.

## Task

Implement `Reduce` in [reduceints.go](reduceints.go).

Do **not** change the function signature or the tests.

## Examples

```go
Reduce([]int{1,2,3,4}, 0, add) // => 10
Reduce([]int{1,2,3,4}, 1, mul) // => 24
Reduce(nil, 5, add)            // => 5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Accumulator threading** | `acc = f(acc, x)` each step. |
| 2 | **Seed value** | `init` is the starting accumulator. |
| 3 | **Binary combiner** | `f(acc, x)` folds two into one. |

## Hint

Start `acc := init`; range `xs` doing `acc = f(acc, v)`; return `acc`.

## Validate

```bash
make verify
```
