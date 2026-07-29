# Append Clobbers via Shared Capacity

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`xs[:2]` shares `xs`'s backing array **and its spare capacity**. Appending to it
writes into `xs[2]`, corrupting the source. A full-slice expression caps the
sub-slice so append must reallocate.

## Task

Fix the sub-slice between the markers in
[appendcapshared.go](appendcapshared.go) to not share spare capacity.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=[1,2,3], extra=9
Output: [1,2,9]
```

_Explanation:_ head is capped to len 2, so append allocates a new array; xs stays [1,2,3].

**Example 2:**

```
Input:  xs=[5,6,7,8], extra=0
Output: [5,6,0]
```

_Explanation:_ xs[2] must remain 7 afterwards.

**Example 3:**

```
Input:  xs=[1,2], extra=3
Output: [1,2,3]
```

_Explanation:_ No spare capacity so append reallocates anyway.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Sub-slice capacity** | `xs[:2]` keeps `xs`'s capacity. |
| 2 | **Full-slice expr** | `xs[:2:2]` caps capacity to 2. |
| 3 | **Force realloc** | Append then can't touch xs. |

## Hint

`head := xs[:2:2]` (three-index slice caps capacity).

## Validate

```bash
make verify
```
