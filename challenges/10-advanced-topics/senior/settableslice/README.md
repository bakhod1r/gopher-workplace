# Write Into A Slice Through Reflection

**Level:** senior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A normalisation pass written with reflection runs without error and changes nothing. The author adds logging inside the loop, sees the right values computed, and cannot see where they go.

## Task

Fix the single planted bug in [settableslice.go](settableslice.go):

1. Double every element of the int slice, in place.
2. Return `ErrShape` for anything that is not a slice of ints.
3. Fix the single bug so the writes actually reach the caller's slice.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  s := []int{1,2,3}; Double(s)
Output: s is [2 4 6]
```

**Example 2:**

```
Input:  s := []int{1,2,3,4}; Double(s[1:3])
Output: s is [1 4 6 4]
```

_Explanation:_ Only the view is written.

**Example 3:**

```
Input:  Double(3)
Output: ErrShape
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Assignment to a Value does nothing** | `e = reflect.ValueOf(x)` rebinds a local variable, it does not store anything. |
| 2 | **Value.SetInt** | The write path; it requires the Value to be addressable. |
| 3 | **Why slice elements are settable** | The elements live in the backing array, which the header points at. |

## Hint

`e` is a handle, not a slot. Which method writes through a handle?

## Validate

```bash
make verify
```
