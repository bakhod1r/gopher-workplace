# Element pointers invalidated by append

## The idea

A reallocating append detaches pointers taken beforehand; use the current slice's index or re-take the address afterward.

## Why it matters

Holding element pointers across appends is a real aliasing bug.

## Watch out

- After a reallocating append, `&s[0]` from before is stale.
- Index the current slice (`s[0]`) or re-take the pointer.
