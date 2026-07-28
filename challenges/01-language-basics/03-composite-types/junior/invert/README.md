# Invert a Map

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Building a reverse lookup: given name→id, produce id→name.

## Task

Implement `Invert(m)` returning value→key (values unique).

## Examples

```go
Invert({one:1, two:2}) // => {1:"one", 2:"two"}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **make with new types** | Result key/value types swap. |
| 2 | **Range k,v** | Read both and store `out[v]=k`. |
| 3 | **Uniqueness** | Duplicate values would collide. |

## Hint

`out := make(map[int]string); for k, v := range m { out[v] = k }`.

## Validate

```bash
make verify
```
