# LEB128 Varint Shift

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A protobuf reader decodes varints. Each byte holds **7** value bits, but the
decoder advances the shift by 8, so multi-byte values are wildly wrong (300
decodes as 556).

## Task

Fix the shift step between the markers in [varint.go](varint.go).

## Examples

```go
Decode([]byte{0xAC, 0x02}) // => 300, 2
Decode([]byte{0x80, 0x01}) // => 128, 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Base-128 groups** | 7 payload bits per byte. |
| 2 | **Continuation bit** | High bit set = more bytes follow. |
| 3 | **Little-endian groups** | Later bytes are more significant. |

## Hint

`shift += 7`.

## Validate

```bash
make verify
```
