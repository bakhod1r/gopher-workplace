# Do These Two Strings Share Their Bytes

## Intuition

Go gives no way to ask whether two strings are the same object, because normally it should not matter. It matters exactly when you are verifying that a cache or interner is doing its job.

## Approach

1. Return false on a length mismatch.
2. Return false for empty strings.
3. Compare `unsafe.StringData(a)` with `unsafe.StringData(b)`.

## Solution

```go
import "unsafe"

// SameBytes reports whether a and b are the same length and start at the
// same address — that is, whether they share their storage.
//
// Two equal strings may or may not share; this asks about identity, not
// equality.
//
// Examples:
//
// 	s := "abc"; SameBytes(s, s) => true
func SameBytes(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return false
	}
	return unsafe.StringData(a) == unsafe.StringData(b)
}
```

## Walkthrough

`s` and `s[:32]` on a 32-byte string share both the length and the start, so the answer is true. A round trip through `[]byte` allocates a new array and the answer is false.

## Pitfalls

- Skipping the length check, which would call two different-length strings identical.
- Reading anything into a false result — sharing is an implementation detail, not a guarantee.
