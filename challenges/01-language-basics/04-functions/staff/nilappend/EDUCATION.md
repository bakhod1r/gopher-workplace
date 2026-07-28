# nil slices are valid append targets

## The idea

A nil slice has len 0 and appends normally; special-casing it is unnecessary and here actively wrong.

## Why it matters

Over-defensive nil-slice guards discard data or complicate code that would just work.

## Watch out

- `append(nil, ...)` allocates a fresh slice — no pre-init needed.
- Ranging or `len` on a nil slice is also fine.
