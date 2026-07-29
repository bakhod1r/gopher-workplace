# The teens exception

## Intuition

English ordinals key on the last digit (1→st, 2→nd, 3→rd, else th) **except**
11, 12, 13, which are all "th". So check the last two digits first:

```go
if n%100 >= 11 && n%100 <= 13 { return "th" }
switch n % 10 { case 1: "st"; case 2: "nd"; case 3: "rd"; default: "th" }
```

## Approach

1. Bug: `switch n % 10` ignored the 11-13 exception, giving 11->"st", 12->"nd", 13->"rd".
2. Fix: first check `n % 100` for 11,12,13 and return "th", then fall through to the mod-10 switch.
3. 21 -> n%100=21 (no match) -> n%10=1 -> "st"; 11 -> "th".

## Solution

```go
import "strconv"

func Suffix(n int) string {
	switch n % 100 {
	case 11, 12, 13:
		return "th"
	}
	switch n % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func Format(n int) string { return strconv.Itoa(n) + Suffix(n) }
```

## Walkthrough

Suffix(11): 11%100==11 -> "th". Suffix(21): 21%100=21 skip, 21%10=1 -> "st".

## Pitfalls

- The exception is on `n%100`, catching 111, 112, 113 as well.
- Guard the teens **before** the last-digit switch.
- Negative inputs need their own policy; this version assumes non-negative.
