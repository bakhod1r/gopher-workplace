# One-shot execution via captured state

## The idea

A captured boolean lets a closure enforce run-once semantics across many calls.

## Why it matters

Lazy initialisation and idempotent setup use exactly this guard.

## Watch out

- Set the flag BEFORE or after calling f consistently; here order doesn't matter (single-threaded).
- Not goroutine-safe; that's `sync.Once`'s job.
