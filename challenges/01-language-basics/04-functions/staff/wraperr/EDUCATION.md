# The defer-wrap-error idiom

## The idea

The idiomatic error decorator captures the NAMED return in the closure body; passing it as a deferred argument freezes the pre-assignment (nil) value.

## Why it matters

This exact mistake makes `defer` error-wrapping silently no-op — a very common real bug.

## Watch out

- `defer func(e error){...}(err)` snapshots err (nil) at defer-time.
- Read the named `err` in the body so you see the final value.
