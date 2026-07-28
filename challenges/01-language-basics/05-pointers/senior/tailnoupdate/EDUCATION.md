# Growing an empty list

## The idea

With a `*Node` signature the empty case must return the new head; returning nil loses the value.

## Why it matters

Empty-collection edge cases are a frequent source of list-append bugs.

## Watch out

- An empty list has no tail — build the head instead.
- Return it so the caller can adopt the new head.
