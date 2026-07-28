# Labeled break vs continue

## The idea

A label lets `break`/`continue` act on an enclosing loop; `break Label` terminates it while `continue Label` moves to its next iteration — a one-word difference with opposite effects.

## Why it matters

Confusing labeled break with labeled continue aborts an entire scan instead of skipping one item.

## Watch out

- `break Rows` ends the outer loop; `continue Rows` skips the current outer iteration.
- Pick the verb that matches 'skip this row' vs 'stop scanning'.
