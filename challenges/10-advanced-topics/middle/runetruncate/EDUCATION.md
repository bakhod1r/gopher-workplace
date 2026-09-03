# Cut A String Without Splitting A Character

## Intuition

UTF-8 is self-synchronising: from any byte you can tell whether it begins a character. So finding a safe cut point is a short walk backwards, never a re-scan from the start.

## Approach

1. Handle `n <= 0` and `n >= len(s)`.
2. While `s[n]` is not a rune start, decrement `n`.
3. Return `s[:n]`.

## Solution

```go
import "unicode/utf8"

// Truncate returns the longest prefix of s that is at most n bytes and
// does not end in the middle of a UTF-8 character.
//
// The result is a substring, so nothing is copied.
//
// Examples:
//
// 	Truncate("héllo", 3) => "hé"
func Truncate(s string, n int) string {
	if n <= 0 {
		return ""
	}
	if n >= len(s) {
		return s
	}
	for n > 0 && !utf8.RuneStart(s[n]) {
		n--
	}
	return s[:n]
}
```

## Walkthrough

In "héllo", é occupies bytes 1 and 2. Cutting at n = 2 lands on the continuation byte, so n drops to 1 and the result is "h".

## Pitfalls

- Converting to `[]rune` to count — correct, and it allocates the whole string again.
- Checking `s[n-1]` instead of `s[n]`; the question is whether the byte you are cutting *before* starts a character.
