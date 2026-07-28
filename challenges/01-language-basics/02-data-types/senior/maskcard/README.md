# Mask All But Last 4

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A receipt shows a card as `************3456`. The condition is inverted — it
masks the last four and reveals the rest, leaking the sensitive part.

## Task

Fix the condition between the markers in [maskcard.go](maskcard.go) so only the
last four are visible.

## Examples

```go
Mask("1234567890123456") // => "************3456"
Mask("12345")            // => "*2345"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Index window** | Last four are indices `>= len-4`. |
| 2 | **Invert the mask** | Mask the *front*, keep the tail. |
| 3 | **Runes** | Count characters, not bytes. |

## Hint

Mask when `i < len(r)-4`.

## Validate

```bash
make verify
```
