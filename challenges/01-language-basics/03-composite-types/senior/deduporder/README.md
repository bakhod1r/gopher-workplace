# Dedupe Preserving Order

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

The condition is inverted: it appends only when the value was **already seen**, so
the output is exactly the duplicates.

## Task

Fix the condition between the markers in [deduporder.go](deduporder.go) to keep
first occurrences.

## Examples

```go
Unique([]int{3,1,3,2,1,4}) // => [3 1 2 4]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Seen set** | Track values already emitted. |
| 2 | **First occurrence** | Emit when NOT seen. |
| 3 | **Order preserved** | Slice keeps insertion order. |

## Hint

`if !ok`.

## Validate

```bash
make verify
```
