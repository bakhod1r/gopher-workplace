# Closures Share One Cell

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

All closures created in the loop capture the SAME outer `c`, so they share and
fight over one counter. Each closure needs its own captured variable, declared
per iteration.

## Task

Fix [sharedcell.go](sharedcell.go) so each counter is independent.

Do **not** change the function signature or the tests.

## Examples

```go
cs := Counters(2); cs[0](); cs[0]() // => 1, 2
cs[1]() // => 1 (independent)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Captured variable identity** | One `c` shared by all closures. |
| 2 | **Per-instance state** | Declare the counter inside the loop. |
| 3 | **Closure environment** | Each closure closes over its own `c`. |

## Hint

Move `c := 0` INSIDE the loop so each closure captures a fresh counter.

## Validate

```bash
make verify
```
