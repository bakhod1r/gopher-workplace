# Rotate a Slice

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

A round-robin scheduler rotates the queue. Rotate left by k, wrapping, for any k.

## Task

Implement `Left(xs, k)` returning a new rotated slice.

## Examples

```go
Left([]int{1,2,3,4,5}, 2) // => [3 4 5 1 2]
Left([]int{1,2,3}, -1)    // => [3 1 2]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Normalize k** | `((k % n) + n) % n`. |
| 2 | **Reassemble** | `xs[k:]` then `xs[:k]`. |
| 3 | **Fresh slice** | Build a new result. |

## Hint

`n := len(xs)`; if n==0 return empty; `k = ((k%n)+n)%n`; `append(append([]int{}, xs[k:]...), xs[:k]...)`.

## Validate

```bash
make verify
```
