# Initialising nested maps

## The idea

Only the outer map exists after construction; each inner map is nil until assigned, and writing to nil panics.

## Why it matters

Grouped counters and two-level indexes hit this nil-inner-map crash.

## Watch out

- `m[group]` is nil, not empty, until you assign a map.
- Create the inner map before incrementing.
