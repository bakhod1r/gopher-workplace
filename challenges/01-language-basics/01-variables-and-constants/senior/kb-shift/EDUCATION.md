# Shift stride in iota units

## The idea

Binary size units scale by 1024 per step, which is a **10-bit** shift, not one
bit. The `iota` slot picks the exponent, so the multiplier is `10 * iota`:

```go
const (
	_  ByteSize = iota          // consumes iota==0
	KB ByteSize = 1 << (10 * iota) // iota==1 -> 1<<10 == 1024
	MB                              // iota==2 -> 1<<20
	GB                              // iota==3 -> 1<<30
)
```

`1 << iota` (stride 1) would give 2, 4, 8 — powers of two, not powers of 1024.

## Why it matters

Off-by-a-stride bugs produce values that are still powers of two, so they look
"binary-ish" and pass a glance. But KB=2 instead of 1024 corrupts every capacity
calculation downstream.

## Watch out

- The blank `_ = iota` line matters: it puts KB on `iota==1`, so `10*iota == 10`.
- Use `uint64`; `1 << 40` (TB) overflows 32-bit ints.
- Verify by pinning KB to exactly 1024 in a test.

## Try it yourself

```go
const (
	_  = iota
	A  = 1 << (10 * iota) // 1024
	B                     // 1048576
)
```
