# String ↔ []byte conversion

## Intuition

Strings and byte slices convert with an ordinary conversion:

```go
[]byte("Go")          // []byte{71, 111}
string([]byte{71, 111}) // "Go"
```

A string is an immutable sequence of bytes; a `[]byte` is a mutable one. Each
conversion **copies**, so mutating the slice never changes the source string.

## Approach

1. FromBytes: return string(b) - a plain conversion that copies.
2. ToBytes: return []byte(s) - a plain conversion that copies.

## Solution

```go
func FromBytes(b []byte) string {
	return string(b)
}

func ToBytes(s string) []byte {
	return []byte(s)
}
```

## Walkthrough

ToBytes("Go"): 'G'=71, 'o'=111, so []byte("Go") = [71 111].

## Pitfalls

- Each conversion allocates and copies (O(n)); avoid it in hot loops.
- Converting a `[]byte` with invalid UTF-8 to a string keeps the bytes as-is;
  ranging over it later yields `utf8.RuneError` for bad sequences.
- `string(65)` is `"A"` (rune conversion), *not* `"65"` — a common surprise;
  `string([]byte{65})` is also `"A"`.
