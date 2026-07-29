# Modular letter shifting

## Intuition

Map a letter to 0..25 by subtracting its case base, add the shift, take it mod 26
(normalized non-negative), and add the base back:

```go
base + rune(((int(r-base)+n)%26+26)%26)
```

Non-letters are left untouched.

## Approach

1. Normalize the shift: shift = ((n%26)+26)%26 so negatives/large values work. 2. For each byte, if it's a-z map within lowercase, if A-Z within uppercase, using (c-base+shift)%26. 3. Non-letters pass through.

## Solution

```go
func Shift(s string, n int) string {
	shift := ((n % 26) + 26) % 26
	out := []byte(s)
	for i := 0; i < len(out); i++ {
		c := out[i]
		switch {
		case c >= 'a' && c <= 'z':
			out[i] = 'a' + byte((int(c-'a')+shift)%26)
		case c >= 'A' && c <= 'Z':
			out[i] = 'A' + byte((int(c-'A')+shift)%26)
		}
	}
	return string(out)
}
```

## Walkthrough

Shift("xyz",3): shift=3. 'x'->(23+3)%26=0->'a', 'y'->1->'b', 'z'->2->'c'. "abc".

## Pitfalls

- Handle upper and lower case with their own base.
- Normalize the modulo (`(x%26+26)%26`) so negative shifts wrap correctly.
- ROT13 is `Shift(s,13)` and is its own inverse.
