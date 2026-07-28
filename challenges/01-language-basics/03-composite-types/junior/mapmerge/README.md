# Merge Maps

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Merging config layers: a base map plus overrides, where later values win.

## Task

Implement `Merge(a, b)` — a new map, `b` overriding `a` on collisions, inputs
untouched.

## Examples

```go
Merge({x:1,y:2}, {y:20,z:3}) // => {x:1, y:20, z:3}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **make(map)** | Create a fresh map to return. |
| 2 | **Range a map** | `for k, v := range m`. |
| 3 | **Override order** | Copy a first, then b. |

## Hint

`out := make(map[string]int)`; copy a, then copy b.

## Validate

```bash
make verify
```
