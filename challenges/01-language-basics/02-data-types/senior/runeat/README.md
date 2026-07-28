# Rune Index Bounds

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A text editor addresses characters by index. The bounds check uses `len(s)` (the
**byte** length), so for multi-byte text it lets an out-of-range rune index
through — `"日本"` has 2 runes but `len` is 6.

## Task

Fix the bounds check between the markers in [runeat.go](runeat.go) to use the
rune count.

## Examples

```go
At("héllo", 2) // => 'l', true
At("日本", 2)   // => 0, false  (only 2 runes)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Byte vs rune length** | `len(s)` counts bytes, `len([]rune(s))` runes. |
| 2 | **Bounds check** | Compare against the rune count. |
| 3 | **Consistency** | Index and bound must use the same unit. |

## Hint

`n >= len(rs)`.

## Validate

```bash
make verify
```
