# Grow A Slice Whose Type You Do Not Know

**Level:** middle
**Topic:** 10-advanced-topics / 03-reflection

## Context

A fixture helper pads every slice in a test table to the same length. Writing it once per element type produced six near-identical functions that drifted apart.

## Task

Implement [appendzero.go](appendzero.go):

1. Append `n` zero values to the slice `slicePtr` points at.
2. Work for any element type, including structs and a nil slice.
3. Return `ErrTarget` for a non-pointer, a nil pointer, a non-slice, or a negative `n`.

Replace the stub body in [appendzero.go](appendzero.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s := []int{1}; AppendZero(&s, 2)
Output: s is [1 0 0]
```

**Example 2:**

```
Input:  s := []string{"a"}; AppendZero(&s, 1)
Output: s is ["a" ""]
```

_Explanation:_ The zero value follows the element type.

**Example 3:**

```
Input:  AppendZero(&s, 0)
Output: s unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **reflect.Append** | The reflective twin of `append`; it returns a new Value that must be stored back. |
| 2 | **reflect.Zero** | Produces the zero Value for any type. |
| 3 | **Type.Elem()** | A slice type knows its element type. |
| 4 | **Set writes the result back** | `Append` alone does not modify the caller's slice. |

## Hint

`reflect.Append` returns a new slice Value. Something has to happen to it.

## Validate

```bash
make verify
```
