# The comma-ok idiom

## The idea

Indexing a map always returns a value — the zero value if the key is absent. To
distinguish "present with value 0" from "missing", use the two-result form:

```go
if v, ok := m[key]; ok { return v }
return def
```

## Why it matters

Config, caches, and counters constantly need "present?" separate from "value".
Relying on the zero value silently treats a real `0`/`""`/`false` as missing — a
frequent bug.

## Watch out

- Reading a **nil** map is safe and reports absent; only *writing* nil panics.
- `ok` is the second result; ignoring it loses the distinction.
- The same idiom appears for type assertions and channel receives.
