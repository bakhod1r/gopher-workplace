# recover requires a deferred context

## The idea

A direct `recover()` call returns nil; only a `recover()` inside a function invoked by `defer` during a panic captures the panic value.

## Why it matters

Misplacing recover silently disables panic handling — the guard looks correct but never works.

## Watch out

- `recover` outside a deferred call always returns nil.
- Put `f()` after scheduling the deferred recover, not before it.
