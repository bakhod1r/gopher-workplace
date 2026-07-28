# JSON String Escaping

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A hand-rolled JSON writer escapes quotes and control chars but forgets the
**backslash** — so `a\b` is emitted as `a\b`, producing invalid JSON that a
strict parser rejects.

## Task

Add the missing `'\\'` case between the markers in
[jsonescape.go](jsonescape.go).

## Examples

```go
Escape(`a"b`) // => `a\"b`
Escape(`a\b`) // => `a\\b`
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Escape set** | `\`, `"`, `\n`, `\t`, `\r` at minimum. |
| 2 | **Backslash first** | It is itself an escape character. |
| 3 | **Raw literals** | `` `\\` `` is a two-char string: backslash, backslash. |

## Hint

Add `case '\\': b.WriteString(`\\\\`)` (a backslash escaped as two).

## Validate

```bash
make verify
```
