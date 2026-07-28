# Defer runs at function scope, not block scope

## The idea

A `defer` inside a loop is deferred to the enclosing function's return, so resources accumulate until then — sometimes the intent, sometimes a leak.

## Why it matters

Misjudging when deferred cleanup runs causes both over-holding (leaks) and miscounts.

## Watch out

- Defers do not run at each iteration's end.
- For per-iteration cleanup, extract the body into its own function.
