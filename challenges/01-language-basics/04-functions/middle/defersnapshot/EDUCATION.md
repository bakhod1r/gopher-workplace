# When defer evaluates its arguments

## The idea

The deferred function's arguments are evaluated when `defer` executes; only the call is postponed.

## Why it matters

This distinction decides whether a deferred log/cleanup sees the old or new value — a frequent source of confusion.

## Watch out

- `defer f(x)` snapshots `x`; `defer func(){ use(x) }()` reads it at return.
- Here the argument form gives 1.
