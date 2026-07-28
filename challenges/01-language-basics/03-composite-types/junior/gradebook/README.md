# Gradebook Averages

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A map from student to their scores; produce each average. Students with no
scores are skipped.

## Task

Implement `Averages(book)` (integer average).

## Examples

```go
Averages({ann:{90,80,100}, bob:{70,75}, cid:{}}) // => {ann:90, bob:72}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map of slices** | Value type is `[]int`. |
| 2 | **Guard empty** | Skip zero-length slices. |
| 3 | **Build a map** | Accumulate results. |

## Hint

Range the map; for non-empty slices, sum and divide, store in the result.

## Validate

```bash
make verify
```
