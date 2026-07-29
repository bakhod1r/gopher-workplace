# Reversing Unicode text

## Intuition

Reversing the **bytes** of "café" splits the two-byte `é` and produces invalid
UTF-8. Convert to a `[]rune` (one element per character), reverse that, and
convert back:

```go
r := []rune(s)
for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 { r[i], r[j] = r[j], r[i] }
return string(r)
```

## Approach

1. Convert to []rune so multi-byte characters are single elements. 2. Two-pointer swap from ends toward middle. 3. Convert back to string.

## Solution

```go
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

## Walkthrough

Reverse("café"): runes [c,a,f,é] -> swap to [é,f,a,c] -> "éfac".

## Pitfalls

- `[]rune(s)` allocates and decodes; fine here, wasteful in tight loops.
- This reverses code points, not grapheme clusters — combining marks or emoji ZWJ
  sequences can still reorder visibly (a deeper Unicode topic).
- `len([]rune(s))` is the rune count, unlike `len(s)`.
