# Array Key Coordinate Order

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Arrays are comparable, so `[2]int` is a valid map key. But the key is built with
the coordinates **swapped** (`{c[1], c[0]}`), so `(1,2)` and `(2,1)` are conflated.

## Task

Fix the key between the markers in [maparraykey.go](maparraykey.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  cells=[[0,1],[0,1],[2,3]]
Output: { [0,1]:2, [2,3]:1 }
```

**Example 2:**

```
Input:  cells=[[1,0]]
Output: { [1,0]:1 }
```

**Example 3:**

```
Input:  cells=[[2,3],[3,2]]
Output: { [2,3]:1, [3,2]:1 }
```

_Explanation:_ row/col order preserved.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Arrays are comparable** | Usable as map keys. |
| 2 | **Slices are not** | `[]int` can't be a key. |
| 3 | **Key identity** | Field order defines the key. |

## Hint

`m[[2]int{c[0], c[1]}]++`.

## Validate

```bash
make verify
```
