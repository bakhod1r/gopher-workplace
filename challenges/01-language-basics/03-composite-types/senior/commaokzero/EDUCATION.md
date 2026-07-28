# Presence, not value

## The idea

Indexing a map returns the zero value for a missing key, so testing `v != 0`
conflates "absent" with "present but zero". The comma-ok form separates them:

```go
if v, ok := m[key]; ok { return v }
return def
```

## Why it matters

Config values, counters, and flags legitimately hold 0/""/false. Using the value
itself as a presence signal silently drops real zero entries — a subtle,
data-dependent bug.

## Watch out

- Branch on `ok`, never on the value, to detect presence.
- The same trap exists for `false` and `""` values.
- Reading a nil map returns `(zero, false)` — safe.
