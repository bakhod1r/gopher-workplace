# Nil Map JSON

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A nil map marshals to JSON `null`; an empty map to `{}`. `Counts` starts with a
nil map and only allocates inside the loop, so **empty input** returns nil → `null`,
breaking clients expecting an object.

## Task

Fix the declaration between the markers in [mapniljson.go](mapniljson.go) to
allocate upfront.

## Examples

```go
json.Marshal(Counts(nil)) // must be "{}", not "null"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil vs empty map** | Both len 0; encode differently. |
| 2 | **JSON encoding** | nil → null, empty → {}. |
| 3 | **Allocate upfront** | `make(map[string]int)`. |

## Hint

`m := make(map[string]int)`.

## Validate

```bash
make verify
```
