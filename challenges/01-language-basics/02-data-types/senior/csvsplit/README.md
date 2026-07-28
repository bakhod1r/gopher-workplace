# Quoted CSV Split

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

An importer splits CSV lines but breaks `a,"b,c",d` into four fields — it splits
on the comma *inside* the quotes. A comma is only a separator when not inside a
quoted field.

## Task

Fix the `case c == ','` guard between the markers in [csvsplit.go](csvsplit.go)
so commas inside quotes are literal.

## Examples

```go
Split(`a,"b,c",d`) // => ["a", "b,c", "d"]
Split(`"x""y",z`)  // => [`x"y`, "z"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Stateful scan** | Track `inQuotes` across characters. |
| 2 | **Context-sensitive delimiter** | Comma splits only outside quotes. |
| 3 | **Quote escaping** | `""` inside quotes is a literal `"`. |

## Hint

`case c == ',' && !inQuotes:`.

## Validate

```bash
make verify
```
