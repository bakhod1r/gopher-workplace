# Accumulator threading in recursion

## The idea

Tail-recursive folds pass the running result forward each call; replacing the accumulator instead of adding to it loses all prior work.

## Why it matters

Accumulator bugs return plausible-looking values (here the last element) that pass shallow tests.

## Watch out

- The recursive call must combine `acc` with the current element.
- Passing only `xs[0]` restarts the fold.
