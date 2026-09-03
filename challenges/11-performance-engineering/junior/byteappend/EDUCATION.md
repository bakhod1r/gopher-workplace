# The Append API Shape

## Intuition

Returning a freshly allocated `[]byte` forces an allocation per call. Appending to the caller's buffer moves that decision — and the cost — out of the hot path.

## Approach

1. Append the key, then `'='`, then the value, then `';'`.
2. Return the final slice.

## Solution

```go
func AppendRecord(dst []byte, key, value string) []byte {
	dst = append(dst, key...)
	dst = append(dst, '=')
	dst = append(dst, value...)
	return append(dst, ';')
}
```

## Walkthrough

`buf[:0]` in the test resets the length while keeping the capacity, so every iteration writes into the same array and the allocation count stays at zero.

## Pitfalls

- `make([]byte, 0, n)` inside the function, which defeats the entire pattern.
- Ignoring `append`'s return value, so a reallocation is silently lost.
- `append(dst, []byte(key)...)`, which allocates a conversion the compiler would otherwise elide.
