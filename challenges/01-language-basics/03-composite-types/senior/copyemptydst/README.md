# copy Into a Length-0 Slice

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`copy` copies `min(len(dst), len(src))` elements. The destination is made with
**length 0** (only capacity), so `copy` copies nothing and `Clone` returns empty.

## Task

Fix the `make` between the markers in [copyemptydst.go](copyemptydst.go) to give
the destination the right length.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3]
Output: [1,2,3]
```

**Example 2:**

```
Input:  []
Output: []
```

**Example 3:**

```
Input:  [9]
Output: [9]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **copy semantics** | Copies `min(len(dst),len(src))`. |
| 2 | **len vs cap** | Capacity alone doesn't hold elements. |
| 3 | **Size dst** | `make([]int, len(xs))`. |

## Hint

`make([]int, len(xs))` (length, not just capacity).

## Validate

```bash
make verify
```
