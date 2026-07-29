# The iota base expression

## Intuition

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

## Approach

1. The first flag must be `1 << iota` (iota 0 → bit 0).
2. The bug's `1 << (iota + 1)` shifts every flag up one, so Read becomes 2.
3. Drop the `+ 1`.

## Solution

```go
type Permission uint8

const (
	Read Permission = 1 << iota
	Write
	Execute
)

func Has(set, want Permission) bool { return set&want == want }
```

## Walkthrough

With the offset, Read is 2 and masks fail; `1 << iota` restores Read=1, Write=2, Execute=4.

## Pitfalls

- Debug flag values by pinning the first: it must be `1 << 0 == 1`.
- The off-by-one hides because the *relationships* still look plausible
  (each is double the last) — only the absolute values are wrong.
- Comment the expected numbers next to each constant during review.
