# Panics during deferred unwinding

## The idea

A panic in a deferred call replaces the panic currently unwinding; the last panic wins, so cleanup code that panics erases the original error.

## Why it matters

Cleanup steps (Close, Rollback) that can panic must contain it, or they mask the real failure.

## Watch out

- A second panic during unwind overwrites the first.
- Keep deferred cleanups panic-free, or recover inside them.
