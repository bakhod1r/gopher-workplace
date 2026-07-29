# Two Sum

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Find two elements adding to a target — the canonical map-as-index-lookup.

## Task

Implement `TwoSum(xs, target)` returning indices `i<j`, or ok=false.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [2,7,11,15], target=9
Output: (0,1,true)
```

_Explanation:_ 2+7=9

**Example 2:**

```
Input:  [3,2,4], target=6
Output: (1,2,true)
```

_Explanation:_ 2+4=6

**Example 3:**

```
Input:  [1,2], target=100
Output: (0,0,false)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Complement lookup** | Need `target - xs[i]`. |
| 2 | **Map value→index** | Remember where each value was. |
| 3 | **Single pass** | Check before inserting. |

## Hint

For each i, if `seen[target-xs[i]]` exists return it; else `seen[xs[i]] = i`.

## Validate

```bash
make verify
```
