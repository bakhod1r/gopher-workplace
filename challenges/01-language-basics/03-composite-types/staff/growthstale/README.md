# Stale Pointer After Growth

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

`append` to a full slice **reallocates** the backing array. The pointer `p`,
taken before the append, still points at the old array — writing `*p = 99` updates
freed-from-view memory, not `s[0]`.

## Task

Fix the write between the markers in [growthstale.go](growthstale.go) to update
the current slice.

## Examples

```go
BumpFirst(20) // => 99, [99 20]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **append reallocates** | When cap is exceeded, a new array is made. |
| 2 | **Stale pointers** | Old element pointers no longer alias the slice. |
| 3 | **Re-address** | Use `s[0]` (or re-take `&s[0]`) after growth. |

## Hint

Re-take the address after the append: `p = &s[0]` then `*p = 99` (or write
`s[0] = 99` directly and drop `p`).

## Validate

```bash
make verify
```
