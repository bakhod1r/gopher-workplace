# The Length Argument Is In Elements

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A frame decoder reinterprets a received buffer as words. It reads plausible garbage past the end of every buffer, and the corruption is blamed on the network for a week.

## Task

Fix the single planted bug in [slicelen.go](slicelen.go):

1. Return a `[]uint32` view sharing `b`'s storage.
2. Report false for an empty buffer, a length that is not a multiple of four, or a misaligned start.
3. Fix the single bug: the view must cover exactly `b` and no more.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Words(alignedBytes(8))
Output: a 2-element view, true
```

**Example 2:**

```
Input:  len(view) * 4
Output: len(b)
```

_Explanation:_ The view covers exactly the buffer.

**Example 3:**

```
Input:  Words(b[:6])
Output: nil, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Slice's length is in elements** | It is multiplied by the element size internally. |
| 2 | **Out-of-bounds without a panic** | The runtime cannot check a length you invented. |
| 3 | **Silent corruption** | The extra elements read whatever follows the buffer in memory. |

## Hint

Everything about the pointer is right. Count what the second argument is counting.

## Validate

```bash
make verify
```
