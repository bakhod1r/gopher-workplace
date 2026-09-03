# Let The Caller Own The Memory

## Intuition

The function does not need to allocate — it needs somewhere to write. Taking that somewhere as a parameter turns a per-call allocation into a one-time one.

## Approach

1. `AppendJoin` appends each part, writing the separator before every part except the first.
2. `Sized` sums the part lengths plus `len(sep) * (len(parts)-1)`.

## Solution

```go
func AppendJoin(scratch []byte, parts []string, sep string) []byte {
	for i, p := range parts {
		if i > 0 {
			scratch = append(scratch, sep...)
		}
		scratch = append(scratch, p...)
	}
	return scratch
}

func Sized(parts []string, sep string) int {
	if len(parts) == 0 {
		return 0
	}
	n := len(sep) * (len(parts) - 1)
	for _, p := range parts {
		n += len(p)
	}
	return n
}
```

## Walkthrough

`scratch[:0]` at the call site is what makes reuse work: the length resets while the array stays, so every subsequent join writes into the same memory.

## Pitfalls

- Allocating inside the function "just to be safe", which throws away the whole benefit.
- Computing `len(sep) * len(parts)` in `Sized`, over-allocating by one separator.
- Returning a slice into the caller's scratch and then holding onto it — the next call overwrites it.
