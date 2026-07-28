# The comma-ok return pattern

## The idea

A trailing boolean result reports whether the other values are meaningful, mirroring map lookups and type assertions.

## Why it matters

Callers can branch on presence without magic numbers, keeping the empty and populated cases explicit.

## Watch out

- Reading `xs[0]` before the length guard panics on an empty slice.
- Return the zero values (`0, 0`) with `false`, not garbage.
