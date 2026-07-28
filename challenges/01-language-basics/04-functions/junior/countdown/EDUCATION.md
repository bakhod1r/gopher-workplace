# The classic three-clause for loop

## The idea

`for init; cond; post` is Go's only loop keyword; it covers counting up, down, and while-style conditions.

## Why it matters

Explicit counters are needed whenever you build or transform by index rather than by range.

## Watch out

- Off-by-one: use `i >= 1` to include 1 and stop before 0.
- Return a non-nil empty slice for `n <= 0` (len 0 is enough here).
