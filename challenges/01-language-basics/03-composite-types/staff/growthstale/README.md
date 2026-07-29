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

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  val=7
Output: first=99, slice=[99 7]
```

**Example 2:**

```
Input:  val=0
Output: first=99, slice=[99 0]
```

**Example 3:**

```
Input:  val=-1
Output: first=99, slice=[99 -1]
```

_Explanation:_ write must target the new backing array, not the stale pointer.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
