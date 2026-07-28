# Scoping defer to a loop iteration

## The idea

Deferred calls run at function return, so a `defer` in a loop body accumulates to the end; wrap each iteration in its own function literal to make the defer fire per iteration.

## Why it matters

Expecting per-iteration cleanup from a function-level defer causes wrong ordering and resource over-holding.

## Watch out

- A loop-body `defer` runs at function exit, not each iteration's end.
- `func(){ ...; defer end() }()` per iteration gives interleaved order.
