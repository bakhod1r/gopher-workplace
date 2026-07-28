# Nil function values

## The idea

A func variable defaults to nil and panics on call; optional callbacks must be checked before invocation.

## Why it matters

Invoking an unset optional callback is a real nil-call panic in plugin/hook APIs.

## Watch out

- A nil func value is not callable — check `hook != nil` first.
- Provide a sensible default (identity) for the nil case.
