# Getting the modulus right

## The idea

Fletcher-16 keeps two running sums: `sum1` accumulates the bytes, `sum2`
accumulates `sum1`. Both are reduced **mod 255** (not 256) each step, and the
result is `(sum2 << 8) | sum1`.

## Why it matters

A one-off modulus (256 vs 255) produces a plausible 16-bit checksum that simply
never matches the standard. Checksums are all-or-nothing: a subtly wrong constant
means every cross-implementation verification fails.

## Watch out

- The modulus is 255 (`2^8 - 1`), a common surprise vs the "natural" 256.
- Reducing every step avoids overflow and matches the spec exactly.
- Constants like this must come from the specification, not intuition.
