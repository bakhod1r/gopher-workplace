# Raw String Path

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

A Windows path is full of backslashes. In an interpreted `"..."` string each `\`
must be doubled; a raw `` `...` `` literal takes the characters verbatim.

## Task

Implement `TempPath()` returning `C:\Users\temp\log.txt` using a **raw** string
literal (backticks).

## Examples

```go
TempPath() // => `C:\Users\temp\log.txt`
len(TempPath()) // => 21
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Raw string literal** | Backtick-delimited; backslashes are literal. |
| 2 | **No escapes** | `\t`, `\n`, `\\` are not interpreted inside backticks. |
| 3 | **Single line vs multi** | Raw strings may span lines and include quotes. |

## Hint

Return `` `C:\Users\temp\log.txt` `` — backticks, no escaping.

## Validate

```bash
make verify
```
