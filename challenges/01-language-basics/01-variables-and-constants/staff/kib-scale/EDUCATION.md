# Kibibytes vs kilobytes

## Intuition

Binary and decimal prefixes are different constants:

```go
const KiB = 1024 // 2^10, kibibyte  (binary)
const kB  = 1000 // 10^3, kilobyte  (decimal)
```

Memory, buffers, and page sizes are binary (1024). Disk marketing and SI units
are decimal (1000). Using 1000 where 1024 is meant under-reports by ~2.4% and
the gap compounds per level (MiB, GiB…).

## Approach

1. A kibibyte is 1024 bytes, not 1000.
2. Set `KiB = 1024`.

## Solution

```go
const KiB = 1024

func Bytes(n int) int { return n * KiB }
```

## Walkthrough

The decimal 1000 makes `Bytes(2)` 2000; the binary 1024 gives 2048.

## Pitfalls

- Prefer `1 << 10`, `1 << 20`… for binary units; it states the intent.
- Name the constant (`KiB`) so the factor lives in one place.
- Be explicit in APIs about which prefix you mean; the ambiguity of "KB" is the
  root cause.
