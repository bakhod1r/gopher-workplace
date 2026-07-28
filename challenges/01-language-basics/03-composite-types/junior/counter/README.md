# Frequency Counter

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Counting occurrences relies on the map zero value: a missing key reads as 0, so
`m[k]++` just works.

## Task

Implement `Count(xs)` returning element→count.

## Examples

```go
Count([]string{"a","b","a"}) // => {a:2, b:1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map zero value** | Missing key reads as 0. |
| 2 | **m[k]++** | Read-modify-write in one step. |
| 3 | **make first** | Allocate before writing. |

## Hint

`m := make(map[string]int); for _, x := range xs { m[x]++ }`.

## Validate

```bash
make verify
```
