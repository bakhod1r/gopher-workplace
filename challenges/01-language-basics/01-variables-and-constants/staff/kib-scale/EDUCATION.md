# Kibibytes vs kilobytes

## The idea

Binary and decimal prefixes are different constants:

```go
const KiB = 1024 // 2^10, kibibyte  (binary)
const kB  = 1000 // 10^3, kilobyte  (decimal)
```

Memory, buffers, and page sizes are binary (1024). Disk marketing and SI units
are decimal (1000). Using 1000 where 1024 is meant under-reports by ~2.4% and
the gap compounds per level (MiB, GiB…).

## Why it matters

A wrong factor is a *unit* bug: every derived size is silently off, and the error
grows with scale. It corrupts capacity planning, quota checks, and progress bars,
yet each individual number looks reasonable.

## Watch out

- Prefer `1 << 10`, `1 << 20`… for binary units; it states the intent.
- Name the constant (`KiB`) so the factor lives in one place.
- Be explicit in APIs about which prefix you mean; the ambiguity of "KB" is the
  root cause.

## Try it yourself

```go
const (
	KiB = 1 << 10 // 1024
	MiB = 1 << 20 // 1048576
)
2 * KiB // 2048
```
