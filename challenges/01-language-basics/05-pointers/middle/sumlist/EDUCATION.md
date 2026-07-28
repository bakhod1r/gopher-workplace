# Recursion over pointer structures

## The idea

Linked lists and trees recurse naturally: base-case on nil, combine the current node with the recursive result.

## Why it matters

Tree/list aggregations (sum, height, count) are recursive pointer walks.

## Watch out

- Base-case nil BEFORE dereferencing.
- Each call handles one node plus the rest.
