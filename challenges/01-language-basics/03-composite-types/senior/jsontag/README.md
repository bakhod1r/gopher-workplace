# JSON Tag Mismatch

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

The API contract is snake_case. `LastName` is tagged `lastName`, so the JSON key
is wrong and clients break.

## Task

Fix the struct tag between the markers in [jsontag.go](jsontag.go).

## Examples

```go
Marshal(User{"Ada","Lovelace"}) // => {"first_name":"Ada","last_name":"Lovelace"}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Struct tags** | `` `json:"name"` `` renames the key. |
| 2 | **Exact key** | The tag string is the literal JSON key. |
| 3 | **encoding/json** | Reads tags via reflection. |

## Hint

`` `json:"last_name"` ``.

## Validate

```bash
make verify
```
