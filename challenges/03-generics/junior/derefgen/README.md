# Dereference Or Default

**Level:** junior  
**Topic:** 03-generics

## Context

An API response uses pointer fields for optional values. Every read site must handle nil before it can use the value.

## Task

Implement the stub(s) in [derefgen.go](derefgen.go):

1. Implement `Deref`, returning the pointed-to value.
2. Return `def` when `p` is nil.
3. `Ptr` is already written for you — do not change it.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Deref(Ptr(7), 0)
Output: 7
```

**Example 2:**

```
Input:  Deref(Ptr("a"), "z")
Output: "a"
```

**Example 3:**

```
Input:  Deref((*int)(nil), 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil pointer checks** | Reused from language basics: dereferencing nil panics, so guard first. |
| 2 | **Typed nil** | `(*int)(nil)` is a nil pointer of a concrete type — inference still works. |
| 3 | **Type parameters** | `[T any]` declares a type parameter; the caller (or inference) picks `T`. |

## Hint

One `if p == nil` guard, then `*p`.

## Validate

```bash
make verify
```
