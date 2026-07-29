# Index windows and inverted conditions

## Intuition

The last four characters are indices `i >= len-4`; everything before is
`i < len-4`. To hide all but the last four, mask the **front**:

```go
if i < len(r)-4 { out[i] = '*' } else { out[i] = r[i] }
```

## Approach

1. Bug: condition `i >= len(r)-4` masked the LAST four and revealed the front.
2. Fix: `i < len(r)-4` masks everything except the last four.
3. Strings of <=4 runes return unchanged.

## Solution

```go
func Mask(s string) string {
	r := []rune(s)
	if len(r) <= 4 {
		return s
	}
	out := make([]rune, len(r))
	for i := range r {
		if i < len(r)-4 {
			out[i] = '*'
		} else {
			out[i] = r[i]
		}
	}
	return string(out)
}
```

## Walkthrough

"12345": len 5, mask indices <1 (index 0) -> '*', keep 2345 -> "*2345".

## Pitfalls

- Operate on runes so multi-byte input isn't miscounted.
- Return short inputs unchanged, or you might reveal everything (define the
  policy deliberately).
- The boundary is `len-4`; off-by-one reveals or hides one extra digit.
