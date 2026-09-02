# Zip To Pairs

**Level:** junior  
**Topic:** 03-generics

## Context

Names and scores arrive as parallel slices. Downstream code is far less error-prone once they travel as pairs.

## Task

Implement the stub(s) in [zippairs.go](zippairs.go):

1. Implement `Zip`, pairing elements at matching positions.
2. Stop at the shorter slice; return an empty (non-nil) result when either slice is empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Zip([]int{1, 2}, []string{"a", "b"})
Output: [{1 a} {2 b}]
```

**Example 2:**

```
Input:  Zip([]int{1, 2, 3}, []string{"a"})
Output: [{1 a}]
```

**Example 3:**

```
Input:  Zip([]int{}, []string{"a"})
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A slice of an instantiated generic type** | `[]Pair[A, B]` is an ordinary slice whose element type happens to be generic. |
| 2 | **Minimum length** | Reused from earlier: zipping is only defined where both slices have elements. |
| 3 | **Instantiating a generic type** | `Stack[int]{}` or `NewStack[int]()` fixes `T` at the use site. |

## Hint

The element type of the result is `Pair[A, B]` — write the instantiation out in `make`.

## Validate

```bash
make verify
```
