# Slice-to-array conversion

## Intuition

Go 1.20+ allows converting a slice to an array or array pointer: `[4]byte(b)`
copies the first 4 bytes. It **panics** if the slice is shorter than the array,
so guard first:

```go
if len(b) < 4 { return [4]byte{}, false }
return [4]byte(b[:4]), true
```

## Approach

1. Bug: `return [4]byte(b[:4]), true` slices b[:4] unconditionally; on input shorter than 4 the slice expression panics.
2. Fix: guard `if len(b) < 4 { return [4]byte{}, false }` before the conversion.

## Solution

```go
func First4(b []byte) ([4]byte, bool) {
	if len(b) < 4 {
		return [4]byte{}, false
	}
	return [4]byte(b[:4]), true
}
```

## Walkthrough

b=[1 2]: len 2 < 4, so the fix returns the zero array and false. b=[1 2 3 4 5]: len >= 4, convert first 4 -> [1 2 3 4], true.

## Pitfalls

- The slice length must be `>=` the array length, or it panics.
- `(*[4]byte)(b)` gives a pointer (no copy) with the same length rule.
- Arrays are comparable and copyable; slices are neither.
