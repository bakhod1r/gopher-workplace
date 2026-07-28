# Clip Capacity

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Handing out a slice with spare capacity is dangerous: the receiver's `append` can
overwrite memory you still use. `Clip` should cap capacity to length via a
full-slice expression. The code returns `xs` unchanged.

## Task

Fix the return between the markers in [clipcap.go](clipcap.go).

## Examples

```go
c := Clip(make([]int,3,10)) // cap(c) == 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Three-index slice** | `xs[:len:len]` caps capacity. |
| 2 | **Append safety** | No spare cap → append reallocates. |
| 3 | **Defensive slicing** | Protect shared backing arrays. |

## Hint

`return xs[:len(xs):len(xs)]`.

## Validate

```bash
make verify
```
