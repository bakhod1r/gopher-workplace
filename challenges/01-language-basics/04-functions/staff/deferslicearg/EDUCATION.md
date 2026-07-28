# Slice headers as deferred arguments

## The idea

A deferred argument snapshots the slice header (including length) at defer-time; since append reassigns the variable, only a body capture sees the final slice.

## Why it matters

Logging a slice's length/contents via a deferred argument records the pre-build state — a subtle reporting bug.

## Watch out

- `defer f(xs)` freezes xs's header (len 0) before the appends.
- Reference `xs` in the closure body for the final value.
