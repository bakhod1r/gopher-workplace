# Conditional in-place mutation

## The idea

Reading `*p`, testing bounds, and writing back constrains the caller's value without a return.

## Why it matters

In-place clamps appear in signal processing and UI value guards.

## Watch out

- Only assign when out of range (optional micro-optimisation).
- Endpoints are inclusive.
