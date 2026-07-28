# Zip to Map

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Pairing parallel arrays (column names + values) into a map.

## Task

Implement `Zip(keys, vals)` over the shorter length.

## Examples

```go
Zip([]string{"a","b"}, []int{1,2,3}) // => {a:1, b:2}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Parallel index** | `keys[i]` with `vals[i]`. |
| 2 | **Min length** | Stop at the shorter slice. |
| 3 | **Build a map** | Insert each pair. |

## Hint

`n := min(len(keys), len(vals)); for i := 0; i < n; i++ { m[keys[i]] = vals[i] }`.

## Validate

```bash
make verify
```
