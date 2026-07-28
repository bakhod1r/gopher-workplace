# Accumulator Reset in Recursion

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A tail-recursive accumulator must carry the running total forward: `acc + xs[0]`.
The bug passes `xs[0]` alone, discarding everything accumulated so far.

## Task

Fix the recursive step in [accreset.go](accreset.go).

Do **not** change the function signature or the tests.

## Examples

```go
Sum([1 2 3 4]) // => 10
Sum(nil)       // => 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Accumulator threading** | Add to `acc`, don't replace it. |
| 2 | **Tail recursion** | Each call passes the updated accumulator. |
| 3 | **Base case returns acc** | The threaded total exits at the base. |

## Hint

Thread the accumulator: `return sumAcc(xs[1:], acc+xs[0])`.

## Validate

```bash
make verify
```
