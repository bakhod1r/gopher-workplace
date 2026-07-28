# Chunk a Slice

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

A batch processor sends records in groups of N. Split a slice into fixed-size
chunks, the last possibly short.

## Task

Implement `Chunk(xs, size)`; `size <= 0` → empty.

## Examples

```go
Chunk([]int{1,2,3,4,5}, 2) // => [[1 2] [3 4] [5]]
Chunk([]int{1,2,3}, 5)     // => [[1 2 3]]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice windows** | `xs[i:end]` where end is clamped. |
| 2 | **Step by size** | Advance i by size each iteration. |
| 3 | **Clamp last** | `min(i+size, len(xs))`. |

## Hint

`for i := 0; i < len(xs); i += size { end := i+size; if end > len(xs) { end = len(xs) }; out = append(out, xs[i:end]) }`.

## Validate

```bash
make verify
```
