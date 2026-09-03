# Read A Wide Value Out Of A Byte Buffer

**Level:** middle
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A zero-copy frame reader casts into the middle of a received buffer. On the first misaligned frame it reads garbage on one architecture and crashes on another.

## Task

Implement [uint32at.go](uint32at.go):

1. Read the native-endian uint32 at byte offset `off`.
2. Report false for a negative offset, a read that would run past the end, or a misaligned address.

Replace the stub body in [uint32at.go](uint32at.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Uint32At(buf, 0)
Output: the first word, true
```

**Example 2:**

```
Input:  Uint32At(buf, 5)
Output: 0, false
```

_Explanation:_ The read would run past the end.

**Example 3:**

```
Input:  Uint32At(buf, 2)
Output: 0, false
```

_Explanation:_ Not 4-byte aligned.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounds checks are yours now** | `unsafe` removes the runtime's check, so the code must do it. |
| 2 | **Alignment for wide loads** | A uint32 read wants a 4-byte-aligned address. |
| 3 | **unsafe.Add plus conversion** | Offset first, then reinterpret the pointer. |

## Hint

Three guards before the read: negative offset, past the end, and the low two bits of the address.

## Validate

```bash
make verify
```
