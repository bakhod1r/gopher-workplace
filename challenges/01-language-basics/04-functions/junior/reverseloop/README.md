# Reverse via Index Loop

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

Walking an index from the end builds a reversed copy without touching the input.

## Task

Implement `Reverse` in [reverseloop.go](reverseloop.go) returning a new slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Reverse([1 2 3])
Output: [3 2 1]
```

**Example 2:**

```
Input:  Reverse([1])
Output: [1]
```

**Example 3:**

```
Input:  Reverse(nil)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Descending index loop** | `for i := len(xs)-1; i >= 0; i--`. |
| 2 | **Build new slice** | append to a fresh result. |
| 3 | **Immutability** | Never write into the caller's slice. |

## Hint

Loop from `len(xs)-1` down to 0, appending each element to `out`.

## Validate

```bash
make verify
```
