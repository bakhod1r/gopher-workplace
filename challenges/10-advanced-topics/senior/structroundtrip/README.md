# A Struct To Bytes And Back

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A framing layer decodes fixed-size records field by field with `binary.LittleEndian`. Both ends are the same architecture, and the field-by-field decode is most of the receive path.

## Task

Implement [structroundtrip.go](structroundtrip.go):

1. Reinterpret `b` as a `Frame` and return a copy of it.
2. Report false unless `len(b)` is exactly the frame's size and the start is correctly aligned.
3. The result must not alias `b`; allocate nothing.

Replace the stub body in [structroundtrip.go](structroundtrip.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Decode(Encode(&f))
Output: f, true
```

**Example 2:**

```
Input:  Decode(b[:4])
Output: zero Frame, false
```

_Explanation:_ Wrong length.

**Example 3:**

```
Input:  a misaligned slice
Output: zero Frame, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Dereference copies** | `*(*Frame)(p)` reads the struct out by value. |
| 2 | **Size and alignment from the type** | `Sizeof` and `Alignof`, never literals. |
| 3 | **Exact length** | Too short reads past the end; too long is a framing error. |
| 4 | **Not portable** | The layout, padding and byte order are the local machine's. |

## Hint

Two guards from the type, then one dereference.

## Validate

```bash
make verify
```
