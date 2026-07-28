# Saturation vs wraparound

## The idea

Converting an out-of-range `int` to `byte` keeps only the low 8 bits, so `300`
becomes `44` and `-20` becomes `236`. Image and audio math need **saturation** —
clamp to the range first, then convert:

```go
if x < 0 { return 0 }
if x > 255 { return 255 }
return byte(x)
```

## Why it matters

Signal processing (pixels, samples) relies on saturating arithmetic: overshoots
should peak, not wrap to a wildly different value. A raw `byte(x)` conversion
produces visible artifacts (bright spots turning dark).

## Watch out

- Narrowing conversions wrap silently; they never clamp for you.
- Clamp in the wide type (`int`) before the conversion.
- Hardware SIMD has saturating adds for exactly this reason.
