# Filter and Project

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Selecting records by a condition and pulling one field — filter + map.

## Task

Implement `ActiveNames(users)` returning names of active users.

## Examples

```go
ActiveNames([{ann,true},{bob,false},{cid,true}]) // => [ann cid]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Filter** | Keep matching structs. |
| 2 | **Project** | Extract one field. |
| 3 | **append** | Build the result slice. |

## Hint

`for _, u := range users { if u.Active { out = append(out, u.Name) } }`.

## Validate

```bash
make verify
```
