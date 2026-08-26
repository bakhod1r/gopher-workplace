# Apply Slice

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A math library adds methods directly to slices. You can mutate the elements of a
slice even with a value receiver, because slices are references to a backing array.

## Task

Implement `Apply` on `IntList` in [applyslice.go](applyslice.go):

1. Multiply each element by `factor`.
2. Do it in place (mutate the elements).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  l := IntList{1, 2, 3}; l.Apply(2)
Output: l == {2, 4, 6}
```

**Example 2:**

```
Input:  l := IntList{5}; l.Apply(0)
Output: l == {0}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods on slice types** | `type IntList []int` can have methods. |
| 2 | **Slice value receiver** | You can mutate *elements* of a slice with a value receiver, because it copies the slice header, not the backing array. |

## Hint

`for i := range l { l[i] *= factor }`. Note: no pointer receiver needed to
change elements.

## Validate

```bash
make verify
```
