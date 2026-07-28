# The single-branch if

## The idea

An `if` without `else` handles the exceptional case and falls through to the common return.

## Why it matters

Guard clauses and early returns keep the happy path un-indented.

## Watch out

- `math.MinInt` has no positive counterpart; ignore that edge for this junior task.
- No `else` is needed after a `return`.
