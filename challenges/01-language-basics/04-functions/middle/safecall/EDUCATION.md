# Panic and recover

## The idea

A panic unwinds the stack running deferred calls; `recover` in one of them captures the panic value and resumes normal return.

## Why it matters

Library boundaries convert unexpected panics into errors so a caller isn't crashed.

## Watch out

- `recover` only works when called directly inside a deferred function.
- Don't use panic/recover for ordinary control flow.
