# Hand Out Bytes That Cannot Be Written

## Intuition

`unsafe.Slice` over `StringData` produces a writable-looking view of memory the whole program treats as constant. There is no way to keep the function's promise without giving the caller bytes of their own.

## Approach

1. Return nil for the empty string.
2. Allocate `len(s)` bytes, copy the string into them, and return the copy.

## Solution

```go
// Snapshot returns a byte slice holding s's bytes that the caller may
// modify freely.
//
// A view over the string's own bytes is not that: strings may live in
// read-only memory, and every other holder of s would see the writes.
//
// Examples:
//
// 	b := Snapshot("hi"); b[0] = 'H' => s is unaffected
func Snapshot(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	out := make([]byte, len(s))
	copy(out, s)
	return out
}
```

## Walkthrough

`Snapshot("abcd")` allocates four bytes and copies. Writing `b[0]` now touches only that array; before the fix it wrote into the string literal's storage.

## Pitfalls

- Assuming a runtime-built string is safe to write — it is still shared with everyone holding that string value.
- `[]byte(s)` is the ordinary spelling of the fix; the point is that the copy is required.
