# Bytes Are Not Characters

## Intuition

Go strings are UTF-8 bytes. `len` answers a storage question and `range` answers a text question — they agree only for ASCII, which is why the bug survives every English test case.

## Approach

1. `bytes = len(s)`.
2. Range over `s`, incrementing a counter.

## Solution

```go
// Counts returns how many bytes and how many characters s holds.
//
// len gives bytes; ranging a string yields characters.
//
// Examples:
//
// 	Counts("héllo") => 6, 5
func Counts(s string) (bytes, runes int) {
	bytes = len(s)
	for range s {
		runes++
	}
	return bytes, runes
}
```

## Walkthrough

"héllo" is six bytes: h, two for é, then l, l, o. Ranging it yields five iterations, one per character.

## Pitfalls

- `utf8.RuneCountInString` is the real answer and is also allocation-free — the point here is why `len` differs.
- `for i := range s` gives byte offsets, not consecutive indices.
