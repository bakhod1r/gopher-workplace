# Sorted Map Keys

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Map iteration order is random, so to print or compare deterministically you
collect the keys and sort them.

## Task

Implement `Sorted(m)` returning keys in ascending order.

## Examples

```go
Sorted({banana:1, apple:2}) // => [apple banana]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Collect keys** | Range the map, append `k`. |
| 2 | **sort.Strings** | Sorts a []string in place. |
| 3 | **Random order** | Map ranging is intentionally unordered. |

## Hint

Append each key, then `sort.Strings(out)`.

## Validate

```bash
make verify
```
