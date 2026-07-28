# Reverse by Runes

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Reversing bytes corrupts multi-byte UTF-8 characters. Convert to `[]rune` first,
reverse, convert back.

## Task

Implement `Reverse(s)` reversing runes, keeping characters intact.

## Examples

```go
Reverse("hello") // => "olleh"
Reverse("café")  // => "éfac"
Reverse("日本語") // => "語本日"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **[]rune(s)** | Decodes UTF-8 into code points. |
| 2 | **Two-pointer swap** | Swap ends moving inward. |
| 3 | **string([]rune)** | Re-encodes to UTF-8. |

## Hint

`r := []rune(s)`; swap `r[i], r[j]` inward; `return string(r)`.

## Validate

```bash
make verify
```
