# Contiguous alphabet offsets

## The idea

The base64 value table concatenates four runs: 26 uppercase (0-25), 26 lowercase
(26-51), 10 digits (52-61), then `+` (62) and `/` (63). Each run's base is the
running total of the runs before it:

```go
case c >= '0' && c <= '9': return int(c-'0') + 52, true
```

## Why it matters

Codec tables must line up exactly; a single-unit offset on one run silently
corrupts every value in that run, and base64 corruption is invisible until you
decode. It is the same "runs start where the last ended" arithmetic as any
alphabet mapping.

## Watch out

- 'a' starts at 26 (after 26 letters), '0' at 52 — count the runs.
- URL-safe base64 uses `-` and `_` instead of `+` and `/`.
- Reject non-alphabet bytes (including padding `=`) explicitly.
