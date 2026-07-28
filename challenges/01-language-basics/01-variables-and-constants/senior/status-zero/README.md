# Zero-Value Enum Trap

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

If the first enum member is `iota` (0), then a freshly declared `var s Status`
is silently a valid state — the zero value collides with a real one. Shift the
run so 0 stays reserved for "unknown".

## Task

`Pending` starts at plain `iota` (0), so a zero-valued `Status` reads as
`Pending` and the values are 0,1,2. Fix only the marked line so
`Pending=1, Shipped=2, Delivered=3` and the zero value is unknown.

## Examples

```go
var s Status       // 0
IsKnown(s)         // => false
Pending            // => 1
Delivered          // => 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Zero value** | Every `Status` defaults to 0. |
| 2 | **iota offset** | `iota + 1` reserves 0 for "unset". |
| 3 | **Enum hygiene** | Keep the zero value meaning "unknown/invalid". |

## Hint

Change `Pending Status = iota` to `Pending Status = iota + 1`.

## Validate

```bash
make verify
```
