# Length vs Capacity for Indexing

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`make([]int, 0, len(xs))` has capacity but **length 0**, so `out[i] = ...` panics
(index out of range). Either index into a full-length slice, or `append`.

## Task

Fix the build between the markers in
[preallocindex.go](preallocindex.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3]
Output: [2,4,6]
```

**Example 2:**

```
Input:  []
Output: []
```

**Example 3:**

```
Input:  [-1,4]
Output: [-2,8]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **len vs cap** | You can only index within length. |
| 2 | **Two idioms** | `make(_, len)`+index, or `make(_,0,cap)`+append. |
| 3 | **Panic** | Indexing past length panics. |

## Hint

Either `out := make([]int, len(xs))` and index, or keep cap and `append`.

## Validate

```bash
make verify
```
