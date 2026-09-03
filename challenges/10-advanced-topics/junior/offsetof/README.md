# Where Does This Field Start

**Level:** junior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A binary protocol is written by copying a struct's bytes. It works on one machine and produces garbage on another, because nobody checked where the fields actually sit.

## Task

Implement [offsetof.go](offsetof.go):

1. Return the byte offset of each field of `Rec`.
2. Derive them with `unsafe.Offsetof`, not by adding up sizes.

Replace the stub body in [offsetof.go](offsetof.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Offsets()
Output: 0, 8, 16
```

_Explanation:_ The one-byte `A` is followed by seven bytes of padding.

**Example 2:**

```
Input:  first offset
Output: 0
```

_Explanation:_ The first field always starts at the struct's address.

**Example 3:**

```
Input:  B - A
Output: 8, not 1
```

_Explanation:_ `int64` must be 8-byte aligned.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Offsetof** | The field's byte offset within its struct, as a compile-time constant. |
| 2 | **Alignment** | A field's offset is rounded up to a multiple of its alignment. |
| 3 | **Padding** | The gap alignment leaves behind is real memory. |

## Hint

`Offsetof` takes a field selector, so you need a variable to select from.

## Validate

```bash
make verify
```
