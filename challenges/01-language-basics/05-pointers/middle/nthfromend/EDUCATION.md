# The two-pointer gap technique

## The idea

Keeping a fixed n-node gap between two pointers locates the n-th-from-end in one pass without counting length first.

## Why it matters

Gap pointers solve nth-from-end, window, and remove-nth problems.

## Watch out

- If advancing the lead n steps runs past the end, the list is too short.
- Both pointers move in lockstep after the gap is set.
