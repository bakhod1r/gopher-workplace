# Last N Elements

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Tailing a log or history buffer: return the last N entries, clamping when N
exceeds the length.

## Task

Implement `Last(xs, n)` — last `n` in order; clamp `n` to `[0, len]`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=[1,2,3,4,5], n=2
Output: []int{4,5}
```

**Example 2:**

```
Input:  n=0
Output: []int{}
```

**Example 3:**

```
Input:  n=10
Output: []int{1,2,3,4,5}
```

_Explanation:_ n >= len returns all of xs.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice expression** | `xs[len(xs)-n:]` takes a tail. |
| 2 | **Clamp bounds** | Guard n>len and n<=0. |
| 3 | **Empty result** | Return `[]int{}`, not a panic. |

## Hint

Clamp `n` to `[0, len(xs)]`, then `return xs[len(xs)-n:]`.

## Validate

```bash
make verify
```
