# Coordinates That Collide

**Level:** staff  
**Topic:** 03-generics

## Context

A sparse grid stores one value per occupied cell. A 200x200 import of forty thousand cells ends up holding a few hundred, and reading back (1,2) returns whatever was written to (2,1).

## Task

Fix the single planted bug in [sparsegridkeybug.go](sparsegridkeybug.go):

1. Find and fix the single bug so every distinct point maps to a distinct key.
2. Negative coordinates must work too.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Set({1,2},"a"); Set({2,1},"b"); Get({1,2})
Output: "a"
```

**Example 2:**

```
Input:  Set({-1,5},"a"); Set({4,0},"b"); Len()
Output: 2
```

**Example 3:**

```
Input:  fill a 200x200 block; Len()
Output: 40000
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Injective keys** | A key function that is not one-to-one silently merges rows of your data. |
| 2 | **Bit packing** | Two 32-bit halves fit in one int64: shift the first, mask the second. |
| 3 | **Sign extension** | `int64(y)` for negative y fills the high bits with ones; convert through an unsigned type to keep the halves apart. |

## Hint

How many different points give the same value for `x + y`?

## Validate

```bash
make verify
```
