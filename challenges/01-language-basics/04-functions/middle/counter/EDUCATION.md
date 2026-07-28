# Closures capturing enclosing state

## The idea

A function literal keeps a live reference to the variables it uses from the outer scope; those variables survive as long as the closure does.

## Why it matters

Stateful callbacks, generators, memoisers, and iterators are all built this way.

## Watch out

- Each call to the factory creates a NEW captured variable — instances don't share state.
- The captured variable escapes to the heap; that's expected.
