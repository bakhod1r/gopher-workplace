# Bucketing with a key function

## The idea

`append(m[k], x)` relies on a missing key yielding a nil slice that append grows — no pre-init per key needed.

## Why it matters

Group-by is a staple of aggregation and indexing.

## Watch out

- You must init the outer map; per-key slices need no init.
- Input order is preserved within each bucket by appending in order.
