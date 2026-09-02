# Curry

**Level:** middle  
**Topic:** 03-generics

## Context

A middleware helper wants `withPrefix("api")` to hand back a ready-made function it can pass around.

## Task

Implement the stub(s) in [currygen.go](currygen.go):

1. Implement `Curry2`, returning a function that takes `A` and returns a function taking `B`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Curry2(add)(1)(2)
Output: 3
```

**Example 2:**

```
Input:  Curry2(repeat)("ab")(2)
Output: "abab"
```

**Example 3:**

```
Input:  plus1 := Curry2(add)(1); plus1(5)
Output: 6
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closures capture instantiated types** | Inside a returned closure, `T` is already fixed by the outer call. |
| 2 | **Partial application** | Fixing the first argument produces a reusable specialised function. |
| 3 | **Higher-order generic functions** | A type parameter may appear in a function-typed parameter or return value. |

## Hint

Two nested `func` literals; the inner one closes over `a`.

## Validate

```bash
make verify
```
