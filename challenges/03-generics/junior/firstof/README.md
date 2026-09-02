# First Element

**Level:** junior  
**Topic:** 03-generics

## Context

A queue reader needs the head element, but callers must be able to tell an empty queue from a real zero value.

## Task

Implement the stub(s) in [firstof.go](firstof.go):

1. Implement `First`, returning the first element and `true`.
2. For an empty slice return the zero value of `T` and `false`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  First([]int{3, 1})
Output: 3, true
```

**Example 2:**

```
Input:  First([]string{"a"})
Output: "a", true
```

**Example 3:**

```
Input:  First([]int{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Zero value of `T`** | `var zero T` is the only way to name the zero value of an unknown type. |
| 2 | **Slices of a type parameter** | `[]T` behaves like any slice: `len`, `range`, `append` all work. |
| 3 | **Comma-ok returns** | Reused from language basics: a second `bool` distinguishes "missing" from "zero". |

## Hint

You cannot write `return 0` — `T` may not be numeric. Declare `var zero T`.

## Validate

```bash
make verify
```
