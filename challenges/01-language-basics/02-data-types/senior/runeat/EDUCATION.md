# Index and bound must agree

## Intuition

`rs := []rune(s)` gives rune-indexed access. The bounds check must use
`len(rs)` (rune count), not `len(s)` (byte count):

```go
if n < 0 || n >= len(rs) { return 0, false }
```

## Approach

1. Bug: bounds used `n >= len(s)` (byte length), allowing indices past the rune count for multibyte strings.
2. Fix: compare against `len(rs)`, the rune-slice length.
3. rs[n] then returns the n-th rune.

## Solution

```go
func At(s string, n int) (rune, bool) {
	rs := []rune(s)
	if n < 0 || n >= len(rs) {
		return 0, false
	}
	return rs[n], true
}
```

## Walkthrough

At("日本",2): rs has 2 runes, 2>=len(rs)=2 -> false.

## Pitfalls

- `len(s)` ≥ `len([]rune(s))`, equal only for pure ASCII.
- The conversion `[]rune(s)` allocates; fine for random access.
- Any time you index one representation, bound with the same one.
