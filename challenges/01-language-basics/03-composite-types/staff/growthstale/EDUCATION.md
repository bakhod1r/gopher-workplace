# append can invalidate pointers

## The idea

When `append` exceeds capacity, it allocates a **new** backing array and copies.
Any pointer or sub-slice into the old array still references the old memory:

```go
p := &s[0]
s = append(s, x) // may reallocate
s[0] = 99        // correct: writes the current array (don't use *p)
```

## Why it matters

Holding a pointer (or sub-slice) across a growth is a real, subtle bug: writes go
to a stale array and silently vanish, while the GC keeps the old array alive.
After any append that may grow, re-derive addresses.

## Watch out

- Growth is capacity-triggered; with spare cap, `p` would still be valid — making
  the bug intermittent.
- Sub-slices taken before an append can likewise detach.
- Prefer indices over long-lived element pointers into growable slices.
