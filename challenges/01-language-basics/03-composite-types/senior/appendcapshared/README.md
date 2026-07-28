# Append Clobbers via Shared Capacity

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`xs[:2]` shares `xs`'s backing array **and its spare capacity**. Appending to it
writes into `xs[2]`, corrupting the source. A full-slice expression caps the
sub-slice so append must reallocate.

## Task

Fix the sub-slice between the markers in
[appendcapshared.go](appendcapshared.go) to not share spare capacity.

## Examples

```go
xs := []int{1,2,3} (cap>3); FirstTwoPlus(xs,99) // [1 2 99], xs[2] stays 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Sub-slice capacity** | `xs[:2]` keeps `xs`'s capacity. |
| 2 | **Full-slice expr** | `xs[:2:2]` caps capacity to 2. |
| 3 | **Force realloc** | Append then can't touch xs. |

## Hint

`head := xs[:2:2]` (three-index slice caps capacity).

## Validate

```bash
make verify
```
