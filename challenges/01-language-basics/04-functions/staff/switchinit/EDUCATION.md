# The switch init statement

## The idea

`switch init; tag {}` (or `switch init; {}`) runs the init once and scopes its variables to the switch, avoiding repeated evaluation across cases.

## Why it matters

Recomputing an expensive or side-effecting expression per case is a real performance and correctness bug.

## Watch out

- Cases in a tagless switch each re-evaluate their own expression.
- Hoist shared work into the switch init statement.
