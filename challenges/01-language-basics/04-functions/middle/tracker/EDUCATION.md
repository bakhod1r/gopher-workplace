# Multiple closures over shared state

## The idea

Closures defined in one scope share its variables, letting a factory return several functions that operate on the same hidden state.

## Why it matters

This is how you build small stateful objects without a struct.

## Watch out

- Both returned functions capture the SAME `sum` by reference.
- The state is unreachable except through the returned closures.
