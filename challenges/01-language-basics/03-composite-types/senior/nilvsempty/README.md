# Nil vs Empty Slice

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

A `nil` slice and an empty slice both have length 0, but they encode differently:
`nil` marshals to JSON `null`, an empty slice to `[]`. The function returns nil
when nothing matches, breaking the API contract.

## Task

Fix the declaration between the markers in [nilvsempty.go](nilvsempty.go) to start
from a non-nil empty slice.

## Examples

```go
json.Marshal(NonEmpty([]string{"",""})) // must be "[]", not "null"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil vs empty** | Same len, different identity. |
| 2 | **JSON encoding** | nil → null, empty → []. |
| 3 | **Init non-nil** | `[]string{}` or `make`. |

## Hint

`out := []string{}`.

## Validate

```bash
make verify
```
