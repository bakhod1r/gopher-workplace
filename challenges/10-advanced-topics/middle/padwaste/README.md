# How Many Bytes Is The Padding

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A team is told to reorder struct fields for memory, and nobody can say which structs are actually wasteful. A number would settle it in a minute.

## Task

Implement [padwaste.go](padwaste.go):

1. Return the struct's size minus the sum of its field sizes.
2. Include unexported fields — they occupy space too.
3. Return 0 for a non-struct or a nil interface.

Replace the stub body in [padwaste.go](padwaste.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Waste(gappy{})
Output: 14
```

_Explanation:_ byte, int64, byte on a 64-bit build.

**Example 2:**

```
Input:  Waste(packed{})
Output: less than gappy
```

_Explanation:_ Widest field first.

**Example 3:**

```
Input:  Waste(none{})
Output: 0
```

_Explanation:_ Two int64s need no padding.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type.Size** | The reflective twin of `unsafe.Sizeof`, available for a type known at run time. |
| 2 | **Padding is size minus content** | Internal gaps and tail padding both show up in this difference. |
| 3 | **Unexported fields count** | They cannot be read, but they still occupy bytes. |

## Hint

Sum the field sizes, subtract from the struct's size.

## Validate

```bash
make verify
```
