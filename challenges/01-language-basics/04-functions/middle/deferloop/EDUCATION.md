# Deferring within loops

## The idea

Every `defer` in a loop is stacked; they all run at function return, not at each iteration's end — reverse order overall.

## Why it matters

Deferring `Close` in a loop keeps all resources open until the function ends — sometimes a leak, sometimes intended.

## Watch out

- Defers do NOT run at the end of each iteration — only at function return.
- For per-iteration cleanup, wrap the body in its own function.
