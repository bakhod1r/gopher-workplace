# CRLF-Safe Line Split

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A log parser splits on `\n` but forgets that Windows files use `\r\n`, so every
line keeps a trailing `\r`. Comparisons and trims downstream then mysteriously
fail.

## Task

Fix the loop body between the markers in [splitlines.go](splitlines.go) to strip
a trailing `\r`.

## Examples

```go
Lines("a\r\nb") // => ["a", "b"]
Lines("a\nb")   // => ["a", "b"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Line endings** | LF vs CRLF across platforms. |
| 2 | **TrimSuffix** | `strings.TrimSuffix(p, "\r")`. |
| 3 | **Idempotence** | LF input already has no `\r` to strip. |

## Hint

`parts[i] = strings.TrimSuffix(p, "\r")`.

## Validate

```bash
make verify
```
