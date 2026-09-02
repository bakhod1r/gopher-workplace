# Fluent Builder

**Level:** middle  
**Topic:** 03-generics

## Context

Test fixtures read better as `New().With(a).With(b).Build()` than as a pile of append statements.

## Task

Implement the stub(s) in [buildergen.go](buildergen.go):

1. Implement `With`, `WithAll`, and `Build`.
2. `With` and `WithAll` return the receiver so calls can chain.
3. `Build` returns a copy, so later chaining cannot mutate an already-built slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&Builder[int]{}).With(1).With(2).Build()
Output: [1 2]
```

**Example 2:**

```
Input:  WithAll(1,2,3).Build()
Output: [1 2 3]
```

**Example 3:**

```
Input:  (&Builder[int]{}).Build()
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fluent interfaces** | Returning `*Builder[T]` is what allows chaining. |
| 2 | **Pointer receiver required** | A value receiver would append to a copy and lose it. |
| 3 | **Defensive copies** | Handing out internal storage lets callers corrupt the structure. |

## Hint

Every mutating method returns `b` itself.

## Validate

```bash
make verify
```
