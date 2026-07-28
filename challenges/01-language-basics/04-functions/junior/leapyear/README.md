# Leap Year

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

The leap-year rule is a layered condition: the 400 exception overrides the 100
exception, which overrides the base 4 rule.

## Task

Implement `IsLeap` in [leapyear.go](leapyear.go).

Do **not** change the function signature or the tests.

## Examples

```go
IsLeap(2000) // => true
IsLeap(1900) // => false
IsLeap(2024) // => true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Boolean composition** | Combine `%4`, `%100`, `%400`. |
| 2 | **Rule precedence** | 400 beats 100 beats 4. |
| 3 | **Return a bool** | Express the whole rule in one expression or a few ifs. |

## Hint

`return y%400 == 0 || (y%4 == 0 && y%100 != 0)`.

## Validate

```bash
make verify
```
