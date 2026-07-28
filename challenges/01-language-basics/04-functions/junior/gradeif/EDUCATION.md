# Modulo with signed integers

## The idea

Go's `%` keeps the sign of the dividend, so `-3 % 2 == -1`; test evenness with `== 0`, never `== 1`.

## Why it matters

A `% 2 == 1` odd-test silently fails for negative inputs — a real bug source.

## Watch out

- `-3 % 2` is `-1`, not `1`; only `== 0` is reliable for parity.
- 0 is even.
