# Seeding a running extremum

## The idea

Seeding from the last element with a strict `<` comparison can return a later tie; seed from the first element so the earliest minimum wins.

## Why it matters

Seed/comparison mismatches produce subtle wrong-index results on ties.

## Watch out

- Seeding `&xs[len-1]` with `<` returns the last min on ties.
- Seed `&xs[0]` to honour the earliest-tie rule.
