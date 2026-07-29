# Two Appends Share Spare Cap

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Because `base` has spare capacity, both appends write into the same backing
array at the same slot; the second clobbers the first. Clip the base so each
append allocates its own array.

## Task

Fix [twoappend.go](twoappend.go) so the first variant keeps its value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Fork([7](cap 4), 100, 200)
Output: x == [7 100], last == 100
```

**Example 2:**

```
Input:  x unaffected by second append
Output: true
```

**Example 3:**

```
Input:  x[1]
Output: 100 (not 200)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Spare-capacity aliasing** | Both appends reuse base's array. |
| 2 | **Full-slice expression** | `base[:len:len]` forces reallocation. |
| 3 | **Independent buffers** | Each fork must own its memory. |

## Hint

Clip before appending: `bc := base[:len(base):len(base)]`, then append `a` and `b` to `bc` separately.

## Validate

```bash
make verify
```
