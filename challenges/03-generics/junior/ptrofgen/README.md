# Pointer To Value

**Level:** junior  
**Topic:** 03-generics

## Context

A request builder has optional fields typed `*int` and `*string`. Literals have no address, so every call site declares a temporary variable.

## Task

Implement the stub(s) in [ptrofgen.go](ptrofgen.go):

1. Implement `Ptr`, returning a pointer to a copy of `v`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  *Ptr(7)
Output: 7
```

**Example 2:**

```
Input:  *Ptr("go")
Output: "go"
```

**Example 3:**

```
Input:  Ptr(7) != nil
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Addressable parameters** | A parameter is a local variable, so `&v` is legal even though `&7` is not. |
| 2 | **Escape to the heap** | Returning `&v` makes the compiler allocate `v` on the heap — safe, unlike C. |
| 3 | **Type parameters** | `[T any]` declares a type parameter; the caller (or inference) picks `T`. |

## Hint

`&v` on the parameter itself is the entire body.

## Validate

```bash
make verify
```
