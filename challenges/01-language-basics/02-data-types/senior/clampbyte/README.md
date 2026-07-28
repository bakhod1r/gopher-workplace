# Saturating Pixel Clamp

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

Brightening an image adds to pixel values; a result of `300` must saturate to
`255`, not wrap to `44`. `byte(x)` wraps — it keeps only the low 8 bits.

## Task

Fix the conversion between the markers in [clampbyte.go](clampbyte.go) to
saturate to [0,255].

## Examples

```go
Clamp(300) // => 255
Clamp(-20) // => 0
Clamp(128) // => 128
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Conversion wraps** | `byte(300) == 44` (mod 256). |
| 2 | **Saturation** | Clamp before converting. |
| 3 | **Range guard** | Two comparisons pin 0 and 255. |

## Hint

`if x < 0 {return 0}; if x > 255 {return 255}; return byte(x)`.

## Validate

```bash
make verify
```
