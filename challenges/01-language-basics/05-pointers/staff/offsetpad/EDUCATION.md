# Padding and field offsets

## The idea

Fields are placed at aligned offsets, so a field's position is not the sum of prior field sizes; `Offsetof` reports the true, padded offset.

## Why it matters

Manual struct parsing must account for padding via Offsetof, not naive size sums.

## Watch out

- `N` follows 1 bool + 7 padding bytes -> offset 8.
- `Offsetof(r.N)` gives the correct position.
