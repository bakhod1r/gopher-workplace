# Compose

**Level:** middle  
**Topic:** 03-generics

## Context

A decoding pipeline parses a string then validates the result. Both steps are small functions the caller wants to bolt together once.

## Task

Implement the stub(s) in [composegen.go](composegen.go):

1. Implement `Compose`, returning a function that applies `f` first and `g` to its result.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Compose(double, itoa)(2)
Output: "4"
```

**Example 2:**

```
Input:  Compose(itoa, length)(12)
Output: 2
```

**Example 3:**

```
Input:  Compose(double, double)(3)
Output: 12
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Higher-order generic functions** | A type parameter may appear in a function-typed parameter or return value. |
| 2 | **Three type parameters** | `A` in, `B` between, `C` out — the middle type never appears in the result. |
| 3 | **Closures capture instantiated types** | Inside a returned closure, `T` is already fixed by the outer call. |

## Hint

Return a closure; the body is one nested call.

## Validate

```bash
make verify
```
