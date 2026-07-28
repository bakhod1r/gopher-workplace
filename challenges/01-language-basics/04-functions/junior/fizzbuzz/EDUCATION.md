# Ordering conditional branches

## The idea

When cases overlap, the most specific condition must come first or a broader branch will shadow it.

## Why it matters

Rule ordering bugs are common in pricing, routing, and validation ladders.

## Watch out

- Checking `%3` before `%15` prints "Fizz" for 15 — wrong.
- `strconv.Itoa` (not string(n)) converts the integer to its decimal text.
