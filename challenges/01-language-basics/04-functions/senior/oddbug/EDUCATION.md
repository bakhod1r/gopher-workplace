# Signed modulo and parity

## The idea

Because `%` follows the dividend's sign, `n%2` is -1 for negative odds; only `!= 0` reliably detects oddness.

## Why it matters

The `%2 == 1` odd-test is a real bug that only shows up on negative inputs.

## Watch out

- `-3 % 2` is `-1`, so `== 1` fails.
- Prefer `v%2 != 0` for parity across all integers.
