# Unsafe Slice And String

## Intuition

`string(b)` copies because the runtime treats strings as immutable — maps, interning and constant folding all depend on it. `unsafe.String` skips the copy and hands you the obligation instead: never write to the source again.

## Approach

1. Guard the empty case: `&b[0]` panics on a zero-length slice.
2. `unsafe.String(&b[0], len(b))` builds a string header over the existing bytes.
3. `unsafe.Slice(unsafe.StringData(s), len(s))` does the reverse.
4. `SafeString` is the ordinary `string(b)` conversion, for comparison.

## Solution

```go
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}

func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func SafeString(b []byte) string { return string(b) }
```

## Walkthrough

`TestBytesToStringAliases` mutates the backing array and expects the string to change. That is the documented behaviour of this conversion — and precisely why it is only safe when the source is dead after the call.

## Pitfalls

- Writing through `StringToBytes` output: string data may live in read-only memory and the program crashes.
- Using `unsafe.String` on a buffer that is reused (a pooled or ring buffer), which silently corrupts every string you handed out.
- Forgetting the empty guard — `&b[0]` on an empty slice panics with an index out of range.
