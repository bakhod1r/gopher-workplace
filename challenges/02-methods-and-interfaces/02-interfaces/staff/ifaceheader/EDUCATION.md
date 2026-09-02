# Interface Header

## Intuition

`err != nil` compares the whole two-word interface header, not the pointer inside it. A typed nil has a real type word, so the comparison is true even though the payload is nil — the bug that has bitten every Go codebase at least once.

## Approach

1. Take the address of the `any` parameter and reinterpret it as `[2]uintptr`.
2. Word 0 is the type pointer, word 1 is the data pointer.
3. `IsTypedNil` is `typ != 0 && data == 0`.
4. `Classify` branches on the two words.

## Solution

```go
func Words(v any) (uintptr, uintptr) {
	w := *(*[2]uintptr)(unsafe.Pointer(&v))
	return w[0], w[1]
}

func IsTypedNil(v any) bool {
	typ, data := Words(v)
	return typ != 0 && data == 0
}

func Classify(v any) string {
	typ, data := Words(v)
	switch {
	case typ == 0:
		return "nil"
	case data == 0:
		return "typed-nil"
	default:
		return "value"
	}
}
```

## Walkthrough

A nil map classifies as `typed-nil` because a map value is a single pointer word. A nil *slice* does not: a slice header is three words, so boxing copies it to the heap and the data word points at that copy — it classifies as `value`.

## Pitfalls

- Relying on this layout in production code: it is unspecified and can change between releases. Use `reflect.ValueOf(v).IsNil()` in real code, guarded by a `Kind` check.
- Passing an already-boxed interface without `any(...)`, which re-boxes and changes what you are inspecting.
- Interpreting a non-zero data word as "the value is non-zero" — it is a pointer, not the value.
