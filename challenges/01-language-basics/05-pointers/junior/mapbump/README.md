# Mutate Through Map of Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Map values that are pointers let you mutate shared variables — the map stores
addresses, not copies.

## Task

Implement `BumpAll` in [mapbump.go](mapbump.go).

Do **not** change the function signature or the tests.

## Examples

```go
BumpAll(map[string]*int{"a":&a}) // *m["a"]++
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map of pointers** | Values are addresses. |
| 2 | **Range values** | `for _, p := range m`. |
| 3 | **Mutate through** | `*p++`. |

## Hint

Range values; `*p++`.

## Validate

```bash
make verify
```
