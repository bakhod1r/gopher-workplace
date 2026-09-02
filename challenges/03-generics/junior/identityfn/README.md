# Identity

**Level:** junior  
**Topic:** 03-generics

## Context

Before a generic helper library grows, it starts with the smallest possible generic function: one that works for every type and does nothing to the value.

## Task

Implement the stub(s) in [identityfn.go](identityfn.go):

1. Implement `Identity`, returning the argument unchanged for any type `T`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Identity(7)
Output: 7
```

**Example 2:**

```
Input:  Identity("go")
Output: "go"
```

**Example 3:**

```
Input:  Identity(true)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type parameters** | `[T any]` declares a type parameter; the caller (or inference) picks `T`. |
| 2 | **Type inference** | You call `F(x)`, not `F[int](x)` — the compiler reads `T` from the argument. |
| 3 | **`any`** | `any` is an alias for `interface{}` — as a constraint it permits every type. |

## Hint

The body is one `return`. The interesting part is the signature.

## Validate

```bash
make verify
```
