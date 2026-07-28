# Initialising accumulators

## The idea

The loop counts every node, so the counter must start at 0; a start of 1 double-counts and breaks the empty case.

## Why it matters

Off-by-one initial values are a classic counting bug.

## Watch out

- `count := 1` makes an empty list report length 1.
- Let the loop do all the counting from 0.
