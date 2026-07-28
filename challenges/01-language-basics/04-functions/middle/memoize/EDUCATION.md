# Stateful closures for caching

## The idea

A closure holding a map memoises across calls; the state is private and survives between invocations.

## Why it matters

Memoisation trades memory for speed on expensive pure functions.

## Watch out

- Use comma-ok to distinguish a stored 0 from a missing key.
- Only correct for pure functions; side-effecting `f` breaks caching semantics.
