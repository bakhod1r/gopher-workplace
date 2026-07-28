# Loop-variable capture semantics

## The idea

Go 1.22 gives each iteration its own loop variable, so closures capture distinct values; pre-1.22 they all shared the final value.

## Why it matters

This removed one of Go's most common footguns in goroutine/closure loops.

## Watch out

- On Go <1.22 you would need `i := i` to shadow; on 1.26 it's automatic.
- Each closure returns its own captured index.
