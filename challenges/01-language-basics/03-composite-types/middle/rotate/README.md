# Rotate a Slice

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

A round-robin scheduler rotates the queue. Rotate left by k, wrapping, for any k.

## Task

Implement `Left(xs, k)` returning a new rotated slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3,4,5], k=2
Output: [3,4,5,1,2]
```

**Example 2:**

```
Input:  [1,2,3], k=4
Output: [2,3,1]
```

_Explanation:_ k mod len

**Example 3:**

```
Input:  [1,2,3], k=-1
Output: [3,1,2]
```

_Explanation:_ negative normalized

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Normalize k** | `((k % n) + n) % n`. |
| 2 | **Reassemble** | `xs[k:]` then `xs[:k]`. |
| 3 | **Fresh slice** | Build a new result. |

## Hint

`n := len(xs)`; if n==0 return empty; `k = ((k%n)+n)%n`; `append(append([]int{}, xs[k:]...), xs[:k]...)`.

## Validate

```bash
make verify
```
