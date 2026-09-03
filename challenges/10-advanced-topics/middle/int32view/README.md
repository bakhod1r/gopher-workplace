# Reinterpret Bytes As Wider Values

**Level:** middle
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A numeric kernel receives its input as bytes off the wire and wants to run over it as `int32` without copying eight megabytes first.

## Task

Implement [int32view.go](int32view.go):

1. Return an `[]int32` view sharing `b`'s storage.
2. Report false for an empty slice, a length that is not a multiple of four, or a misaligned start.
3. Zero allocations.

Replace the stub body in [int32view.go](int32view.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Int32s(alignedBytes(8))
Output: a 2-element view, true
```

**Example 2:**

```
Input:  Int32s(b[:6])
Output: nil, false
```

_Explanation:_ 6 is not a multiple of 4.

**Example 3:**

```
Input:  Int32s(b[1:13])
Output: nil, false
```

_Explanation:_ Misaligned start.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Slice** | Reinterprets a typed pointer and a length as a slice. |
| 2 | **Element count, not byte count** | The length argument is in elements — `len(b)/4`. |
| 3 | **Two preconditions** | Length divisibility and address alignment are both required. |

## Hint

`unsafe.Slice((*int32)(p), n)` — think carefully about what `n` is.

## Validate

```bash
make verify
```
