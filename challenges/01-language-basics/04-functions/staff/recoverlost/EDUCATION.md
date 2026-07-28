# Setting named returns inside deferred recover

## The idea

The panic-to-error idiom must write the NAMED result variable in the deferred closure; a local assignment is invisible to the caller.

## Why it matters

Losing the recovered error to a local is a subtle way to report success after a crash.

## Watch out

- Deferred recover must set the named return, not a fresh local.
- `err = ...` in the closure body reaches the caller.
