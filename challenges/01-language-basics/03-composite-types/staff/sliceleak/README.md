# Sub-slice Memory Leak

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Returning `xs[:3]` keeps a reference to the **entire** backing array — even if the
source had a million-element capacity, none of it can be freed while the small
head lives. Copy to release it.

## Task

Fix the return between the markers in [sliceleak.go](sliceleak.go) to return an
independent 3-element slice.

## Examples

```go
h := Head3(bigSlice) // cap(h) == 3, source array can be GC'd
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Sub-slice retains array** | Cap covers the whole backing array. |
| 2 | **GC reachability** | The array lives as long as any slice into it. |
| 3 | **Copy to release** | A fresh 3-cap slice frees the source. |

## Hint

`append([]int{}, xs[:3]...)` (or make+copy).

## Validate

```bash
make verify
```
