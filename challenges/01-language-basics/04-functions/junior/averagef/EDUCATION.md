# Guarding aggregation

## The idea

An average needs at least one sample; returning a validity flag lets callers avoid a division-by-zero or a misleading 0.

## Why it matters

Real aggregations (metrics, ratings) must distinguish 'no data' from 'average is zero'.

## Watch out

- Dividing by `len(nums)` as an int truncates; convert to `float64`.
- Guard the empty slice before the division.
