# Presence, not value

## Intuition

Indexing a map returns the zero value for a missing key, so testing `v != 0`
conflates "absent" with "present but zero". The comma-ok form separates them:

```go
if v, ok := m[key]; ok { return v }
return def
```

## Approach

1. Bug: if v := m[key]; v != 0 treats a stored zero as missing, returning def. 2. A missing key and a stored 0 both yield v==0, so value alone cannot distinguish them. 3. Fix: use comma-ok: if v, ok := m[key]; ok { return v }.

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

For m={a:0}, key=a: m[key] gives (0,true). Old code sees v==0 and falls through to def=7 (wrong). comma-ok sees ok=true and returns 0.

## Pitfalls

- Branch on `ok`, never on the value, to detect presence.
- The same trap exists for `false` and `""` values.
- Reading a nil map returns `(zero, false)` — safe.
