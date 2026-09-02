# Pipeline

**Level:** middle  
**Topic:** 03-generics

## Context

A text normaliser trims, lowercases, and collapses whitespace. The steps are configured per tenant.

## Task

Implement the stub(s) in [pipelinegen.go](pipelinegen.go):

1. Implement `Pipeline`, returning a function that applies every stage left to right.
2. With no stages the returned function is the identity.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Pipeline(double, inc)(3)
Output: 7
```

**Example 2:**

```
Input:  Pipeline[int]()(3)
Output: 3
```

**Example 3:**

```
Input:  Pipeline(inc, double)(3)
Output: 8
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variadic higher-order functions** | `...func(T) T` composes any number of same-typed stages. |
| 2 | **Identity as the empty case** | Zero stages must return the input untouched. |
| 3 | **Closures capture instantiated types** | Inside a returned closure, `T` is already fixed by the outer call. |

## Hint

Reassign `v` inside the loop; the empty case falls out for free.

## Validate

```bash
make verify
```
