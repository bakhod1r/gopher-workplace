# Method Expression Receiver

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A method EXPRESSION `(*Counter).Add` is a function whose FIRST argument is the
receiver: `f(c, d)` calls `c.Add(d)`. The bug ignores the passed receiver and
mutates an unrelated local.

## Task

Fix [methodexpr.go](methodexpr.go) to return the method expression.

Do **not** change the function signature or the tests.

## Examples

```go
f := AdderExpr(); f(c, 5) // c.N += 5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Method expression** | `(*Counter).Add` takes the receiver as arg 1. |
| 2 | **Explicit receiver** | The caller passes which Counter to mutate. |
| 3 | **Value vs expression** | Method value binds a receiver; expression does not. |

## Hint

Return the method expression: `return (*Counter).Add`.

## Validate

```bash
make verify
```
