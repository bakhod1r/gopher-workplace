# Decode 2-Byte UTF-8

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A hand-written decoder combines a 2-byte sequence into a rune. It masks the lead
byte with `0x0F` (4 bits), but a 2-byte lead carries **5** payload bits
(`110xxxxx`), so any code point ≥ U+0100 loses its top bit.

## Task

Fix the lead-byte mask between the markers in [rune2decode.go](rune2decode.go).

## Examples

```go
Decode2(0xC3, 0xA9) // => 'é' (U+00E9)
Decode2(0xC3, 0xB1) // => 'ñ' (U+00F1)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Payload bits** | 2-byte lead has 5 payload bits (`0x1F`). |
| 2 | **Continuation** | 6 payload bits (`0x3F`). |
| 3 | **Assembly** | `lead5<<6 | cont6`. |

## Hint

`rune(lead&0x1F)<<6 | rune(cont&0x3F)`.

## Validate

```bash
make verify
```
