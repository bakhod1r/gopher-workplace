# Sign via Switch

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

A tagless `switch` (no expression after `switch`) reads like an if-else ladder,
matching the first true case.

## Task

Implement `Sign` in [signswitch.go](signswitch.go) using a switch.

Do **not** change the function signature or the tests.

## Examples

```go
Sign(-4) // => "negative"
Sign(0)  // => "zero"
Sign(7)  // => "positive"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Tagless switch** | `switch { case cond: }` with boolean cases. |
| 2 | **First match wins** | Cases are tried top to bottom. |
| 3 | **No fallthrough** | Go stops after the matched case. |

## Hint

`switch { case n < 0: return "negative"; case n == 0: return "zero"; default: return "positive" }`.

## Validate

```bash
make verify
```
