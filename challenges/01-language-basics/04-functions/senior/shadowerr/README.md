# Shadowed Named Return

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

`:=` inside a block creates NEW variables, shadowing the outer named returns.
Here the inner `err` never reaches the outer `err`, and the final `return n, err`
sends the zero-valued outer `err` (nil) even after a failure.

## Task

Fix the shadowing bug between the markers in [shadowerr.go](shadowerr.go) so parse failures propagate.

Do **not** change the function signature or the tests.

## Examples

```go
Parse("42")   // => 42, nil
Parse("nope") // => 0, non-nil error
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variable shadowing** | `:=` in a block hides the outer name. |
| 2 | **Named returns** | `err` is meant to carry the failure out. |
| 3 | **`:=` vs `=`** | Use `=` to assign the existing named return. |

## Hint

The `if v, err := ...` declares a fresh `err`. Assign to the outer one instead: `v, err = strconv.Atoi(s)` after declaring `v`, or restructure so the outer `err` is set.

## Validate

```bash
make verify
```
