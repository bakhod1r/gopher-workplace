# Clearing pointer slices without leaking

## The idea

Reducing length leaves element pointers in the backing array; for pointer/interface elements you must nil them (or use `clear`) so the referents can be collected.

## Why it matters

The nil-out-then-truncate pattern prevents retention leaks in reusable buffers.

## Watch out

- `s[:0]` keeps pointers in the array's spare capacity.
- `clear(s)` (Go 1.21+) zeroes elements; then re-slice.
