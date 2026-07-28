# Predicate over a pointer structure

## The idea

Threading a `func(int) bool` through a list traversal generalises counting/filtering to any condition.

## Why it matters

Higher-order traversals keep list algorithms reusable.

## Watch out

- Stop at nil; test `n.Val` each step.
- The predicate decouples the condition from the walk.
