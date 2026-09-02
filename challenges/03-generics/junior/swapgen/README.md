# Generic Swap

**Level:** junior  
**Topic:** 03-generics

## Context

The team keeps writing three-line swaps with a temporary variable for every new type. One generic helper replaces them all.

## Task

Implement the stub(s) in [swapgen.go](swapgen.go):

1. Implement `Swap`, returning `b` then `a`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Swap(1, 2)
Output: 2, 1
```

**Example 2:**

```
Input:  Swap("a", "b")
Output: "b", "a"
```

**Example 3:**

```
Input:  Swap(true, false)
Output: false, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type parameters** | `[T any]` declares a type parameter; the caller (or inference) picks `T`. |
| 2 | **Type inference** | You call `F(x)`, not `F[int](x)` — the compiler reads `T` from the argument. |
| 3 | **Multiple return values** | Reused from language basics: Go returns tuples directly, no temp needed. |

## Hint

Both parameters share one type parameter, so both arguments must be the same type.

## Validate

```bash
make verify
```
