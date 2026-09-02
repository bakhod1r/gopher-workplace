# The []any Trap

**Level:** junior  
**Topic:** 03-generics

## Context

A legacy API takes `[]any`. Someone tried passing a `[]int` and the compiler refused; the reason is worth knowing.

## Task

Implement the stub(s) in [anyslicepitfall.go](anyslicepitfall.go):

1. Implement `ToAny`, converting a typed slice element by element.
2. `SumInts` is provided — note it needs no conversion at all.
3. Understand why `[]int` cannot be used directly as `[]any`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ToAny([]int{1, 2})
Output: []any{1, 2}
```

**Example 2:**

```
Input:  len(ToAny([]string{"a"}))
Output: 1
```

**Example 3:**

```
Input:  ToAny([]int{})
Output: []any{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slice types are invariant** | `[]int` and `[]any` have different memory layouts, so no conversion exists. |
| 2 | **Boxing costs** | Each element becomes an interface value — a copy plus, potentially, an allocation. |
| 3 | **Generics avoid the problem** | A type parameter keeps the elements as they are, which is why `SumInts` needs no conversion. |

## Hint

You must loop: there is no cast from `[]T` to `[]any`.

## Validate

```bash
make verify
```
