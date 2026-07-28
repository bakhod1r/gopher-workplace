# Slice Equality

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

Slices can't be compared with `==` (that's a compile error), so equality must be
element-wise. The code checks only the length, so different contents of equal
length compare "equal".

## Task

Fix the body between the markers in [sliceequal.go](sliceequal.go) to compare
elements.

## Examples

```go
Equal([]int{1,2,3}, []int{1,9,3}) // => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **No slice ==** | Only nil comparison is allowed. |
| 2 | **Element-wise** | Loop and compare each index. |
| 3 | **Length first** | Fast reject on differing lengths. |

## Hint

`for i := range a { if a[i] != b[i] { return false } }; return true`.

## Validate

```bash
make verify
```
