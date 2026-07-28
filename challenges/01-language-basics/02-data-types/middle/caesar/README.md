# Caesar Cipher

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Each letter shifts by `n` within its case, wrapping z→a. Non-letters stay put.
Rune arithmetic plus a modulo does the wrap.

## Task

Implement `Shift(s, n)` (letters only, case preserved, wrap, `n` any int).

## Examples

```go
Shift("abc", 1)          // => "bcd"
Shift("xyz", 3)          // => "abc"
Shift("Hello, World!",13)// => "Uryyb, Jbeyq!"
Shift("abc", -1)         // => "zab"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Rune ranges** | `'a'..'z'`, `'A'..'Z'` are contiguous. |
| 2 | **Modular wrap** | `((r-base+n)%26+26)%26` keeps 0..25. |
| 3 | **Building strings** | Accumulate in a `[]rune` or strings.Builder. |

## Hint

For a letter, base = `'a'` or `'A'`; new = `base + ((r-base+n)%26+26)%26`.

## Validate

```bash
make verify
```
