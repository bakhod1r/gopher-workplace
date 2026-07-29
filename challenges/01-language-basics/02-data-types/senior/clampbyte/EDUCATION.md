# Saturation vs wraparound

## Intuition

Converting an out-of-range `int` to `byte` keeps only the low 8 bits, so `300`
becomes `44` and `-20` becomes `236`. Image and audio math need **saturation** —
clamp to the range first, then convert:

```go
if x < 0 { return 0 }
if x > 255 { return 255 }
return byte(x)
```

## Approach

1. Bug: `return byte(x)` truncates to 8 bits, so 300 wraps to 44 and -20 wraps to 236.
2. Fix: clamp before converting — if x<0 return 0, if x>255 return 255, else byte(x).
3. This saturates out-of-range values instead of wrapping.

## Solution

```go
func Clamp(x int) byte {
	if x < 0 {
		return 0
	}
	if x > 255 {
		return 255
	}
	return byte(x)
}
```

## Walkthrough

Clamp(300): 300>255 -> return 255. Clamp(-20): -20<0 -> return 0. Clamp(128): in range -> byte(128).

## Pitfalls

- Narrowing conversions wrap silently; they never clamp for you.
- Clamp in the wide type (`int`) before the conversion.
- Hardware SIMD has saturating adds for exactly this reason.
