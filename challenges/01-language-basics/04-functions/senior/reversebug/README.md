# Two-Pointer Reverse Bound

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

The last valid index of a slice is `len(xs)-1`. Initialising the right pointer
to `len(xs)` makes the first `xs[j]` read one past the end, panicking with an
index-out-of-range error.

## Task

Fix the right-pointer initialisation in [reversebug.go](reversebug.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Reverse([1 2 3 4])
Output: [4 3 2 1]
```

**Example 2:**

```
Input:  Reverse([1 2 3])
Output: [3 2 1]
```

**Example 3:**

```
Input:  Reverse([1])
Output: [1]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Last index** | The final element is at `len(xs)-1`. |
| 2 | **Two-pointer technique** | Swap ends, move inward. |
| 3 | **Index bounds** | `xs[len(xs)]` is out of range. |

## Hint

Start the right pointer at the last index: `i, j := 0, len(xs)-1`.

## Validate

```bash
make verify
```
