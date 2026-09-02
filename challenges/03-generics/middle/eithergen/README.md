# Either

**Level:** middle  
**Topic:** 03-generics

## Context

A parser step yields either a parsed record or the raw line it could not handle, and both need to travel together through the pipeline.

## Task

Implement the stub(s) in [eithergen.go](eithergen.go):

1. Implement `Left`, `Right`, and `Unwrap`.
2. `Unwrap` reports which side is set; the unset side is the zero value.
3. Both constructors need explicit type arguments — understand why.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Left[string, int]("e").Unwrap()
Output: "e", 0, true
```

**Example 2:**

```
Input:  Right[string, int](5).Unwrap()
Output: "", 5, false
```

**Example 3:**

```
Input:  zero Either
Output: zero, zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two independent parameters** | Neither side constrains the other. |
| 2 | **Inference cannot help** | `Left(v)` fixes only `L`, so `R` must be written out. |
| 3 | **Sum types in Go** | Go has no built-in union, so a struct with a discriminator is the idiom. |

## Hint

`Left("e")` cannot infer `R` — the call site must name both parameters.

## Validate

```bash
make verify
```
