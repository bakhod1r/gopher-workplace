# Hex Encode

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Each byte becomes two hex chars: the high nibble `b>>4` and the low nibble
`b&0x0f`, each mapped through `"0123456789abcdef"`.

## Task

Implement `Encode(b)` (lowercase, two chars/byte). No `encoding/hex`.

## Examples

```go
Encode([]byte{0xFF})     // => "ff"
Encode([]byte{0x1a,0x2b})// => "1a2b"
Encode([]byte("Go"))     // => "476f"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nibble split** | `b>>4` and `b&0x0f`. |
| 2 | **Digit table** | Index `"0..f"` by nibble value. |
| 3 | **Builder** | Append two chars per byte. |

## Hint

`const hexd = "0123456789abcdef"`; append `hexd[b>>4]` and `hexd[b&0x0f]`.

## Validate

```bash
make verify
```
