# 24-Hour Clock

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Adding hours wraps around midnight. Because Go's `%` can be negative, wrapping
backward needs an extra step to stay in `0..23`.

## Task

Implement `AddHours(h, add)` returning the resulting hour in `0..23`. `add` may
be negative or larger than 24.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AddHours(10, 5)
Output: 15
```

_Explanation:_ No wrap needed.

**Example 2:**

```
Input:  AddHours(23, 1)
Output: 0
```

_Explanation:_ Past midnight wraps to 0.

**Example 3:**

```
Input:  AddHours(0, -1)
Output: 23
```

_Explanation:_ Going back an hour from 0 wraps to 23.

**Example 4:**

```
Input:  AddHours(6, -30)
Output: 0
```

_Explanation:_ 6-30=-24; normalized to 0.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Modulo wrap** | `(h+add) % 24` folds into range. |
| 2 | **Negative remainder** | Go's `%` can be negative; normalize it. |
| 3 | **Normalize** | `((x % 24) + 24) % 24` is always 0..23. |

## Hint

`((h+add)%24 + 24) % 24` guarantees a non-negative result.

## Validate

```bash
make verify
```
