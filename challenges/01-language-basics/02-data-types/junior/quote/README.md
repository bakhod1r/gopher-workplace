# Quote a String

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

To embed a double quote inside a double-quoted string you escape it with `\"`.

## Task

Implement `Wrap(s)` returning `s` surrounded by literal `"` characters.

## Examples

```go
Wrap("hello") // => "\"hello\""
Wrap("")      // => "\"\""
Wrap("a b")   // => "\"a b\""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Escaped quote** | `\"` puts a quote char inside a `"..."` literal. |
| 2 | **Concatenation** | Build the result with `+`. |
| 3 | **Empty string** | Wrapping "" yields just two quote chars. |

## Hint

`return "\"" + s + "\""`.

## Validate

```bash
make verify
```
