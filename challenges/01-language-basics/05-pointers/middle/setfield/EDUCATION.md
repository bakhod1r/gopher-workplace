# Fetch-then-mutate on pointer maps

## The idea

The comma-ok read distinguishes absence; the fetched pointer allows in-place mutation of the stored object.

## Why it matters

Update-by-key APIs over object registries follow this pattern.

## Watch out

- Guard the missing key before dereferencing.
- The fetched pointer aliases the stored object.
