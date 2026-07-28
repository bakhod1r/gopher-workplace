# Runes and character tests

## The idea

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

## Why it matters

Working at the rune level keeps the count correct for UTF-8 text: `é` is one
rune, not two bytes, and it correctly does *not* match an ASCII vowel here.

## Watch out

- `'a'` is a rune (int32) constant; `"a"` is a string — different types.
- To fold case, add both `'a'` and `'A'`, or convert with
  `unicode.ToLower`.
- Accented vowels are distinct runes; decide deliberately whether they count.

## Try it yourself

```go
'a' == 97          // true (rune is int32)
switch 'E' { case 'e', 'E': } // matches
```
