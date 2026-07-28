# Accidental Fallthrough

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

In Go a case does NOT fall through by default; `fallthrough` forces entry into
the next case unconditionally (ignoring its condition). Here it makes level 1
leak into the level-2 body, so `Label(1)` returns "mid". Remove the stray keyword.

## Task

Fix [tierbug.go](tierbug.go) so case 1 yields "low".

Do **not** change the function signature or the tests.

## Examples

```go
Label(1) // => "low"
Label(2) // => "mid"
Label(9) // => "?"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Go switch semantics** | No implicit fallthrough between cases. |
| 2 | **fallthrough keyword** | Enters the next case without testing it. |
| 3 | **Case isolation** | Each case should set its own label and stop. |

## Hint

Remove the `fallthrough`; `label = "low"` is enough.

## Validate

```bash
make verify
```
