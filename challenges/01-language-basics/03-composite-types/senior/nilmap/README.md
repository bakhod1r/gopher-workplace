# Nil Map Write

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`Count` declares the map with `var m map[int]int`, which is **nil**. Reading a
nil map is fine, but the first `m[x]++` (a write) panics.

## Task

Fix the declaration between the markers in [nilmap.go](nilmap.go) to allocate the
map.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,1,2]
Output: {1:2, 2:1}
```

**Example 2:**

```
Input:  []
Output: {}
```

**Example 3:**

```
Input:  [5,5,5]
Output: {5:3}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil map** | `var m map[K]V` is nil. |
| 2 | **Read vs write** | Reading nil is ok; writing panics. |
| 3 | **make** | `make(map[int]int)` allocates. |

## Hint

`m := make(map[int]int)`.

## Validate

```bash
make verify
```
