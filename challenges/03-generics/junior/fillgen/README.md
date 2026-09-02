# Fill

**Level:** junior  
**Topic:** 03-generics

## Context

A grid renderer needs a row pre-populated with a default cell before the real data arrives.

## Task

Implement the stub(s) in [fillgen.go](fillgen.go):

1. Implement `Fill`, returning a slice containing `n` copies of `v`.
2. Return an empty (non-nil) slice when `n` is zero or negative.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Fill(3, 7)
Output: []int{7, 7, 7}
```

**Example 2:**

```
Input:  Fill(2, "x")
Output: []string{"x", "x"}
```

**Example 3:**

```
Input:  Fill(0, 7)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`make([]T, ...)`** | The element type can be a type parameter — the compiler knows its size per instantiation. |
| 2 | **Guarding negative counts** | Reused from language basics: `make` panics on a negative length. |
| 3 | **Type parameters** | `[T any]` declares a type parameter; the caller (or inference) picks `T`. |

## Hint

A plain counting `for` loop is enough; the negative case falls out of the loop condition.

## Validate

```bash
make verify
```
