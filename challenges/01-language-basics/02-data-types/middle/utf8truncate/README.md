# Byte-Safe Truncate

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A database column is `VARCHAR(N)` counted in **bytes**. Truncating UTF-8 text to
fit must not cut a multi-byte character in half, or you store invalid UTF-8.

## Task

Implement `Truncate(s, max)` = longest prefix ≤ `max` bytes, never splitting a
rune.

## Examples

```go
Truncate("héllo", 2) // => "h"   (é is 2 bytes)
Truncate("héllo", 3) // => "hé"
Truncate("日本", 4)   // => "日"  (3 bytes each)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Byte vs rune width** | A rune spans 1–4 bytes. |
| 2 | **range byte index** | The index is where each rune starts. |
| 3 | **Prefix cut** | Cut at the last rune boundary ≤ max. |

## Hint

`for i := range s { if i+len(rune-at-i-bytes) > max ... }` — simpler: track the
last boundary `i` where `i <= max`; when a rune would cross `max`, stop and
return `s[:i]`.

## Validate

```bash
make verify
```
