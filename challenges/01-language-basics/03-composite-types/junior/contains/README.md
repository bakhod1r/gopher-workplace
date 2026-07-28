# Slice Contains

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Checking membership in a small slice is a linear scan — return as soon as you
find it.

## Task

Implement `Contains(xs, target)`.

## Examples

```go
Contains([]string{"a","b"}, "a") // => true
Contains([]string{"a","b"}, "z") // => false
Contains(nil, "x")               // => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Linear scan** | Compare each element. |
| 2 | **Early return** | Return true on the first match. |
| 3 | **Default false** | No match after the loop → false. |

## Hint

`for _, x := range xs { if x == target { return true } }; return false`.

## Validate

```bash
make verify
```
