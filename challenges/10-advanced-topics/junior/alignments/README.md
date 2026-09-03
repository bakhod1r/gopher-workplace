# What Each Type Must Line Up On

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A struct is laid out by hand for a binary format. Guessing where each field lands works until the first field wider than a byte.

## Task

Implement [alignments.go](alignments.go):

1. Return the alignment of `byte`, `int32`, `int64` and `string`.
2. Derive them with `unsafe.Alignof` rather than writing numbers.

Replace the stub body in [alignments.go](alignments.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Alignments()
Output: 1, 4, 8, 8
```

_Explanation:_ On a 64-bit build.

**Example 2:**

```
Input:  byte's alignment
Output: 1
```

_Explanation:_ A byte can sit anywhere.

**Example 3:**

```
Input:  every result
Output: a power of two
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Alignof** | The boundary a type's address must be a multiple of. |
| 2 | **Alignment drives padding** | A field's offset is rounded up to its alignment. |
| 3 | **Composite alignment** | A struct's alignment is its widest field's. |

## Hint

`Alignof` takes a value, so declare one of each.

## Validate

```bash
make verify
```
