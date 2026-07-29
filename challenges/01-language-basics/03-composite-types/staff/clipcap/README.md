# Clip Capacity

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Handing out a slice with spare capacity is dangerous: the receiver's `append` can
overwrite memory you still use. `Clip` should cap capacity to length via a
full-slice expression. The code returns `xs` unchanged.

## Task

Fix the return between the markers in [clipcap.go](clipcap.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=make([]int,3,10)
Output: len 3, cap 3
```

**Example 2:**

```
Input:  xs len 3
Output: len stays 3
```

**Example 3:**

```
Input:  cap 10 clipped
Output: cap 3
```

_Explanation:_ three-index slice caps future append growth.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Three-index slice** | `xs[:len:len]` caps capacity. |
| 2 | **Append safety** | No spare cap → append reallocates. |
| 3 | **Defensive slicing** | Protect shared backing arrays. |

## Hint

`return xs[:len(xs):len(xs)]`.

## Validate

```bash
make verify
```
