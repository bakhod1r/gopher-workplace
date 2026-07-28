# Predicate-driven filtering

## The idea

Accepting a `func(T) bool` generalises selection so callers supply the rule, not a new function per condition.

## Why it matters

Filter/where operations across collections and queries share this shape.

## Watch out

- Return empty (not nil-panic) when nothing matches.
- Don't mutate the input slice.
