# Capturing shared loop variables

## The idea

A variable declared outside the loop is shared by every closure; only per-iteration variables (or an explicit inner copy) give each closure its own value.

## Why it matters

Even with Go 1.22's per-iteration range vars, a manually hoisted counter reintroduces the classic capture bug.

## Watch out

- `var i int; for i = 0; ...` shares one `i` across all closures.
- Use `for i := range` or copy into a fresh variable inside the loop.
