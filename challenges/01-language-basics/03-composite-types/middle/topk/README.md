# Top K Frequent

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

"Top N" analytics: count occurrences, then rank.

## Task

Implement `TopK(xs, k)` — k most frequent, ties broken alphabetically.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ["a","b","a","c","b","a","d"], k=2
Output: ["a","b"]
```

_Explanation:_ a=3,b=2

**Example 2:**

```
Input:  same, k=4
Output: ["a","b","c","d"]
```

_Explanation:_ c,d tie count1 -> alphabetical

**Example 3:**

```
Input:  ["z","z"], k=5
Output: ["z"]
```

_Explanation:_ k clamped to distinct count

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Count map** | `m[x]++`. |
| 2 | **Custom sort** | `sort.Slice` with count then name. |
| 3 | **Clamp k** | Not more than distinct count. |

## Hint

Count into a map, collect keys, `sort.Slice` by `(count desc, key asc)`, take
first k.

## Validate

```bash
make verify
```
