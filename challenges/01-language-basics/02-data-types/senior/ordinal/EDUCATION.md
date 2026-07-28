# The teens exception

## The idea

English ordinals key on the last digit (1→st, 2→nd, 3→rd, else th) **except**
11, 12, 13, which are all "th". So check the last two digits first:

```go
if n%100 >= 11 && n%100 <= 13 { return "th" }
switch n % 10 { case 1: "st"; case 2: "nd"; case 3: "rd"; default: "th" }
```

## Why it matters

Ordinals show up in dates and rankings. Keying only on the last digit prints
`11st`/`12nd`/`13rd`, an obvious visible bug — a good lesson that a
"mod 10" rule needs a "mod 100" exception.

## Watch out

- The exception is on `n%100`, catching 111, 112, 113 as well.
- Guard the teens **before** the last-digit switch.
- Negative inputs need their own policy; this version assumes non-negative.
