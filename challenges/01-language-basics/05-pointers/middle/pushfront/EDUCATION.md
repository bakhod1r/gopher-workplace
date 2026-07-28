# Prepending to a linked list

## The idea

A new node pointing at the current head becomes the new head — an O(1) push front.

## Why it matters

Stack and list builders prepend nodes this way.

## Watch out

- Return the new node; the old head is now second.
- Works with a nil head (empty list).
