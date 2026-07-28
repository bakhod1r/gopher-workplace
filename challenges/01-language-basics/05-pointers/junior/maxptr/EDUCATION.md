# Returning a selected pointer

## The idea

Comparing dereferenced values then returning one address hands the caller a live reference to the winner.

## Why it matters

Selecting a mutable element (max, chosen node) by pointer is common in data structures.

## Watch out

- Compare `*a`/`*b`, but return `a`/`b`.
- The returned pointer still aliases the caller's variable.
