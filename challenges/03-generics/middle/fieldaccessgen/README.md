# No Field Access On T

**Level:** middle  
**Topic:** 03-generics

## Context

Someone tried to write `func Total[T any](items []T) int { ... v.Price ... }` and the compiler refused. The fix is worth internalising.

## Task

Implement the stub(s) in [fieldaccessgen.go](fieldaccessgen.go):

1. Implement `TotalPrice`, summing `price(v)` over the items.
2. Understand why the field cannot be read directly from a `T`.
3. Return `0` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TotalPrice(books, bookPrice)
Output: sum of prices
```

**Example 2:**

```
Input:  TotalPrice(coffees, coffeePrice)
Output: sum of prices
```

**Example 3:**

```
Input:  TotalPrice([]Book{}, bookPrice)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A real limit of Go generics** | Some things simply cannot be expressed — knowing which saves hours. |
| 2 | **No structural constraints** | Go constraints describe methods and type sets — never fields. |
| 3 | **Two workarounds** | Pass a projection function, or require a method through a constraint. |

## Hint

Constraints can demand methods, never fields — so pass a projection.

## Validate

```bash
make verify
```
