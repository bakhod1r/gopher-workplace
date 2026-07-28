# Returning reused buffers

## The idea

Reusing a buffer with `buf[:0]` avoids allocations, but returning it aliases shared memory that the next call overwrites; hand out a copy when the caller keeps the result.

## Why it matters

Buffer-reuse APIs (scanners, encoders) that leak the internal slice corrupt earlier results — a real, subtle bug.

## Watch out

- `buf[:0]` reuses the same array; returning it shares mutable memory.
- Copy with `append([]T(nil), buf...)` when the caller retains the slice.
