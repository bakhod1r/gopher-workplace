# The comma-ok idiom

## Intuition

Indexing a map always returns a value — the zero value if the key is absent. To
distinguish "present with value 0" from "missing", use the two-result form:

```go
if v, ok := m[key]; ok { return v }
return def
```

## Approach

1. Read v, ok := m[key] using the comma-ok form.
2. If ok, the key exists — return v (even if it is 0).
3. Otherwise return def.

## Solution

```go
func GetOr(m map[string]int, key string, def int) int {
	if v, ok := m[key]; ok {
		return v
	}
	return def
}
```

## Walkthrough

GetOr({"zero":0},"zero",99): comma-ok gives v=0, ok=true, so return 0 rather than the default.

## Pitfalls

- Reading a **nil** map is safe and reports absent; only *writing* nil panics.
- `ok` is the second result; ignoring it loses the distinction.
- The same idiom appears for type assertions and channel receives.
