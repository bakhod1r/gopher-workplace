# The Conversion You Do Not Need

## Intuition

A string is already a byte sequence in memory. Converting it to `[]byte` only exists so you can *write* to it — reading needs no copy at all.

## Approach

1. Loop with an index and compare `s[i]`.
2. Guard the empty string in the prefix check.

## Solution

```go
func CountByte(s string, b byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			count++
		}
	}
	return count
}

func HasPrefixByte(s string, b byte) bool {
	return len(s) > 0 && s[0] == b
}
```

## Walkthrough

`for _, r := range s` would decode UTF-8 and give runes, so the two-byte `é` would never match a single byte — hence the index form.

## Pitfalls

- `for _, c := range []byte(s)`, which allocates a copy the compiler cannot always elide.
- Ranging with the two-variable form and comparing a `rune` to a `byte`.
- Indexing `s[0]` without the length guard.
