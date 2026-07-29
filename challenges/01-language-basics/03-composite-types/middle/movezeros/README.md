# Move Zeros

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Compacting non-zero values to the front (stable), zeros to the back — a common
array-shuffling task.

## Task

Implement `MoveZeros(xs)` (stable order of non-zeros).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [0,1,0,3,12]
Output: [1,3,12,0,0]
```

**Example 2:**

```
Input:  [1,2,3]
Output: [1,2,3]
```

**Example 3:**

```
Input:  nil
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Stable compaction** | Keep non-zero order. |
| 2 | **Two-phase build** | Non-zeros, then zeros. |
| 3 | **Length preserved** | Output same length. |

## Hint

Append non-zeros, then append `len(xs)-count` zeros.

## Validate

```bash
make verify
```
