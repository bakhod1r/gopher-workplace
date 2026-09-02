# Take While

**Level:** middle  
**Topic:** 03-generics

## Context

A parser consumes the leading digits of a token and stops at the first character that is not one.

## Task

Implement the stub(s) in [takewhilegen.go](takewhilegen.go):

1. Implement `TakeWhile`, returning the leading elements that satisfy `pred`.
2. Stop at the first rejected element — later matches are not collected.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TakeWhile([]int{2,4,5,6}, isEven)
Output: []int{2,4}
```

**Example 2:**

```
Input:  TakeWhile([]int{1,2}, isEven)
Output: []int{}
```

**Example 3:**

```
Input:  TakeWhile([]int{2,4}, isEven)
Output: []int{2,4}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Prefix versus filter** | `Filter` keeps every match; `TakeWhile` stops at the first failure. |
| 2 | **`break` in a range loop** | Reused from language basics: leaving the loop early is the whole semantic. |
| 3 | **Higher-order generic functions** | A type parameter may appear in a function-typed parameter or return value. |

## Hint

`break` on the first rejection — do not `continue`.

## Validate

```bash
make verify
```
