# Snapshot vs live capture in defer

## The idea

Passing a value as a deferred argument freezes it at defer-time; referencing it in the closure body reads it at return-time.

## Why it matters

Deferred logging/metrics that pass the value by argument record stale data — a subtle, common bug.

## Watch out

- `defer f(c)` uses c's value NOW; `defer func(){ use(c) }()` uses it at return.
- For final values, capture in the body.
