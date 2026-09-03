# Is This Address Aligned

**Level:** middle
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A decoder reinterprets an incoming buffer as a slice of `uint64`. On x86 it is merely slow when the buffer is misaligned; on other architectures it faults.

## Task

Implement [alignment.go](alignment.go):

1. Report whether `b`'s first byte is at an address that is a multiple of `n`.
2. Return false for an empty slice, for `n == 0`, and for any `n` that is not a power of two.

Replace the stub body in [alignment.go](alignment.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Aligned(bufFromUint64s, 8)
Output: true
```

**Example 2:**

```
Input:  Aligned(buf[1:], 8)
Output: false
```

_Explanation:_ One byte in is no longer aligned.

**Example 3:**

```
Input:  Aligned(buf, 3)
Output: false
```

_Explanation:_ 3 is not a power of two.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer to uintptr** | The numeric address is only meaningful for arithmetic like this. |
| 2 | **Power-of-two masking** | `x & (n-1) == 0` is the alignment test when n is a power of two. |
| 3 | **Why alignment matters** | Wide loads on a misaligned address are slow or illegal depending on the machine. |

## Hint

A power of two has exactly one bit set: `n & (n-1)` is zero.

## Validate

```bash
make verify
```
