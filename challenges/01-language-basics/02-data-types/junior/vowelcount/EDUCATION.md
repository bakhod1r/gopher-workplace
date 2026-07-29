# Runes and character tests

## Intuition

Ranging over a string yields runes, which you compare against **rune literals**
like `'a'`. A `switch` with a case list is the clean way to test membership:

```go
for _, r := range s {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		n++
	}
}
```

## Approach

1. Range over the string rune by rune.
2. For each rune, check membership in the ASCII vowel set (both cases).
3. Increment a counter and return it.

## Solution

```go
func Vowels(s string) int {
	count := 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}
```

## Walkthrough

Vowels("hello"): runes h,e,l,l,o; e and o match -> 2.

## Pitfalls

- `'a'` is a rune (int32) constant; `"a"` is a string — different types.
- To fold case, add both `'a'` and `'A'`, or convert with
  `unicode.ToLower`.
- Accented vowels are distinct runes; decide deliberately whether they count.
