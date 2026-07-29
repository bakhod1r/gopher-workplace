# Append Invalidates Alias

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

When append exceeds capacity it allocates a new array; pointers taken before
the append still point at the OLD array. Writing through the stale pointer does
not change the current slice.

## Task

Fix [stalecap.go](stalecap.go) so the write updates the current slice's element 0.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstAfterGrow(7)
Output: 99
```

**Example 2:**

```
Input:  result reflects write through fresh pointer
Output: 99
```

**Example 3:**

```
Input:  element after grow
Output: 99
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Append reallocation** | Exceeding cap copies to a new array. |
| 2 | **Stale element pointers** | Old `&xs[0]` no longer aliases the slice. |
| 3 | **Re-address after growth** | Take the address again or index directly. |

## Hint

Write to the current slice instead: `xs[0] = 99` (or re-take `p = &xs[0]` after the append).

## Validate

```bash
make verify
```
