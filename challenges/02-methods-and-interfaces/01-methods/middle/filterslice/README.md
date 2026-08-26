# Filter Slice

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Continuing the math library, `FilterEvens` returns a subset of the slice.
Because it returns a new slice, it doesn't need to mutate the original.

## Task

Implement `FilterEvens` on `IntList` in [filterslice.go](filterslice.go):

1. Iterate over the elements.
2. Append evens (`v%2 == 0`) to a new `IntList`.
3. Return the new list (empty list if no evens).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IntList{1, 2, 3, 4}.FilterEvens()
Output: {2, 4}
```

**Example 2:**

```
Input:  IntList{1, 3}.FilterEvens()
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Read-only scan, returns new value. |
| 2 | **Slice appending** | `var result IntList; result = append(result, v)` |

## Hint

`var res IntList`; loop and append; return `res`. (If `res` is nil, it works like `{}`).

## Validate

```bash
make verify
```
