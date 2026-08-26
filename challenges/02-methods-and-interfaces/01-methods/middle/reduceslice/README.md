# Reduce Slice

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Finishing the map/filter/reduce trio. `Product` reduces a list to a single value.

## Task

Implement `Product` on `IntList` in [reduceslice.go](reduceslice.go):

1. Multiply all elements together.
2. The product of an empty list is 1 (the multiplicative identity).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IntList{2, 3, 4}.Product()
Output: 24
```

**Example 2:**

```
Input:  IntList{5, 0}.Product()
Output: 0
```

**Example 3:**

```
Input:  IntList{}.Product()
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reduce operation** | Accumulating state in a loop. |
| 2 | **Multiplicative identity** | Start accumulation at 1, not 0. |

## Hint

`p := 1`, loop `p *= v`, return `p`.

## Validate

```bash
make verify
```
