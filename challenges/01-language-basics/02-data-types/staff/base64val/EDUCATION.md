# Contiguous alphabet offsets

## Intuition

The base64 value table concatenates four runs: 26 uppercase (0-25), 26 lowercase
(26-51), 10 digits (52-61), then `+` (62) and `/` (63). Each run's base is the
running total of the runs before it:

```go
case c >= '0' && c <= '9': return int(c-'0') + 52, true
```

## Approach

1. Bug: digit branch adds 53, so '0' yields 53 (off by one).
2. Base64 alphabet: A-Z=0..25, a-z=26..51, 0-9=52..61.
3. Fix: int(c-'0') + 52.

## Solution

```go
func Value(c byte) (int, bool) {
	switch {
	case c >= 'A' && c <= 'Z':
		return int(c - 'A'), true
	case c >= 'a' && c <= 'z':
		return int(c-'a') + 26, true
	case c >= '0' && c <= '9':
		return int(c-'0') + 52, true
	case c == '+':
		return 62, true
	case c == '/':
		return 63, true
	}
	return 0, false
}
```

## Walkthrough

'0'-'0'=0, +52 = 52. '9'-'0'=9, +52 = 61. Ranges now abut correctly.

## Pitfalls

- 'a' starts at 26 (after 26 letters), '0' at 52 — count the runs.
- URL-safe base64 uses `-` and `_` instead of `+` and `/`.
- Reject non-alphabet bytes (including padding `=`) explicitly.
