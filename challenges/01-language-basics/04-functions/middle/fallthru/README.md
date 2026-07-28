# Switch Fallthrough

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

`fallthrough` forces control into the next case even though its condition
isn't tested — the opposite of Go's default break-after-case.

## Task

Implement `Rank` in [fallthru.go](fallthru.go). Build the label from the highest tier down using fallthrough.

Do **not** change the function signature or the tests.

## Examples

```go
Rank(3) // => "bronze"
Rank(6) // => "silver/bronze"
Rank(9) // => "gold/silver/bronze"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **fallthrough keyword** | Continues into the next case unconditionally. |
| 2 | **Ordered tiers** | Start at the highest matched tier. |
| 3 | **Accumulate** | Each fallen-through case appends its label. |

## Hint

Use a tagless switch on tiers (`case s>=9:` add "gold", `fallthrough`; `case s>=6:` ...). Note fallthrough ignores the next case's condition, so structure tiers carefully — or build with explicit if-appends. See EDUCATION.

## Validate

```bash
make verify
```
