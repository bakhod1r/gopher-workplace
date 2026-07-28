# Deep vs shallow copy

## The idea

Copying the head pointer shares the whole list; a deep copy allocates a new node per element so mutations don't leak.

## Why it matters

Deep-copying linked structures is required before independent edits.

## Watch out

- Returning `head` (shallow) shares every node.
- Allocate a new node for each; recurse on Next.
