# Dropping references with nil

## The idea

Assigning nil to a pointer field removes the link; if no other reference remains, the target becomes eligible for garbage collection.

## Why it matters

Detaching nodes prevents memory leaks in long-lived structures.

## Watch out

- Only the Next field changes; Val stays.
- Nil-ing references is how you let the GC reclaim memory.
