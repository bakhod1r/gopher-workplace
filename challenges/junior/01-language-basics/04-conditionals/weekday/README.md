# Weekday Switch

**Level:** junior
**Topic:** 01-language-basics → 04-conditionals
**Estimated time:** 10 min

## Context

A scheduler prints a day name from its number. A `switch` is cleaner than a
long if-ladder here — but an out-of-range number still needs a sensible answer.

## Task

Implement `Weekday` in [weekday.go](weekday.go) using a `switch`: 1→"Monday" …
7→"Sunday", anything else → "Unknown".

Do **not** change the function signature or the tests.

## Examples

```go
Weekday(1) // => "Monday"
Weekday(7) // => "Sunday"
Weekday(0) // => "Unknown"
Weekday(8) // => "Unknown"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Expression switch** | `switch d { case 1: … }` compares `d` against each case value. |
| 2 | **No auto fallthrough** | Go stops after a matching case — no `break` needed, cases don't leak into each other. |
| 3 | **default** | The `default` case handles every value no `case` matched — use it for "Unknown". |

## Hint

`switch d { case 1: return "Monday"; …; case 7: return "Sunday"; default:
return "Unknown" }`. Each case returns on its own; there is no fallthrough.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
