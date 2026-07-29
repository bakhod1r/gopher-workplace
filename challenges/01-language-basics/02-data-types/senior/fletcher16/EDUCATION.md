# Getting the modulus right

## Intuition

Fletcher-16 keeps two running sums: `sum1` accumulates the bytes, `sum2`
accumulates `sum1`. Both are reduced **mod 255** (not 256) each step, and the
result is `(sum2 << 8) | sum1`.

## Approach

1. Bug: both running sums used `% 256` instead of the Fletcher-16 modulus `% 255`.
2. Fix: change both moduli to 255.
3. Combine as (sum2<<8)|sum1.

## Solution

```go
func Checksum(data []byte) uint16 {
	var sum1, sum2 uint16
	for _, b := range data {
		sum1 = (sum1 + uint16(b)) % 255
		sum2 = (sum2 + sum1) % 255
	}
	return sum2<<8 | sum1
}
```

## Walkthrough

"abcde": running sums mod 255 accumulate to sum2=0xC8, sum1=0xF0 -> 0xC8F0.

## Pitfalls

- The modulus is 255 (`2^8 - 1`), a common surprise vs the "natural" 256.
- Reducing every step avoids overflow and matches the spec exactly.
- Constants like this must come from the specification, not intuition.
