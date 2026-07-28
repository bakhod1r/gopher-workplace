# Ordering defer relative to acquisition

## The idea

A `defer cleanup()` placed before the acquisition check runs on every exit path, including failures where nothing was acquired.

## Why it matters

Deferring Close before confirming the open succeeded double-frees or closes nil resources — a real bug.

## Watch out

- Put `defer resource.Close()` immediately AFTER a successful open, not before.
- A defer above an early error return still executes.
