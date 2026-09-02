# Slice Retention Leak

## Intuition

The GC frees whole objects, not parts of them. A slice header pointing anywhere into a 1MB array keeps the entire array reachable — so an 8-byte "result" can pin a megabyte indefinitely.

## Approach

1. `Prefix` clamps `n` and returns `b[:n]`, deliberately sharing the array.
2. `PrefixCopy` allocates exactly `n` bytes and copies into it.
3. `RetainedBytes` returns `cap(b)`, which is what the reference actually holds alive.
4. Both clamp negative and oversized `n`.

## Solution

```go
func Prefix(b []byte, n int) []byte {
	if n > len(b) {
		n = len(b)
	}
	if n < 0 {
		n = 0
	}
	return b[:n]
}

func PrefixCopy(b []byte, n int) []byte {
	if n > len(b) {
		n = len(b)
	}
	if n < 0 {
		n = 0
	}
	out := make([]byte, n)
	copy(out, b)
	return out
}

func RetainedBytes(b []byte) int { return cap(b) }
```

## Walkthrough

`TestManyExtractsRetention` keeps 8 bytes from a hundred 64KB messages. The aliasing version pins 6.4MB; the copying version pins 800 bytes. Same visible data, four orders of magnitude apart.

## Pitfalls

- `b[:n:n]` limits the capacity but still points into the same array — the memory is not released.
- Assuming a short slice is cheap: `len` is what you see, `cap` is what you pay for.
- Copying everything defensively: when the source is short-lived and small, the alias is the right call.
