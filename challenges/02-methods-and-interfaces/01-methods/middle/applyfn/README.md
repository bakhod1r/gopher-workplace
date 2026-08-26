# Apply All

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A data pipeline transforms records. Each stage is a function — method values
let you plug object methods directly into the pipeline.

## Task

Implement `ApplyAll` in [applyfn.go](applyfn.go):

1. Apply `fn` to each element of `nums`.
2. Return a new slice with the results.
3. Return empty slice (not nil) for empty/nil input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ApplyAll(Transformer{2}.Transform, []int{1, 2, 3})
Output: [2, 4, 6]
```

**Example 2:**

```
Input:  ApplyAll(Transformer{3}.Transform, []int{0, 5})
Output: [0, 15]
```

**Example 3:**

```
Input:  ApplyAll(Transformer{10}.Transform, nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method values as callbacks** | `tr.Transform` is a `func(int) int` — pass it directly. |
| 2 | **Higher-order functions** | `ApplyAll` accepts a function argument. |
| 3 | **Slice mapping** | Build result slice from transformation. |

## Hint

`result := make([]int, 0, len(nums))` then `append(result, fn(n))` in a loop.

## Validate

```bash
make verify
```
