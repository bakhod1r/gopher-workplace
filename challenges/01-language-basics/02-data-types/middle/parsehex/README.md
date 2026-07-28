# Parse Hex

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Fold hex digits into a value: `n = n*16 + d`, where `d` comes from `'0'..'9'`,
`'a'..'f'`, or `'A'..'F'`.

## Task

Implement `Parse(s)` (no `0x`), returning `(value, ok)`; false on non-hex/empty.

## Examples

```go
Parse("ff")   // => 255, true
Parse("1A2B") // => 6699, true
Parse("xy")   // => 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Hex digit value** | Map three char ranges to 0..15. |
| 2 | **Horner fold** | `n = n*16 + d`. |
| 3 | **Case folding** | Accept both `a-f` and `A-F`. |

## Hint

Per byte: if `'0'..'9'` → `c-'0'`; `'a'..'f'` → `c-'a'+10`; `'A'..'F'` →
`c-'A'+10`; else fail.

## Validate

```bash
make verify
```
