# Deferred mutation of named returns

## The idea

Because defers run after the return value is assigned, a deferred closure that writes the named result overrides whatever the body computed.

## Why it matters

A stray or mis-ordered deferred assignment is a subtle way to lose a return value.

## Watch out

- Deferred writes to a named return WIN over the body's assignment.
- Only mutate the result in a defer when you intend to (e.g. wrapping errors).
