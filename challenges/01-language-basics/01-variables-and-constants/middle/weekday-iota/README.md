# Weekday Enum

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Days of the week are a natural enum. `iota` numbers them 0..6 with no manual
values.

## Task

In [weekday.go](weekday.go):

1. Define `Sunday..Saturday` in one const block (`Sunday=0`, `Saturday=6`).
2. Implement `IsWeekend` returning true for Saturday and Sunday.

## Examples

```go
Monday        // => 1
Saturday      // => 6
IsWeekend(Sunday)   // => true
IsWeekend(Wednesday)// => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **iota counter** | Resets to 0 per const block, +1 each line. |
| 2 | **Named typed enum** | `type Day int` gives the constants a distinct type. |
| 3 | **Range over ints** | `for d := Sunday; d <= Saturday; d++` walks the enum. |

## Hint

After `Sunday Day = iota`, the remaining six names on their own lines fill in
1..6 automatically.

## Validate

```bash
make verify
```
