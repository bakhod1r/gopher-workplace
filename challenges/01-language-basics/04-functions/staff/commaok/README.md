# Ignoring comma-ok

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Indexing a map returns the zero value for a missing key, indistinguishable from
a stored zero. The two-result form `v, ok := m[k]` reports presence; hardcoding
`ok = true` misclassifies absent keys.

## Task

Fix [commaok.go](commaok.go) so absent keys report ok=false.

Do **not** change the function signature or the tests.

## Examples

```go
Lookup({a:0,b:5}, "a") // => 0, true
Lookup({a:0,b:5}, "z") // => 0, false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **comma-ok map read** | `v, ok := m[k]` reports presence. |
| 2 | **Zero vs absent** | A stored 0 and a missing key both index to 0. |
| 3 | **Propagate ok** | Return the map's ok, not a constant. |

## Hint

Use the two-result form: `score, ok = scores[name]`.

## Validate

```bash
make verify
```
