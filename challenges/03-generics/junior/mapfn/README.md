# Map Over a Slice

**Level:** junior  
**Topic:** 03-generics

## Context

A report layer turns records into display strings. The transformation differs per screen, but the looping never does.

## Task

Implement the stub(s) in [mapfn.go](mapfn.go):

1. Implement `Map`, applying `f` to every element of `s` in order.
2. Return a slice of the results, one per input element.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Map([]int{1, 2}, double)
Output: []int{2, 4}
```

**Example 2:**

```
Input:  Map([]int{1, 2}, itoa)
Output: []string{"1", "2"}
```

**Example 3:**

```
Input:  Map([]int{}, double)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two type parameters** | `[T, U any]` lets the input and output element types differ. |
| 2 | **Functions as values** | Reused from language basics: a `func(T) U` parameter is an ordinary value. |
| 3 | **Type inference** | You call `F(x)`, not `F[int](x)` — the compiler reads `T` from the argument. |

## Hint

The output slice is `[]U`, not `[]T` — that is the whole point of the second type parameter.

## Validate

```bash
make verify
```
