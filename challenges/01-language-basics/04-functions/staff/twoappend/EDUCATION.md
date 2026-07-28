# Append aliasing through shared capacity

## The idea

Two appends to a base with spare capacity write the same memory; a three-index slice (`s[:len:len]`) caps it so the next append copies.

## Why it matters

This is a real corruption bug when one buffer seeds multiple derived slices.

## Watch out

- Appending to a base with spare cap twice makes the results alias.
- `base[:len:len]` guarantees the next append reallocates.
