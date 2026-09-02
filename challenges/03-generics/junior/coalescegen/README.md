# Coalesce

**Level:** junior  
**Topic:** 03-generics

## Context

Configuration is layered: command-line flag, then environment, then a built-in default. The first value that was actually set wins.

## Task

Implement the stub(s) in [coalescegen.go](coalescegen.go):

1. Implement `Coalesce`, returning the first argument that is not the zero value of `T`.
2. Return the zero value when every argument is zero or no arguments are given.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Coalesce(0, 0, 5)
Output: 5
```

**Example 2:**

```
Input:  Coalesce("", "a", "b")
Output: "a"
```

**Example 3:**

```
Input:  Coalesce(0, 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variadic type parameters** | `vals ...T` collects any number of arguments of the same type. |
| 2 | **Zero value of `T`** | `var zero T` is the only way to name the zero value of an unknown type. |
| 3 | **The `comparable` constraint** | `comparable` is what lets you use `==` and `!=` on a type parameter. |

## Hint

Same `var zero T` trick as `IsZero`, inside a loop.

## Validate

```bash
make verify
```
