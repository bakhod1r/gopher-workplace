# Clip Spare Capacity

**Level:** middle  
**Topic:** 03-generics

## Context

A handler returns a sub-slice of a large buffer. The caller appends to it and silently overwrites the next record.

## Task

Implement the stub(s) in [slicesclipgen.go](slicesclipgen.go):

1. Implement `Freeze` using `slices.Clip`.
2. After clipping, the capacity must equal the length.
3. Return an empty (non-nil) slice for nil input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  cap after Freeze
Output: == len
```

**Example 2:**

```
Input:  appending to the result
Output: allocates instead of overwriting
```

**Example 3:**

```
Input:  Freeze(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Clip`** | Returns the slice reduced to `s[:len(s):len(s)]`. |
| 2 | **Full slice expressions** | The three-index form is what caps the capacity. |
| 3 | **Append aliasing** | Spare capacity is shared storage — appending into it is a real source of corruption. |

## Hint

Clipping caps the capacity so the next `append` must allocate.

## Validate

```bash
make verify
```
