# The goto statement

## The idea

`goto` transfers control to a label in the same function; it cannot jump into a block or over variable declarations.

## Why it matters

Rare in Go, but it appears in generated code and tight state machines; understanding it clarifies how loops desugar.

## Watch out

- `goto` can't jump over a variable's declaration into its scope.
- Prefer `for`; use `goto` only when it genuinely reads better.
