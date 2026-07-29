# Frequency Counter

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Counting occurrences relies on the map zero value: a missing key reads as 0, so
`m[k]++` just works.

## Task

Implement `Count(xs)` returning element→count.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ["a","b","a","c","a","b"]
Output: {"a":3,"b":2,"c":1}
```

**Example 2:**

```
Input:  nil
Output: {} (empty map)
```

**Example 3:**

```
Input:  ["x"]
Output: {"x":1}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
