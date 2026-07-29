# Editing strings

## Intuition

A Go string is immutable — you cannot assign to `s[i]`. To transform it, copy to a
`[]byte`, edit that, and convert back:

```go
b := []byte(s)
// mutate b...
return string(b)
```

Returning the original `s` discards all the work.

## Approach

1. Bug: the loop upper-cases the []byte copy b, but the function returns the original immutable string s, so the change is discarded.
2. Fix: `return string(b)` returns the mutated bytes as a new string.

## Solution

```go
func Upper(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = c - 32
		}
	}
	return string(b)
}
```

## Walkthrough

s="abc": b becomes ['A','B','C'] but s is unchanged. Returning s gives "abc"; returning string(b) gives "ABC".

## Pitfalls

- `s[i] = x` is a compile error; strings are read-only.
- Each conversion (`[]byte(s)`, `string(b)`) copies.
- For rune-level edits, use `[]rune`; for ASCII, `[]byte` suffices.
