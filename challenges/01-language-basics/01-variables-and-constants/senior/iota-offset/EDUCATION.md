# The iota base expression

## The idea

The **first** line of a `const` block sets the expression that every bare line
below repeats. `iota` starts at 0 on that first line, so the base expression is
evaluated with `iota == 0`:

```go
const (
	Read Permission = 1 << iota // iota==0 -> 1
	Write                       // iota==1 -> 2
	Execute                     // iota==2 -> 4
)
```

If the base expression is skewed — say `1 << (iota + 1)` — the whole run shifts:
Read becomes 2, Write 4, Execute 8. The code compiles and *looks* like flags,
but the values are wrong from the very first one.

## Why it matters

For bit flags the exact value **is** the contract. Doubling every flag makes the
lowest bit (1) unused and can cause masks to overlap with a neighbouring field,
so `Has` and `&` checks silently misbehave.

## Watch out

- Debug flag values by pinning the first: it must be `1 << 0 == 1`.
- The off-by-one hides because the *relationships* still look plausible
  (each is double the last) — only the absolute values are wrong.
- Comment the expected numbers next to each constant during review.

## Try it yourself

```go
const (
	A = 1 << iota // 1  ✓
	B             // 2
)
const (
	X = 1 << (iota + 1) // 2  ✗ starts too high
	Y                   // 4
)
```
