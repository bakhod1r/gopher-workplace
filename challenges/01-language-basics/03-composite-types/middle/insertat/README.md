# Insert at Index

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Inserting into an ordered slice: split at the index and stitch the value in.

## Task

Implement `InsertAt(xs, i, v)`, clamping `i` to `[0, len]`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3], i=1, v=9
Output: [1,9,2,3]
```

**Example 2:**

```
Input:  [1,2,3], i=10, v=9
Output: [1,2,3,9]
```

_Explanation:_ i clamps to len

**Example 3:**

```
Input:  [], i=0, v=9
Output: [9]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Split & stitch** | `append(head, v)` then tail. |
| 2 | **Aliasing hazard** | `append(xs[:i], ...)` can clobber the tail. |
| 3 | **Clamp index** | i in `[0, len]`. |

## Hint

Safest: `out := append([]int{}, xs[:i]...); out = append(out, v); out = append(out, xs[i:]...)`.

## Validate

```bash
make verify
```
