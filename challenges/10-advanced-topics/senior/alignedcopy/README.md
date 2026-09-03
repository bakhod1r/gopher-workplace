# Copy Words Only When The Buffer Allows

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A decoder copies a received byte buffer into a word array with a loop and four shifts per word. On a same-endian link the shifts are pure overhead.

## Task

Implement [alignedcopy.go](alignedcopy.go):

1. Copy as many whole `uint32` values as fit from `src` into `dst`.
2. Return the count, and false when `src`'s length is not a multiple of four or its start is misaligned.
3. `dst` must not alias `src`; allocate nothing.

Replace the stub body in [alignedcopy.go](alignedcopy.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  CopyWords(make([]uint32,3), bytesOf(1,2,3))
Output: 3, true
```

**Example 2:**

```
Input:  a 2-element dst
Output: 2, true
```

_Explanation:_ `copy` stops at the shorter side.

**Example 3:**

```
Input:  src[1:13]
Output: 0, false
```

_Explanation:_ Misaligned start.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reinterpret then copy** | The view costs nothing; the copy is what makes `dst` independent. |
| 2 | **copy bounds both sides** | It moves `min(len(dst), len(view))` elements. |
| 3 | **Two preconditions** | Length divisibility and address alignment. |
| 4 | **Not a wire format** | The interpretation is the host's byte order. |

## Hint

Build the view, then let `copy` do the bounding.

## Validate

```bash
make verify
```
