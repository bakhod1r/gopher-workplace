# Joining linked lists

## The idea

Linking the first list's tail to the second's head concatenates them in O(len a); the empty first-list case returns the second directly.

## Why it matters

List concatenation appears in queue merging and rope structures.

## Watch out

- Handle a nil first list.
- Only a's tail's Next changes.
