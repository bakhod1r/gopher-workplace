# Deferred call ordering

## The idea

Each `defer` pushes onto a stack unwound in reverse at function exit; deferred closures see and can modify named return values.

## Why it matters

Cleanup that must mirror acquisition order (locks, files, transactions) relies on this LIFO guarantee.

## Watch out

- Schedule order is source order; execution order is reversed.
- Append 1,2,3 in source → result [3 2 1].
