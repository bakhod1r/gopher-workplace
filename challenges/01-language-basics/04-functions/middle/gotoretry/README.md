# Retry with Goto

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`goto` jumps to a label in the same function. It can express a loop, though a
normal `for` is clearer — this exercise makes the control flow explicit.

## Task

Implement `SumUntil` in [gotoretry.go](gotoretry.go) using a `goto` label to repeat the add step.

Do **not** change the function signature or the tests.

## Examples

```go
SumUntil([]int{1,2,3,4}, 5) // => 6, 3
SumUntil([]int{1,1}, 100)   // => 2, 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **goto and labels** | `loop:` ... `goto loop`. |
| 2 | **Loop condition by hand** | Check index bound and limit before jumping. |
| 3 | **Two accumulators** | sum and used advance together. |

## Hint

Label `loop:`; if `used < len(xs) && sum < limit` add `xs[used]`, `used++`, `goto loop`.

## Validate

```bash
make verify
```
