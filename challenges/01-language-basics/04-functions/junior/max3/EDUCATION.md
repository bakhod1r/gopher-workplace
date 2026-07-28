# Reducing several values

## The idea

A running accumulator seeded from the first input generalises to any number of comparisons.

## Why it matters

The pattern extends directly to `MinN`/`MaxN` over slices.

## Watch out

- Seeding `m := 0` fails when every argument is negative.
- Use `>` (or `>=`) consistently.
