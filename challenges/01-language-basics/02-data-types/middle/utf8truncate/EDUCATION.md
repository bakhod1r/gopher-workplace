# Truncating UTF-8 safely

## Intuition

`for i, r := range s` gives the **byte index** where each rune starts. A rune
occupies `utf8.RuneLen(r)` bytes. Keep runes while the next one still fits under
the byte budget; stop at the boundary:

```go
for i, r := range s {
	if i+utf8.RuneLen(r) > max { return s[:i] }
}
return s
```

## Approach

1. Range over s, which yields the starting byte index i of each rune. 2. Track the last rune boundary that is still <= max. 3. When an index i exceeds max, return s[:last]. 4. If the whole string fits (len<=max), return it all.

## Solution

```go
func Truncate(s string, max int) string {
	last := 0
	for i := range s {
		if i > max {
			return s[:last]
		}
		last = i
	}
	if len(s) <= max {
		return s
	}
	return s[:last]
}
```

## Walkthrough

Truncate("héllo",2): boundaries i=0(h),1(é),3(l)... at i=3 > 2 return s[:last] where last=1 -> "h".

## Pitfalls

- The range index is a byte offset, exactly the safe cut point.
- `s[:i]` slices bytes without copying — cheap.
- `max` larger than `len(s)` returns the whole string; `max==0` returns "".
