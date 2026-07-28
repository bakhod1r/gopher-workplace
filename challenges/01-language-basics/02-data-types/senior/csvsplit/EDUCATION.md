# Context-sensitive delimiters

## The idea

A CSV comma is a field separator only when you are **not** inside a quoted field.
The parser carries an `inQuotes` flag; the split case must respect it:

```go
case c == ',' && !inQuotes:
```

## Why it matters

Real formats (CSV, shell, CSVs of addresses) have delimiters whose meaning
depends on surrounding state. Ignoring the state is the classic CSV bug that
shreds any field containing a comma.

## Watch out

- Toggle `inQuotes` on a lone quote; a doubled `""` inside quotes is a literal.
- The final field is flushed after the loop.
- Trailing/empty fields matter: `""` line yields one empty field.
