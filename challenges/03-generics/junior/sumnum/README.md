# Sum Numbers

**Level:** junior  
**Topic:** 03-generics

## Context

An invoice line adds up amounts that are sometimes integer cents and sometimes floating-point rates. One function should cover both.

## Task

Implement the stub(s) in [sumnum.go](sumnum.go):

1. Implement `Sum`, returning the total of all elements.
2. Return the zero value for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Sum([]int{1, 2, 3})
Output: 6
```

**Example 2:**

```
Input:  Sum([]float64{0.5, 0.5})
Output: 1
```

**Example 3:**

```
Input:  Sum([]int{})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Constraints permit operations** | A type parameter can only do what every type in its set can do — that is why `+` needs a numeric constraint. |
| 2 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |
| 3 | **Zero value of `T`** | `var zero T` names the zero value of an unknown type. |

## Hint

`var total T` starts the accumulator; `+=` is allowed because every type in `Number` supports it.

## Validate

```bash
make verify
```
